package postgrest

import "fmt"

type RestError struct {
	Code    string `json:"code"`
	Details string `json:"details"`
	Hint    string `json:"hint"`
	Message string `json:"message"`
}

func (e *RestError) Error() string {
	return fmt.Sprintf("code: %s, details: %s, hint: %s, message: %s", e.Code, e.Details, e.Hint, e.Message)
}
