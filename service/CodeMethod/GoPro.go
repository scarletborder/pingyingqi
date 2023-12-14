package codemethod

// 用于绑定CodeMan所用到的方法

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"path"
	"pingyingqi/config"
	redis2 "pingyingqi/utils/redis"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type GoPro struct{}

func (g GoPro) Exec(code string, outData *string, statusCode *int32) {
	ctx := context.Background()

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
		*outData = err.Error()
		*statusCode = 1
		return
	}

	cmd := exec.Command("go", "mod", "init")
	cmd.Dir = tempPath
	err = cmd.Run()
	if err != nil {
		*outData = err.Error()
		*statusCode = 1
		return
	}

	cmd = exec.Command("go", "fmt")
	cmd.Dir = tempPath
	err = cmd.Run()
	if err != nil {
		*outData = err.Error()
		*statusCode = 1
		return
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
					*outData = err.Error()
					*statusCode = 1
					return
				}
				if block {
					*outData = "blocked package " + importPack
					*statusCode = 2
					return
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
				*outData = err.Error()
				*statusCode = 1
				return
			}
			if block {
				*outData = "blocked package " + importPack
				*statusCode = 2
				return
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
		*outData = err.Error()
		*statusCode = 1
		return
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
			*outData = outBut.String()
			*statusCode = 3
			return
		case <-done:
			*outData = outBut.String()
			*statusCode = 0
			return
		default:
			if len(outBut.String()) > config.EnvCfg.MaximumOutput {
				logrus.Println("kill")
				err = cmd.Process.Kill()
				logrus.Println(err)
				*outData = outBut.String()
				*statusCode = 0
				return
			}
		}
	}

}
