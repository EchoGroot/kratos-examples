// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/user-manage/v1/livez.proto

package adminv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LivezRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LivezRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LivezRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LivezRequestMultiError, or
// nil if none found.
func (m *LivezRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LivezRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LivezRequestMultiError(errors)
	}

	return nil
}

// LivezRequestMultiError is an error wrapping multiple validation errors
// returned by LivezRequest.ValidateAll() if the designated constraints aren't met.
type LivezRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LivezRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LivezRequestMultiError) AllErrors() []error { return m }

// LivezRequestValidationError is the validation error returned by
// LivezRequest.Validate if the designated constraints aren't met.
type LivezRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LivezRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LivezRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LivezRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LivezRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LivezRequestValidationError) ErrorName() string { return "LivezRequestValidationError" }

// Error satisfies the builtin error interface
func (e LivezRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLivezRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LivezRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LivezRequestValidationError{}

// Validate checks the field values on LivezResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LivezResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LivezResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LivezResponseMultiError, or
// nil if none found.
func (m *LivezResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LivezResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LivezResponseMultiError(errors)
	}

	return nil
}

// LivezResponseMultiError is an error wrapping multiple validation errors
// returned by LivezResponse.ValidateAll() if the designated constraints
// aren't met.
type LivezResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LivezResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LivezResponseMultiError) AllErrors() []error { return m }

// LivezResponseValidationError is the validation error returned by
// LivezResponse.Validate if the designated constraints aren't met.
type LivezResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LivezResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LivezResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LivezResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LivezResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LivezResponseValidationError) ErrorName() string { return "LivezResponseValidationError" }

// Error satisfies the builtin error interface
func (e LivezResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLivezResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LivezResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LivezResponseValidationError{}