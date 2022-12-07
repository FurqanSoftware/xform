package pipe

import (
	"reflect"
	"strings"
)

type StepError struct {
	Step Step
	err  error
}

func (e StepError) Error() string {
	return strings.ToLower(reflect.TypeOf(e.Step).Name()) + ": " + e.err.Error()
}

func (e StepError) Unwrap() error {
	return e.err
}
