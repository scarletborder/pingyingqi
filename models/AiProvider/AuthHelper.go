package aiprovider

import (
	"pingyingqi/config"
	"pingyingqi/service/CodeAi/provider"
)

type AiInterface interface {
	// 获取AI助手服务的名称
	GetServiceName() string
	// 通过凭证，保存AI助手的所有登录所需凭据
	Authme(...interface{})
	// 通过prompt获得代码执行信息(之后的业务逻辑进行解析)
	Prompt(string) string
}

var AiHelper aiHelper

type aiHelper struct {
	AuthTool AiInterface
}

func init() {
	AiHelper = aiHelper{}
	switch config.EnvCfg.AiserviceProvider {
	case "nil":

		return
	case "tongyi":
		AiHelper.AuthTool = provider.NewTongyi(config.EnvCfg.AiserviceKey1)
		return
	case "wenxin":
		AiHelper.AuthTool = provider.NewWenxin(config.EnvCfg.AiserviceKey1, config.EnvCfg.AiserviceKey2)
	}
}
