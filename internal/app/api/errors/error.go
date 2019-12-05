package errors

type (
	//ErrorUnit is the smallest unit of service
	ErrorUnit struct {
		Code    int    `mapstruct:"code"`
		Message string `mapstruct:"message"`
	}
	//AppErrors the struct hold all error message and error information in this service, for return to a other servic and web UI
	AppErrors struct {
		ExternalError OpenAPI      `mapstruct:"external_error"`
		InternalError InternalCall `mapstruct:"internal_error"`
	}
	//OpenAPI to return common code and error message to a external service outside our infrastructure
	OpenAPI struct {
		Authent AuthenticationError `mapstruct:"authent"`
		Request RequestError        `mapstruct:"request"`
	}
	//InternalCall is a struct to hold respons message to internal service when those service call eachother help our developer to know what happen when service is called
	InternalCall struct {
	}
	//AuthenticationError to tell with service what happen with authentication issuse e.g invalid user....
	AuthenticationError struct {
		ExpriteSession   ErrorUnit `mapstruct:"exprite_session"`
		PermistionDenied ErrorUnit `mapstruct:"permistion_denied"`
	}
	//RequestError to tell with service what happen with its request
	RequestError struct {
		RequestInvalid ErrorUnit `mapstruct:"request_invalid"`
	}
)
