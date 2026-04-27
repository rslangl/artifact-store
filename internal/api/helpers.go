package api

func NewError(code int, message string) Error {
	return Error{
		Code: &code,
		Message: &message,
	}
}
