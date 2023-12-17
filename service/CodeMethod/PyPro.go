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
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

type PyPro struct{}

func (p PyPro) Exec(code string, outData *string, statusCode *int32) {
	// cmd := exec.Command("python3")
	// 审查PyPro
	ctx := context.Background()

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
		*outData = err.Error()
		*statusCode = 1
		return
	}

	cmd := exec.Command("python", "main.py")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, //使得 Shell 进程开辟新的 PGID, 即 Shell 进程的 PID, 它后面创建的所有子进程都属于该进程组
	}
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
			p, err := os.FindProcess(cmd.Process.Pid)
			if err != nil {
				logrus.Errorf("Couldn't close process, %s", err.Error())
			}
			err = p.Kill()
			if err != nil {
				logrus.Errorf("Couldn't close process, %s", err.Error())
			}
			*outData = outBut.String()
			*statusCode = 3
			return
		case <-done:
			*outData = outBut.String()
			*statusCode = 0
			return
		default:
			if len(outBut.String()) > config.EnvCfg.MaximumOutput {
				cmd.Process.Kill()
				*outData = outBut.String()
				*statusCode = 0
				return
			}
		}
	}
}
