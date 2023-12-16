package CodePro

import (
	"context"
	"pingyingqi/config"
	myrpc "pingyingqi/idl"
	cm "pingyingqi/models/CodeManual"
	codeai "pingyingqi/service/CodeAi"
	redis2 "pingyingqi/utils/redis"
	"strings"
)

/*
code:
- 0 success
- 1 inner wrong
- 2 illegal code
*/

type server struct {
	myrpc.UnimplementedCodeProProgramerServer
}

func (s *server) CodePro(ctx context.Context, in *myrpc.CodeProRequest) (*myrpc.CodeProResp, error) {
	var data string
	var code int
	var extraInfo string
	var code32 int32
	var err error
	// 如果config设置优先ai，那么直接丢给ai
	if config.EnvCfg.CompilerQueue == 1 {
		data, extraInfo, code = codeai.MainPrompt(in.Code, in.Lang)
		return &myrpc.CodeProResp{Data: data, Code: int32(code), Extra: extraInfo}, nil
	}

	// 人工CodeManual
	if cm.CodeMan.CouldProgram(in.Lang) {
		data, code32, err = cm.CodeMan.Exec(in.Code, in.Lang)
		return &myrpc.CodeProResp{Data: data, Code: code32, Extra: ``}, err
	}

	// 缺省CodeAi
	// ai()
	data, extraInfo, code = codeai.MainPrompt(in.Code, in.Lang)
	return &myrpc.CodeProResp{Data: data, Code: int32(code), Extra: extraInfo}, nil
}

func (s *server) Dislike(ctx context.Context, in *myrpc.DislikedPackage) (*myrpc.DislikedResp, error) {
	rawstr := in.GetPack()
	modstr := strings.Split(rawstr, ";")
	if modstr[0] == "0" {
		// pass
		_, err := redis2.Client.SRem(ctx, "dislike", modstr[1]).Result()
		if err != nil {
			return &myrpc.DislikedResp{Data: err.Error(), Code: 1}, nil
		} else {
			return &myrpc.DislikedResp{Data: "successfully pass", Code: 0}, nil
		}
	} else {
		// block
		_, err := redis2.Client.SAdd(ctx, "dislike", modstr[1]).Result()
		if err != nil {
			return &myrpc.DislikedResp{Data: err.Error(), Code: 1}, nil
		} else {
			return &myrpc.DislikedResp{Data: "successfully block", Code: 0}, nil
		}
	}
}
