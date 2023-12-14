package provider

type Nil struct {
	serviceName string
}

func NewNil() Nil {
	return Nil{serviceName: "empty AI provider"}
}
func (n Nil) GetServiceName() string {
	return n.serviceName
}

func (n Nil) Prompt(string) string {
	return `[WARNING]no available AI provider!
This may caused by your not offering any AISERVICE_PROVIDER setting.
You can change the AISERVICE_PROVIDER in your "config/*.env".
Also, if all AI providers can not prompt, it will also cause this problem
You may need to check your AI providers' authentication in log`
}
