package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"scarletborders.top/pingyingqi/config"
	exlang "scarletborders.top/pingyingqi/exlang"
	_ "scarletborders.top/pingyingqi/redis"
)

func main() {
	defer os.RemoveAll(config.EnvCfg.DefaultDir + "/pyq")
	defer func() {
		if err := recover(); err != nil {
			logrus.Println(err)
		}
	}()
	exlang.Exlanglis()
}
