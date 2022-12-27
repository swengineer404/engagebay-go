package engagebay

import "fmt"

type APIError struct {
	Message string `json:"error"`
	Code    int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api_error[%d]: %s", e.Code, e.Message)
}
