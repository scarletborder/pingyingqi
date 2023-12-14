package main

import (
	"os"

	config "pingyingqi/config"
	CodePro "pingyingqi/service/CodePro"
	_ "pingyingqi/utils/redis"

	"github.com/sirupsen/logrus"
)

func main() {
	defer os.RemoveAll(config.EnvCfg.DefaultDir + "/pyq")
	defer func() {
		if err := recover(); err != nil {
			logrus.Println(err)
		}
	}()
	CodePro.CodeProListen()
}
