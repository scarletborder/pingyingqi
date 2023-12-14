package aiprovider

import (
	"pingyingqi/config"
	"pingyingqi/service/CodeAi/provider"
	_ "pingyingqi/utils/logging"

	"github.com/sirupsen/logrus"
)

type AiInterface interface {
	// 获取AI助手服务的名称
	GetServiceName() string
	//[deprecated! now through New*() method to get and maintain] 通过凭证，保存AI助手的所有登录所需凭据
	// Authme(...interface{})
	// 通过prompt获得代码执行信息(之后的业务逻辑进行解析)
	Prompt(string) string
}

var AiHelper aiHelper

type aiHelper struct {
	Provider AiInterface
}

// 创建(包括认证过程)中出错导致无法创建Provider的情况
func FailOnAiProviderCreate(ServiceName string, err error) {
	logrus.Errorf("AI provider %s couldn't create, reason %s", ServiceName, err.Error())
}

func init() {
	AiHelper = aiHelper{}

	// 首先创建默认的Provider,并放在AiHelper第一位

	// 轮流创建Provider

	switch config.EnvCfg.AiserviceProvider {
	case "nil":

		return
	case "tongyi":
		AiHelper.Provider = provider.NewTongyi(config.EnvCfg.AiserviceKey1)
		return
	case "wenxin":
		var err error
		AiHelper.Provider, err = provider.NewWenxin(config.EnvCfg.AiserviceKey1, config.EnvCfg.AiserviceKey2)
		if err != nil {
			FailOnAiProviderCreate("wenxin", err)
		}
	}
}

// 轮训prompt
// func MainPrompt()
