package common

import (
	"fmt"
	"github.com/rozturac/cerror"
	"reflect"
)

func UnExpectedCommand(expectedCommand interface{}) cerror.Error {
	return cerror.New(
		cerror.ApplicationError,
		fmt.Sprintf("Can not cast command to '%s' model", reflect.TypeOf(expectedCommand).Name()),
	)
}

// NullOrEmptyReferenceError Create an instance of Null or Empty Reference Error with object name
func NullOrEmptyReferenceError(fieldName string) cerror.Error {
	return cerror.New(cerror.ApplicationError, fmt.Sprintf("%s is null or empty!", fieldName))
}

// InvalidValueError Create an instance of Invalid Value Error with message
func InvalidValueError(message string) cerror.Error {
	return cerror.New(cerror.ApplicationError, message)
}
