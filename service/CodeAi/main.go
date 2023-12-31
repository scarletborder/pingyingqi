package codeai

import (
	aiprovider "pingyingqi/models/AiProvider"
	_ "pingyingqi/service/CodeAi/provider"

	"github.com/sirupsen/logrus"
)

func FailOnAiPrompt(ServiceName string, err error) {
	logrus.WithField("ServiceName", ServiceName).Errorf("AI Provider prompt error\n%s", err.Error())
}

func MainPrompt(code string, lang string) (string, string, int) {
	var data string
	var extraInfo string
	var statusCode int
	var err error
	code, err = WrapperCode(code, lang)
	if err != nil {
		logrus.Errorf("Unable to create wrapped prompt\n%s", err)
		return `inner wrong`, ``, 1
	}

	aiprovider.AiHelper.ResetQueue() // 按照default provider，重新规划顺序
	for _, provider := range aiprovider.AiHelper.Provider {
		data, extraInfo, statusCode, err = DecodeJson(provider.Prompt(code))
		if err != nil {
			FailOnAiPrompt(provider.GetServiceName(), err)
		} else {
			// 第一次成功就break loop
			// 保证"nil"的prompt的err一定是nil
			break
		}
	}
	return data, extraInfo, statusCode
}
