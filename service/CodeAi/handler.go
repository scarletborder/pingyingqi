package codeai

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

/** 包装code
 * @return string 完整的prompt(提供给json序列化)
 */
func WrapperCode(code string, lang string) (string, error) {
	f, err := os.ReadFile("./service/CodeAi/wrapper.txt")
	if err != nil {
		return "", err
	}
	return string(f) + lang + "\n" + code + "\n```", nil
}

// 解析json
type AiProviderResp struct {
	Code int    `json:"Code"`
	Data string `json:"Data"`
}

func DecodeJson(inData string, inerr error) (data string, extraInfo string, statusCode int, err error) {
	if inerr != nil {
		// 内部错误
		err = inerr
		return
	}
	jsonCodeBegin := strings.Index(inData, "```")
	if jsonCodeBegin == -1 {
		data = ``
		extraInfo = inData
		statusCode = 1
		err = nil
		return
	}
	jsonCodeEnd := strings.Index(inData[jsonCodeBegin+3:], "```")
	if jsonCodeEnd == -1 {
		data = ``
		extraInfo = inData
		statusCode = 1
		err = nil
		return
	}
	inData = inData[jsonCodeBegin : jsonCodeBegin+jsonCodeEnd+6]
	inData = strings.Trim(inData, "`")
	inData = strings.TrimSpace(inData)
	inData = strings.Trim(inData, "json")
	inData = strings.Trim(inData, "Json")
	inData = strings.TrimSpace(inData)
	var resp AiProviderResp
	err = json.Unmarshal([]byte(inData), &resp)

	if err != nil {
		err = nil
		data = inData
		statusCode = 0
		logrus.Warnf("Wrong result format\n%s", data)
	} else {
		data = resp.Data
		statusCode = resp.Code
	}
	return
}
