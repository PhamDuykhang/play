package errors

import (
	"fmt"

	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/spf13/viper"
)

type (
	//ErrorUnit is the smallest unit of service
	ErrorUnit struct {
		Code    int    `mapstructure:"code"`
		Message string `mapstructure:"message"`
	}
	//AppErrors the struct hold all error message and error information in this service, for return to a other servic and web UI
	AppErrors struct {
		Success        ErrorUnit           `mapstructure:"success"`
		Common         ErrorUnit           `mapstructure:"common"`
		Authentication AuthenticationError `mapstructure:"authentication"`
		Request        RequestError        `mapstructure:"request"`
		DatabaseError  DataBaseError       `mapstructure:"database_error"`
	}
	//AuthenticationError to tell with a orther service what happen with authentication issuse e.g invalid user....
	AuthenticationError struct {
		ExpireSession    ErrorUnit `mapstructure:"expire_session"`
		PermissionDenied ErrorUnit `mapstructure:"permission_denied"`
	}
	//RequestError to tell with service what happen with its request
	RequestError struct {
		RequestInvalid ErrorUnit `mapstructure:"request_invalid"`
	}
	//DataBaseError to tell with a other service what happen with database
	DataBaseError struct {
		Connection ErrorUnit `mapstructure:"connection"`
	}
)

//Init to load data of error message
//Next sprint will be design auto reload config when config change
func Init(state, confDir *string) (*AppErrors, error) {
	var envStage conf.Stage
	switch *state {
	case "local", "localhost", "l":
		envStage = conf.DevStage
		break
	case "production", "prob", "p":
		envStage = conf.ProductionStage
	default:
		envStage = conf.DevStage
	}
	vp := viper.New()
	var e AppErrors
	vp.AddConfigPath(*confDir)
	vp.SetConfigName(fmt.Sprintf("error.%s", envStage))
	if err := vp.ReadInConfig(); err != nil {
		return &e, err
	}
	if err := vp.Unmarshal(&e); err != nil {
		return &e, err
	}
	return &e, nil
}
