package common

import (
	"fmt"
	"github.com/rozturac/cerror"
	"net/http"
	"reflect"
)

func UnExpectedCommand(from string, expectedCommand interface{}) cerror.Error {
	return cerror.New(
		cerror.ApplicationError,
		fmt.Sprintf("%s can not cast to '%s' model", from, reflect.TypeOf(expectedCommand).Name()),
	)
}

// NullOrEmptyArgumentError Create an instance of Null or Empty Argument Error with object name
func NullOrEmptyArgumentError(fieldName string) cerror.Error {
	return cerror.NewWithHttpStatusCode(cerror.ApplicationError, fmt.Sprintf("%s argument is null or empty!", fieldName), http.StatusBadRequest)
}

// NullOrEmptyReferenceError Create an instance of Null or Empty Reference Error with object name
func NullOrEmptyReferenceError(fieldName string) cerror.Error {
	return cerror.New(cerror.ApplicationError, fmt.Sprintf("%s is null or empty!", fieldName))
}

// InvalidValueError Create an instance of Invalid Value Error with message
func InvalidValueError(message string) cerror.Error {
	return cerror.New(cerror.ApplicationError, message)
}
