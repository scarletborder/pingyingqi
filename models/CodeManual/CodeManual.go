package codemanual

import (
	cm "pingyingqi/service/CodeMethod"

	mapset "github.com/deckarep/golang-set"
)

type ExecInterface interface {
	Exec(string, *string, *int32)
}

type CodeManual struct {
	lang      mapset.Set
	langAlias map[string]string
	langExec  map[string]ExecInterface
}

var CodeMan CodeManual

func init() {
	CodeMan = CodeManual{}
	CodeMan.lang = mapset.NewSet()
	CodeMan.langAlias = make(map[string]string)
	CodeMan.langExec = make(map[string]ExecInterface)

	// 注册语言
	CodeMan.lang.Add("golang")
	CodeMan.lang.Add("python")

	// 注册别名
	CodeMan.langAlias["go"] = "golang"
	CodeMan.langAlias["py"] = "python"

	// 绑定函数
	CodeMan.langExec["golang"] = cm.GoPro{}
	CodeMan.langExec["python"] = cm.PyPro{}
}
