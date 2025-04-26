package shared

import "net/http"

type AppError struct {
	Code    int    `json:"code"`    // e.g., 400
	Message string `json:"message"` // e.g., "invalid input"
	Err     error  `json:"-"`       // actual Go error, for logs/debug
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func GetCode(err error) int {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code
	}
	return http.StatusInternalServerError
}
