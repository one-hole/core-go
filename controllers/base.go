package controllers

type APIException struct {
	Message string `json:"message"`
}

// 这里实现了 error 接口
func (e *APIException) Error() string {
	return  e.Message
}

func newAPIException(msg string) *APIException {
	return &APIException{
		Message: msg,
	}
}


func ServerError() *APIException {
	return newAPIException("出错了")
}