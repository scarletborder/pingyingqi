package exlang

import (
	"context"
	"fmt"
	"net"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"scarletborders.top/pingyingqi/config"
	myrpc "scarletborders.top/pingyingqi/idl"
	"scarletborders.top/pingyingqi/redis"
)

func init() {
	// 创建用到的名为dislike的set
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := redis.Client.SAdd(ctx, "dislike", 0).Result()
	if err != nil {
		logrus.Errorln("redis set initial meet some wrong", err)
	}

	dir := path.Join(config.EnvCfg.DefaultDir, "pyq")
	if err := os.MkdirAll(dir, os.FileMode(0755)); err != nil {
		panic(err)
	}
	logrus.Println("已经创建文件夹了", dir)
}

func Exlanglis() {
	lis, err := net.Listen("tcp", config.EnvCfg.ListenHost+":"+fmt.Sprint(config.EnvCfg.ListenPort))
	if err != nil {
		logrus.Fatalln("rpc fail to listen:", err)
	}
	s := grpc.NewServer()
	myrpc.RegisterExlangProgramerServer(s, &server{})
	logrus.Println("ready to serve", config.EnvCfg.ListenHost+":"+fmt.Sprint(config.EnvCfg.ListenPort))
	err = s.Serve(lis)

	if err != nil {
		logrus.Fatalln("rpc fail to serve:", err)
		return
	}
}
