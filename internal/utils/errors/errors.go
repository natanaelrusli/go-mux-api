package errors

var (
	CodeClientError = 400
	CodeServerError = 500
)

var (
	MsgServerError               = "server error"
	MsgClientBadFormattedRequest = "bad format request"
)

type ErrorWrapper struct {
	Message string `json:"message"` // human readable error
	Code    int    `json:"-"`       // code
	Err     error  `json:"-"`       // original error
}

func (w *ErrorWrapper) Error() string {
	if w.Err != nil {
		return w.Err.Error()
	}

	return w.Message
}

func NewErrorWrapper(code int, msg string, err error) *ErrorWrapper {
	return &ErrorWrapper{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func NewBadFormattedRequest() *ErrorWrapper {
	return NewErrorWrapper(CodeClientError, MsgClientBadFormattedRequest, nil)
}
