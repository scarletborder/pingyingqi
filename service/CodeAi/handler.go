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

func DecodeJson(inData string, inerr error) (data string, statusCode int, err error) {
	if inerr != nil {
		// 内部错误
		err = inerr
		return
	}
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
