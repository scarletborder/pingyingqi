package aiprovider

type AiProvider interface {
	// 获取AI助手服务的名称
	GetServiceName() string
	//[deprecated! now through New*() method to get and maintain] 通过凭证，保存AI助手的所有登录所需凭据
	// Authme(...interface{})

	// 将code包装后，通过prompt获得代码执行信息(之后的业务逻辑进行解析)
	Prompt(string) (string, error)
}
