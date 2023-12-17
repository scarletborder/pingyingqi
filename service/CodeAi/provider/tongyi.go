package provider

import (
	"encoding/json"
	"io"
	"net/http"
	"pingyingqi/config"
	aiprovider "pingyingqi/models/AiProvider"
	"strings"

	"github.com/sirupsen/logrus"
)

type Tongyi struct {
	serviceName string
	website     string
	api_key     string
	// 服务商提供的api-key是永久的，不会expire
	// https://help.aliyun.com/zh/dashscope/developer-reference/activate-dashscope-and-create-an-api-key
}

func NewTongyi(Api_Key string) *Tongyi {
	return &Tongyi{serviceName: "通义千问", website: "https://qianwen.aliyun.com/", api_key: Api_Key}
}
func (t *Tongyi) GetServiceName() string {
	return t.serviceName
}

type tongyiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type tongyiInput struct {
	Messages []tongyiMessage `json:"messages"`
}
type tongyiParameter struct{}

type tongyiRequest struct {
	Model      string          `json:"model"`
	Input      tongyiInput     `json:"input"`
	Parameters tongyiParameter `json:"parameters"`
}

type tongyiOutput struct {
	Text string `json:"text"`
}
type tongyiResp struct {
	Output tongyiOutput `json:"output"`
}

func (t *Tongyi) Prompt(myPrompt string) (string, error) {
	var err error
	url := `https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation`
	payloadData := tongyiRequest{Model: `qwen-turbo`, Input: tongyiInput{Messages: []tongyiMessage{{Role: "system", Content: "You are a experienced programmer"}, {Role: "user", Content: myPrompt}}}}
	payloadJson, _ := json.Marshal(&payloadData)
	payload := strings.NewReader(string(payloadJson))

	// 	payload := strings.NewReader(`{
	//     "model": "qwen-turbo",
	//     "input":{
	//         "messages":[
	//             {
	//                 "role": "system",
	//                 "content": "You are a helpful assistant."
	//             },
	//             {
	//                 "role": "user",
	//                 "content": "你好，哪个公园距离我最近？"
	//             }
	//         ]
	//     },
	//     "parameters": {
	//     }
	// }`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return ``, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.EnvCfg.TongyiApiKey)
	res, err := client.Do(req)
	if err != nil {
		return ``, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ``, err
	}
	var tongyiresp tongyiResp
	err = json.Unmarshal(body, &tongyiresp)
	if err != nil {
		return ``, err
	}
	logrus.Debugf("The resp content is %s", tongyiresp.Output.Text)

	return tongyiresp.Output.Text, nil
}

func init() {
	if config.EnvCfg.TongyiApiKey != `` && config.EnvCfg.DefaultProvider != `nil` {
		aiprovider.AiHelper.Provider = append(aiprovider.AiHelper.Provider, NewTongyi(config.EnvCfg.TongyiApiKey))
	}
}
