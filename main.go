package main

import (
	"os"

	config "pingyingqi/config"
	_ "pingyingqi/redis"
	CodePro "pingyingqi/service/CodePro"

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
