package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var EnvCfg envConfig

type envConfig struct {
	ListenHost string `env:"LISTEN_HOST" envDefault:"127.0.0.1"`
	ListenPort uint16 `env:"LISTEN_PORT" envDefault:"28000"`

	CompilerQueue int `env:"COMPILERQUEUE" envDefault:"1"`

	DefaultDir    string `env:"DEFAULT_DIR" envDefault:"./temp"`
	MaximumOutput int    `env:"MAXIMUM_OUTPUT" envDefault:"5700"`
	MaximumDelay  int    `env:"MAXIMUM_DELAY" envDefault:"3"`

	RedisAddr     string `env:"REDIS_ADDR" envDefault:"127.0.0.1:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`
	RedisMaster   string `env:"REDIS_MASTER"`

	DefaultProvider string `env:"DEFAULT_PROVIDER" envDefault:"nil"`
	MaxRetryTime    int    `env:"MAXRETRYTIME" envDefault:"1"`
	WenxinApiKey    string `env:"WENXIN_API_KEY" envDefault:"nil"`
	WenxinSecretKey string `env:"WENXIN_SECRET_KEY" envDefault:"nil"`
	TongyiApiKey    string `env:"TONGYI_API_KEY" envDefault:"nil"`

	LoggerLevel string `env:"LOGGERLEVEL" envDefault:"DEBUG"`
	SuperUser   uint64 `env:"SUPER_USER" envDefault:"1581822568"`
}

func init() {
	logrus.Println(os.Getwd())
	if err := godotenv.Load("./config/.env.config"); err != nil {
		logrus.Errorln("Can not read env from file system, please check the right this program owned.")
	}

	EnvCfg = envConfig{}

	if err := env.Parse(&EnvCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}
}
