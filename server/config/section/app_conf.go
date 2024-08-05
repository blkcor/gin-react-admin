package section

import "gopkg.in/ini.v1"

var AppConfig App

type App struct {
	AccessKey   string
	RefreshKey  string
	JwtExpire   string
	Port        string
	RedirectUrl string
}

func InitAPP(appSection *ini.Section) {
	AppConfig = App{
		AccessKey:   appSection.Key("access_key").String(),
		RefreshKey:  appSection.Key("refresh_key").String(),
		JwtExpire:   appSection.Key("jwt_expire").String(),
		Port:        appSection.Key("port").String(),
		RedirectUrl: appSection.Key("redirect_url").String(),
	}
}
