package codemanual

// 检验某语言(或其别名)是否可以被人工编译
func (c CodeManual) CouldProgram(lang string) bool {
	if c.lang.Contains(lang) {
		return true
	}
	_, ok := c.langAlias[lang]
	return ok
}

func (c CodeManual) Exec(inCode string, lang string) (data string, code int32, err error) {
	var realLang string // 真实的语言名称,如python,golang,c,cpp

	if c.lang.Contains(lang) {
		realLang = lang
	} else {
		var ok bool
		realLang, ok = c.langAlias[lang]
		if !ok {
			return "This language is not supported manually", 1, nil
		}
	}
	c.langExec[realLang].Exec(inCode, &data, &code)
	err = nil
	return
}
