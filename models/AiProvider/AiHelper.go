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

	// 将code包装后，通过prompt获得代码执行信息(之后的业务逻辑进行解析)
	Prompt(string) (string, error)
}

var AiHelper aiHelper

type aiHelper struct {
	Provider []AiInterface
}

// 创建(包括认证过程)中出错导致无法创建Provider的情况
func FailOnAiProviderCreate(ServiceName string, err error) {
	logrus.Errorf("AI provider %s couldn't create, reason %s", ServiceName, err.Error())
}

func init() {
	AiHelper = aiHelper{}

	// 轮流创建Provider
	if config.EnvCfg.TongyiApiKey != "nil" {
		AiHelper.Provider = append(AiHelper.Provider, provider.NewTongyi(config.EnvCfg.TongyiApiKey))
	}
	if config.EnvCfg.WenxinApiKey != "nil" && config.EnvCfg.WenxinSecretKey != "nil" {
		Wenxin, err := provider.NewWenxin(config.EnvCfg.WenxinApiKey, config.EnvCfg.WenxinSecretKey)
		if err != nil {
			FailOnAiProviderCreate("wenxin", err)
		} else {
			AiHelper.Provider = append(AiHelper.Provider, Wenxin)
		}
	}
	AiHelper.Provider = append(AiHelper.Provider, provider.NewNil())

	// 交换
	for idx := range AiHelper.Provider {
		if AiHelper.Provider[idx].GetServiceName() == config.EnvCfg.DefaultProvider {
			tmp := AiHelper.Provider[0]
			AiHelper.Provider[0] = AiHelper.Provider[idx]
			AiHelper.Provider[idx] = tmp
		}
	}
}
