package section

import "gopkg.in/ini.v1"

var DBConfig DataBase

type DataBase struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func InitDataBase(dbSection *ini.Section) {
	DBConfig = DataBase{
		Host:     dbSection.Key("host").String(),
		Port:     dbSection.Key("port").String(),
		User:     dbSection.Key("user").String(),
		Password: dbSection.Key("password").String(),
		DBName:   dbSection.Key("dbname").String(),
	}
}
