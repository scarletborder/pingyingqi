package provider

import "errors"

type Wenxin struct {
	serviceName  string
	website      string
	api_key      string
	secret_key   string
	access_token string
}

func NewWenxin(Api_Key string, Secret_Key string) (*Wenxin, error) {
	// 过期自动更新
	// go Ticker
	err := errors.New("fail")
	return &Wenxin{serviceName: "文心一言", website: "https://yiyan.baidu.com/"}, err
}

func (w *Wenxin) GetServiceName() string {
	return w.serviceName
}

func (w *Wenxin) Prompt(string) (string, error) {
	return "123"
}
