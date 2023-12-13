package exlang

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"scarletborders.top/pingyingqi/config"
	myrpc "scarletborders.top/pingyingqi/idl"
	redis2 "scarletborders.top/pingyingqi/redis"
)

/*
code:
- 0 success
- 1 inner wrong
- 2 illegal code
*/

type server struct {
	myrpc.UnimplementedExlangProgramerServer
}

func (s *server) PyPro(ctx context.Context, in *myrpc.ExlangRequest) (*myrpc.ExlangResp, error) {
	// cmd := exec.Command("python3")
	// 审查PyPro
	code := in.GetCode()

	if strings.Contains(code, "import") {
		code += "\n"
		leftCode := code
		var importCode []string
		var importIndex int
		for {
			if importIndex = strings.Index(leftCode, "import"); importIndex == -1 {
				break
			}
			newLine := strings.Index(leftCode[importIndex:], "\n")
			newImport := leftCode[importIndex+6 : importIndex+newLine]
			newImport = strings.TrimSpace(newImport)
			importCode = append(importCode, newImport)
			logrus.Println("new pack detect ", newImport)
			leftCode = leftCode[importIndex+newLine:]
		}
		// 轮训
		for _, importPack := range importCode {
			block, err := redis2.Client.SIsMember(ctx, "dislike", importPack).Result()
			if err != nil {
				return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
			}
			if block {
				return &myrpc.ExlangResp{Data: "blocked package " + importPack, Code: 2}, nil
			}
		}
	}

	tempPath, err := os.MkdirTemp(config.EnvCfg.DefaultDir+"/pyq", "python*")
	if err != nil {
		panic(err)
	}

	filePath := path.Join(tempPath, "main.py")
	defer func() {
		err = os.RemoveAll(tempPath)
	}()

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString(code)
	if err = f.Close(); err != nil {
		return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
	}

	cmd := exec.Command("python", "main.py")
	cmd.Dir = tempPath
	var outBut bytes.Buffer
	cmd.Stdout = io.MultiWriter(io.Discard, &outBut)
	proTimer := time.NewTimer(time.Duration(config.EnvCfg.MaximumDelay) * time.Second)
	defer proTimer.Stop()

	err = cmd.Start()
	if err != nil {
		return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
	}

	// 逻辑，先开协程跑并有wait，wait完外部直接break，外部计时器如果触发也可以break
	var done chan bool = make(chan bool, 1)
	defer close(done)
	go func(thecmd *exec.Cmd) {
		thecmd.Wait()
		done <- true
	}(cmd)

	for {
		select {
		case <-proTimer.C:
			cmd.Process.Kill()
			return &myrpc.ExlangResp{Data: outBut.String(), Code: 3}, nil
		case <-done:
			return &myrpc.ExlangResp{Data: outBut.String(), Code: 0}, nil
		default:
			if len(outBut.String()) > config.EnvCfg.MaximumOutput {
				cmd.Process.Kill()
				return &myrpc.ExlangResp{Data: outBut.String(), Code: 0}, nil
			}
		}
	}
}

func (s *server) GoPro(ctx context.Context, in *myrpc.ExlangRequest) (*myrpc.ExlangResp, error) {
	code := in.GetCode()

	tempPath, err := os.MkdirTemp(config.EnvCfg.DefaultDir+"/pyq", "golang*")
	if err != nil {
		panic(err)
	}

	filePath := path.Join(tempPath, "main.go")
	defer func() { os.RemoveAll(tempPath) }()

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString(code)
	if err = f.Close(); err != nil {
		return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
	}

	cmd := exec.Command("go", "mod", "init")
	cmd.Dir = tempPath
	err = cmd.Run()
	if err != nil {
		return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
	}

	cmd = exec.Command("go", "fmt")
	cmd.Dir = tempPath
	err = cmd.Run()
	if err != nil {
		return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
	}

	// new import inspect
	// delay, in order to detect wrong code
	buf, _ := os.ReadFile(filePath)
	code = string(buf)
	if strings.Contains(code, "import") {
		code += "\n"
		// 鉴别是否是单个导入
		var newImport string
		var importIndex int = strings.Index(code, "import")
		newLine := strings.Index(code[importIndex+6:], "\n")
		newImport = code[importIndex+6 : importIndex+6+newLine]
		if strings.Contains(newImport, "(") {
			// 多行导入
			logrus.Println("mutli line import")
			var importCode []string

			endImport := strings.Index(code[importIndex:], ")")
			code = code[importIndex+6+newLine+1 : importIndex+6+newLine+1+endImport]
			for {
				if newLine = strings.Index(code, "\n"); strings.Contains(code[:newLine], ")") {
					break
				}
				newImport = code[:newLine]
				newImport = strings.TrimSpace(newImport)
				newImport = strings.Trim(newImport, "\"")
				importCode = append(importCode, newImport)
				logrus.Println("new pack detect ", newImport)
				code = code[newLine+1:]
			}

			// 轮训
			for _, importPack := range importCode {
				block, err := redis2.Client.SIsMember(ctx, "dislike", importPack).Result()
				if err != nil {
					return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
				}
				if block {
					return &myrpc.ExlangResp{Data: "blocked package " + importPack, Code: 2}, nil
				}
			}
		} else {
			// 单行导入
			logrus.Println("single line import")
			importPack := code[importIndex+6 : importIndex+6+newLine]
			importPack = strings.TrimSpace(importPack)
			importPack = strings.Trim(importPack, "\"")
			logrus.Println("pack ", importPack, " detect")
			block, err := redis2.Client.SIsMember(ctx, "dislike", importPack).Result()
			if err != nil {
				return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
			}
			if block {
				return &myrpc.ExlangResp{Data: "blocked package " + importPack, Code: 2}, nil
			}
		}
	}
	//end new inspect

	cmd = exec.Command("go", "run", "main.go")
	cmd.Dir = tempPath
	var outBut bytes.Buffer
	cmd.Stdout = io.MultiWriter(io.Discard, &outBut)
	proTimer := time.NewTimer(time.Duration(config.EnvCfg.MaximumDelay) * time.Second)
	defer proTimer.Stop()

	err = cmd.Start()
	if err != nil {
		return &myrpc.ExlangResp{Data: err.Error(), Code: 1}, nil
	}

	// 逻辑，先开协程跑并有wait，wait完外部直接break，外部计时器如果触发也可以break
	var done chan bool = make(chan bool, 1)
	defer close(done)
	go func(thecmd *exec.Cmd) {
		thecmd.Wait()
		done <- true
	}(cmd)

	for {
		select {
		case <-proTimer.C:
			cmd.Process.Kill()
			return &myrpc.ExlangResp{Data: outBut.String(), Code: 3}, nil
		case <-done:
			return &myrpc.ExlangResp{Data: outBut.String(), Code: 0}, nil
		default:
			if len(outBut.String()) > config.EnvCfg.MaximumOutput {
				logrus.Println("kill")
				err = cmd.Process.Kill()
				logrus.Println(err)
				return &myrpc.ExlangResp{Data: outBut.String(), Code: 0}, nil
			}
		}
	}

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
