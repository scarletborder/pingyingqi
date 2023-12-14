package CodePro

import (
	"context"
	"pingyingqi/config"
	myrpc "pingyingqi/idl"
	cm "pingyingqi/models/CodeManual"
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
	var code int32
	var err error = nil
	// 如果config设置优先ai，那么直接丢给ai
	if config.EnvCfg.CompilerQueue == 1 {
		// ai()
		return &myrpc.CodeProResp{Data: "还没做好ai功能", Code: 1}, err
	}

	// 人工CodeManual
	if cm.CodeMan.CouldProgram(in.Lang) {
		data, code, err = cm.CodeMan.Exec(in.Code, in.Lang)
		return &myrpc.CodeProResp{Data: data, Code: code}, err
	}

	// 缺省CodeAi
	// ai()
	return &myrpc.CodeProResp{Data: "还没做好ai功能", Code: 1}, err
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
