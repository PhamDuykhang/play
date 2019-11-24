package conf

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	//DevStage the env when we develop the service in individual machine
	DevStage Stage = "dev"
	//ProductionStage when we deploy the service in real sever
	ProductionStage Stage = "prod"
)

type (
	//Stage the abstract type to hole current servic's stage
	Stage string
	// Config struct hold the config structure of all servicer
	// All your config should be a child of config struct to making service is flexable than
	Config struct {
		Stage          Stage
		HTTPServer     HTTPServer     `mapstructure:"http_server"`
		Infrastructure Infrastructure `mapstructure:"infrastructure"`
	}
	// HTTPServer hold config service's http gateway e.g HTTP request time out
	HTTPServer struct {
		ReadTimeout       time.Duration `mapstructure:"read_timeout"`
		WriteTimeout      time.Duration `mapstructure:"write_timeout"`
		ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
		ShutdownTimeout   time.Duration `mapstructure:"shutdown_timeout"`
	}
	//Infrastructure the struct being use to hold all connection to out side infrastructure
	//e.g database,message queue, cloud computing
	Infrastructure struct {
		MongoDB MongoDB `mapstructure:"mongo_db"`
	}
	//MongoDB the struct hold all information to connect to mongo db server
	MongoDB struct {
		Address      []string `mapstructure:"address"`
		Username     string   `mapstructure:"user_name"`
		Password     string   `mapstructure:"pass_word"`
		DatabaseName string   `mapstructure:"database_name"`
		Enable       bool     `mapstructure:"enable"`
	}
)

//LoadConfig this func run and read config file in config folder
func LoadConfig(state, confDir *string) (*Config, error) {
	var envStage Stage
	switch *state {
	case "local", "localhost", "l":
		envStage = DevStage
		break
	case "production", "prob", "p":
		envStage = ProductionStage
	default:
		envStage = DevStage
	}
	var c Config
	vp := viper.New()
	vp.AddConfigPath(*confDir)
	vp.SetConfigName(fmt.Sprintf("config.%s", envStage))
	if err := vp.ReadInConfig(); err != nil {
		return &c, err
	}
	if err := vp.Unmarshal(&c); err != nil {
		return &c, err
	}
	return &c, nil
}
