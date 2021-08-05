package errors

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(text string) error {
	return &Error{
		Message: text,
		Code:    500,
	}
}

//Tuân thủ interface chuẩn của errors package
func (e *Error) Error() string {
	return e.Message
}

//Đây là những hàm bổ xung để tạo error dễ hơn
func Generic(text string, code int) error {
	return &Error{
		Message: text,
		Code:    code,
	}
}

func UnAuthorized(text string) error {
	return &Error{
		Message: text,
		Code:    401,
	}
}

func BadRequest(text string) error {
	return &Error{
		Message: text,
		Code:    400,
	}
}
func NotFound(text string) error {
	return &Error{
		Message: text,
		Code:    404,
	}
}
