package handler

type (
	//Response hold common information to return to FE or user side
	Response struct {
		StatusCode int         `json:"status_code,omitempty"`
		Message    string      `json:"message,omitempty"`
		Data       interface{} `json:"data,omitempty"`
	}
)
