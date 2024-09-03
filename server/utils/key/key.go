package key

import "strconv"

// OnlineUserSetKey 在线用户集合key
const OnlineUserSetKey = "online_user_set"
const OnlineUserPrefix = "online_user:"

// GenerateOnlineUserKey 生成在线用户key
func GenerateOnlineUserKey(userId uint32) string {
	return OnlineUserPrefix + strconv.Itoa(int(userId))
}
