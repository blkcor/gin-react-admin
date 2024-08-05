package main

import (
	"github.com/blkcor/gin-react-admin/config"
	"github.com/blkcor/gin-react-admin/core/cache"
	"github.com/blkcor/gin-react-admin/core/db"
	"github.com/blkcor/gin-react-admin/core/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	configDone := make(chan bool, 1)
	//初始化配置信息
	config.Init(configDone)
	<-configDone
	wg.Add(2)
	//初始化数据库
	go func() {
		defer wg.Done()
		db.Init()
	}()
	//初始化redis
	go func() {
		defer wg.Done()
		cache.Init()
	}()
	wg.Wait()
	//初始化服务
	server.Init()
}
