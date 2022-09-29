package errs

import (
	"fmt"
	// "encoding/json"
)

type Error struct {
	Code int `json:"errCode"`
	Message string `json:"message"`
	isDbErr bool `json:"-"`
	Err string `json:"error"`
	Module string `json:"-"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s\n Error: %v", e.Code, e.Message, e.Err)
}

// func (e Error) Stack() string {
// 	if e.Err != nil {
// 		return ""
// 	}
// 	return e.Err.Error()
// }