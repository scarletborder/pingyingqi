package provider

type Wenxin struct {
	serviceName  string
	website      string
	api_key      string
	secret_key   string
	access_token string
}

func NewWenxin(Api_Key string, Secret_Key string) Wenxin {
	// 过期自动更新
	// go Ticker

	return Wenxin{serviceName: "文心一言", website: "https://yiyan.baidu.com/"}
}

func (w Wenxin) GetServiceName() string {
	return w.serviceName
}

func (w Wenxin) Authme(...interface{}) {

}
func (w Wenxin) Prompt(string) string {
	return "123"
}
