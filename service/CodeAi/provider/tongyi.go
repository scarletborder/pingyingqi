package provider

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

func (t *Tongyi) Prompt(myPrompt string) (string, error) {
	
}
