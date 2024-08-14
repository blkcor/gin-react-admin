package randx

import (
	"github.com/blkcor/gin-react-admin/consts"
	"math/rand"
	"time"
)

// GetRandomDefaultAvatar 获取随机默认头像
func GetRandomDefaultAvatar() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return consts.UserDefaultAvatars[r.Intn(len(consts.UserDefaultAvatars))]
}
