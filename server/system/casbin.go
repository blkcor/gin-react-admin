package system

import (
	"github.com/blkcor/gin-react-admin/core/db"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

// 持久化到数据库
var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (c *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(db.DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer("rbac_model.conf", a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
