package config

import (
	"github.com/astaxie/beego/config"
)

var conf config.Configer

func init() {
	conf, _ = config.NewConfig("ini", "config.ini")
}

func GetSetting(keystr string, typec int) (res interface{}) {
	if 0 == typec {
		res = conf.String(keystr)
	} else if 1 == typec {
		res, _ = conf.Int(keystr)
	}
	return
}
