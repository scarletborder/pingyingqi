package codeai

import (
	"encoding/json"
	"os"
)

// 包装code
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
	var resp AiProviderResp
	err = json.Unmarshal([]byte(inData), &resp)
	if err != nil {
		err = nil
		data = inData
		statusCode = 0
	} else {
		data = resp.Data
		statusCode = resp.Code
	}
	return
}
