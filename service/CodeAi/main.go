package codeai

import (
	aiprovider "pingyingqi/models/AiProvider"

	"github.com/sirupsen/logrus"
)

func FailOnAiPrompt(ServiceName string, err error) {
	logrus.WithField("ServiceName", ServiceName).Errorf("AI Provider prompt error %s", err.Error())
}

func MainPrompt(code string, lang string) (string, int) {
	var data string
	var statusCode int
	var err error
	code, err = WrapperCode(code, lang)

	for _, provider := range aiprovider.AiHelper.Provider {
		data, statusCode, err = DecodeJson(provider.Prompt(code))
		if err != nil {
			FailOnAiPrompt(provider.GetServiceName(), err)
		} else {
			// 第一次成功就break loop
			break
		}
	}
	return data, statusCode

}
