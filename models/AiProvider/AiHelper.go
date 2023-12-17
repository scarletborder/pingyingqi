package aiprovider

import (
	"pingyingqi/config"
	// "pingyingqi/service/CodeAi/provider"
	_ "pingyingqi/utils/logging"

	"github.com/sirupsen/logrus"
)

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
