package CodePro

import (
	"context"
	"strings"

	myrpc "pingyingqi/idl"
	redis2 "pingyingqi/redis"
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
	// 如果config设置优先ai，那么直接丢给ai

	// 人工CodeManual

	// 缺省CodeAi
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
