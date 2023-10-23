package controllers

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func DataResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

func ErrorResponse(err error) *Response {
	return &Response{
		Message: err.Error(),
	}
}

func TextResponse(msg string) *Response {
	return &Response{
		Message: msg,
	}
}
