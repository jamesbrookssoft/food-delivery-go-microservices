package httpErrors

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StackTraceWithErrors(t *testing.T) {
	err := errors.WithMessage(errors.New("handling bad request"), "this is a bad-request")
	err = errors.Wrap(err, "outer error message")

	res := ErrorsWithStack(err)
	fmt.Println(res)
}

func Test_StackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("handling bad request"), "this is a bad-request")
	err = errors.Wrap(err, "outer error message")

	res := StackTrace(err)
	fmt.Println(res)
}

func Test_RootStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("handling bad request"), "this is a bad-request")
	err = errors.Wrap(err, "outer error message")

	res := RootStackTrace(err)
	fmt.Println(res)
}

func Test_All_Level_Stack_Trace(t *testing.T) {
	err := errors.Wrap(errors.New("handling bad request"), "this is a bad-request")
	err = errors.Wrap(err, "outer error message")

	res := ErrorsWithStack(err)
	fmt.Println(res)
}

func Test_Errors_Without_Stack_Trace(t *testing.T) {
	err := errors.Wrap(errors.New("handling bad request"), "this is a bad-request")
	err = errors.Wrap(err, "outer error message")

	res := ErrorsWithoutTrace(err, true)
	fmt.Println(res)
	assert.Contains(t, res, "outer error message\nthis is a bad-request\nhandling bad request")

	res = ErrorsWithoutTrace(err, false)
	fmt.Println(res)
	assert.Contains(t, res, "outer error message: this is a bad-request: handling bad request")
}
