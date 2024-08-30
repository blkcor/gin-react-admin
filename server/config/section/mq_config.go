package section

import "gopkg.in/ini.v1"

var MQConfig MQ

type MQ struct {
	URL string
}

func InitMQ(mqSection *ini.Section) {
	MQConfig = MQ{
		URL: mqSection.Key("url").String(),
	}
}
