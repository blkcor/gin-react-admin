package system

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/models/vo"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"net/http"
)

type ResourceMonitor struct {
	upgrader websocket.Upgrader
	mqConn   *amqp.Connection
	mqChan   *amqp.Channel
}

func NewResourceMonitor() *ResourceMonitor {
	return &ResourceMonitor{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // 允许所有来源，生产环境中应该更严格
			},
		},
	}
}

func (rm *ResourceMonitor) setupMQ() error {
	var err error
	rm.mqConn, err = amqp.Dial(section.MQConfig.URL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	rm.mqChan, err = rm.mqConn.Channel()
	if err != nil {
		rm.mqConn.Close() // Close the connection if channel creation fails
		return fmt.Errorf("failed to open a channel: %v", err)
	}

	return nil
}

func (rm *ResourceMonitor) cleanupMQ() {
	if rm.mqChan != nil {
		rm.mqChan.Close()
	}
	if rm.mqConn != nil {
		rm.mqConn.Close()
	}
}

func (rm *ResourceMonitor) ServerResourceMonitor(c *gin.Context) {
	// 设置 MQ
	if err := rm.setupMQ(); err != nil {
		logger.Errorf("Failed to setup MQ: %v", err)
		c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Success: false,
			Message: "Internal server error: " + err.Error(),
			Data:    nil,
		})
		return
	}
	defer rm.cleanupMQ()

	// 升级到 WebSocket
	ws, err := rm.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Errorf("Failed to upgrade to WebSocket: %v", err)
		c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Success: false,
			Message: "无法升级到 WebSocket: " + err.Error(),
			Data:    nil,
		})
		return
	}
	defer ws.Close()

	// 创建一个 context，用于控制 goroutines
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	// 启动消息处理
	errChan := make(chan error, 1)
	go rm.handleMessages(ctx, ws, errChan)

	// 启动定时任务
	if err := rm.startCron(ctx); err != nil {
		logger.Errorf("Failed to start cron job: %v", err)
		c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Success: false,
			Message: "无法开启定时任务: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 等待连接关闭或错误发生
	select {
	case <-ctx.Done():
		logger.Info("WebSocket connection closed")
	case err := <-errChan:
		logger.Errorf("Error in message handling: %v", err)
		ws.WriteMessage(websocket.CloseMessage, []byte("Internal server error"))
	}
}

func (rm *ResourceMonitor) handleMessages(ctx context.Context, ws *websocket.Conn, errChan chan<- error) {
	msgs, err := rm.mqChan.Consume(
		"system_monitor",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		errChan <- fmt.Errorf("failed to register a consumer: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				errChan <- fmt.Errorf("RabbitMQ channel closed")
				return
			}
			err := ws.WriteMessage(websocket.TextMessage, msg.Body)
			if err != nil {
				errChan <- fmt.Errorf("failed to write message: %v", err)
				return
			}
		}
	}
}

func (rm *ResourceMonitor) startCron(ctx context.Context) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return fmt.Errorf("failed to create scheduler: %v", err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(2*time.Second),
		gocron.NewTask(rm.getAndPublishServerInfo),
	)
	if err != nil {
		return fmt.Errorf("failed to create job: %v", err)
	}

	s.Start()

	go func() {
		<-ctx.Done()
		s.Shutdown()
	}()

	return nil
}

func (rm *ResourceMonitor) getAndPublishServerInfo() {
	info := rm.getServerInfo()
	data, err := json.Marshal(info)
	if err != nil {
		logger.Errorf("Failed to marshal server info: %v", err)
		return
	}

	err = rm.mqChan.Publish(
		"",
		"system_monitor",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	if err != nil {
		logger.Errorf("Failed to publish message: %v", err)
		// Implement retry mechanism here
		return
	}
}

// getServerInfo 获取服务器信息
func (rm *ResourceMonitor) getServerInfo() vo.SystemVO {
	// CPU 信息
	cpuUsage := vo.CPUUsage{
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	p, err := cpu.Percent(time.Second, false)
	if err != nil {
		logger.Errorf("Failed to get CPU usage: %v", err)
	}
	cpuUsage.Value = p[0]

	// 内存信息
	available := vo.MemoryUsage{
		Name: "可用内存",
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		logger.Errorf("Failed to get memory info: %v", err)
	}
	available.Value = float64(memInfo.Available) / 1024 / 1024
	used := vo.MemoryUsage{
		Name: "已用内存",
	}
	used.Value = float64(memInfo.Used) / 1024 / 1024

	// 磁盘信息
	diskPaths := []string{"/", "/home", "/var"}
	disks := make([]vo.DiskUsage, 3)
	for i, path := range diskPaths {
		usage, err := disk.Usage(path)
		if err != nil {
			logger.Errorf("Failed to get disk usage: %v", err)
		}
		disks[i] = vo.DiskUsage{
			Name: path,
			Used: float64(usage.Used) / 1024 / 1024,
			Free: float64(usage.Free) / 1024 / 1024,
		}
	}
	// 网络信息
	networkUsage := vo.NetWorkUsage{
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	ioStat, err := net.IOCounters(true)
	if err != nil {
		logger.Errorf("Failed to get network IO: %v", err)
	}
	networkUsage.In = ioStat[0].BytesRecv
	networkUsage.Out = ioStat[0].BytesSent

	return vo.SystemVO{
		CPUUsage:     cpuUsage,
		MemoryUsage:  available,
		DiskUsage:    disks,
		NetWorkUsage: networkUsage,
	}
}
