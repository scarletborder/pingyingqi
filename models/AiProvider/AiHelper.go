package aiprovider

import (
	"pingyingqi/config"
	// "pingyingqi/service/CodeAi/provider"
	_ "pingyingqi/utils/logging"

	"github.com/sirupsen/logrus"
)

type AiProvider interface {
	// 获取AI助手服务的名称
	GetServiceName() string
	//[deprecated! now through New*() method to get and maintain] 通过凭证，保存AI助手的所有登录所需凭据
	// Authme(...interface{})

	// 将code包装后，通过prompt获得代码执行信息(之后的业务逻辑进行解析)
	Prompt(string) (string, error)
}

var AiHelper aiHelper

type aiHelper struct {
	Provider []AiProvider
}

// 创建(包括认证过程)中出错导致无法创建Provider的情况
func FailOnAiProviderCreate(ServiceName string, err error) {
	logrus.Errorf("AI provider %s couldn't create, reason %s", ServiceName, err.Error())
}

func (a aiHelper) ResetQueue() {
	// 交换
	for idx := range AiHelper.Provider {
		if AiHelper.Provider[idx].GetServiceName() == config.EnvCfg.DefaultProvider {
			tmp := AiHelper.Provider[0]
			AiHelper.Provider[0] = AiHelper.Provider[idx]
			AiHelper.Provider[idx] = tmp
			break
		}
	}
}

func init() {
	AiHelper = aiHelper{}
}
