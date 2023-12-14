package provider

type Nil struct {
	serviceName string
}

func NewNil() Nil {
	return Nil{serviceName: "通义千问"}
}
func (n Nil) GetServiceName() string {
	return n.serviceName
}

func (n Nil) Prompt(string) string {
	return "123"
}
