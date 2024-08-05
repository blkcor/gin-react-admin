package section

import "gopkg.in/ini.v1"

var RedisConfig Redis

type Redis struct {
	Host string
	Port string
	Db   int
}

func InitRedis(redisSection *ini.Section) {
	RedisConfig = Redis{
		Host: redisSection.Key("host").String(),
		Port: redisSection.Key("port").String(),
		Db:   redisSection.Key("db").MustInt(),
	}
}
