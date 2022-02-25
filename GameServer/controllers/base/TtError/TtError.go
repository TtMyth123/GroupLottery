package TtError

import (
	"fmt"
)

type TtError struct {
	E  string
	Mp []interface{}
}

func New(text string, a ...interface{}) *TtError {
	aTtError := new(TtError)
	aTtError.E = text
	i := len(a)
	if i > 0 {
		aTtError.Mp = make([]interface{}, 0)
		aTtError.Mp = append(aTtError.Mp, a...)
	}

	return aTtError
}

func (e *TtError) Error() string {
	return fmt.Sprintf(e.E, e.Mp...)
}
