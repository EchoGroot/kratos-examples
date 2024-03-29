// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package adminv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNameRepeat(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_NAME_REPEAT.String() && e.Code == 409
}

func ErrorUserNameRepeat(format string, args ...interface{}) *errors.Error {
	return errors.New(409, UserServiceErrorReason_USER_NAME_REPEAT.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserServiceErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserServiceErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}
