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
		ExternalError OpenAPI      `mapstructure:"external_error"`
		InternalError InternalCall `mapstructure:"internal_error"`
	}
	//OpenAPI to return common code and error message to a external service outside our infrastructure
	OpenAPI struct {
		Authentication AuthenticationError `mapstructure:"authentication"`
		Request        RequestError        `mapstructure:"request"`
	}
	//InternalCall is a struct to hold respons message to internal service when those service call eachother help our developer to know what happen when service is called
	InternalCall struct {
		DatabaseError DataBaseError `mapstructure:"database_error"`
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

//I return internal call error data directly
func (a *AppErrors) I() *InternalCall {
	return &a.InternalError
}

//E return external call error data directly
func (a *AppErrors) E() *OpenAPI {
	return &a.ExternalError
}
