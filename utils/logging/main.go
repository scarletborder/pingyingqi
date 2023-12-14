package logging

// 默认先初始化我使得logrus的定制正确
import (
	"os"
	"pingyingqi/config"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	LevelBuffer   []log.Level
	ContentBuffer []string
	BufferSize    = 8
)

type FileLogHook struct{}

var LogFilePath string

func (h *FileLogHook) Levels() []log.Level {
	return []log.Level{log.WarnLevel, log.ErrorLevel, log.FatalLevel}
}

func (h *FileLogHook) Fire(entry *log.Entry) error {
	ContentBuffer = append(ContentBuffer, entry.Message)
	LevelBuffer = append(LevelBuffer, entry.Level)
	if len(ContentBuffer) > BufferSize {
		// 输出
		fout, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Error("无法创建或打开log文件")
		}
		for idx := range ContentBuffer {
			fout.WriteString(time.Now().Format("01_02_15:04:05") + "[" + LevelBuffer[idx].String() + "]:" + ContentBuffer[idx])
			fout.WriteString("\n")
		}
		ContentBuffer = nil
		LevelBuffer = nil
	}
	return nil
}

func init() {
	log.SetReportCaller(true)

	switch config.EnvCfg.LoggerLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN", "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	}

	LogFilePath = "./log/" + time.Now().Format("15_04_05") + ".log"
	log.AddHook(&FileLogHook{})
}
