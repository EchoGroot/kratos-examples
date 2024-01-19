// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/conf/conf.proto

package conf

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

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Bootstrap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bootstrap with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BootstrapMultiError, or nil
// if none found.
func (m *Bootstrap) ValidateAll() error {
	return m.validate(true)
}

func (m *Bootstrap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetApp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetServer() == nil {
		err := BootstrapValidationError{
			field:  "Server",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetServer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Server",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Server",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Server",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetData() == nil {
		err := BootstrapValidationError{
			field:  "Data",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BootstrapMultiError(errors)
	}

	return nil
}

// BootstrapMultiError is an error wrapping multiple validation errors returned
// by Bootstrap.ValidateAll() if the designated constraints aren't met.
type BootstrapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BootstrapMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BootstrapMultiError) AllErrors() []error { return m }

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}

// Validate checks the field values on App with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *App) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on App with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AppMultiError, or nil if none found.
func (m *App) ValidateAll() error {
	return m.validate(true)
}

func (m *App) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AppMultiError(errors)
	}

	return nil
}

// AppMultiError is an error wrapping multiple validation errors returned by
// App.ValidateAll() if the designated constraints aren't met.
type AppMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppMultiError) AllErrors() []error { return m }

// AppValidationError is the validation error returned by App.Validate if the
// designated constraints aren't met.
type AppValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppValidationError) ErrorName() string { return "AppValidationError" }

// Error satisfies the builtin error interface
func (e AppValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppValidationError{}

// Validate checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServerMultiError, or nil if none found.
func (m *Server) ValidateAll() error {
	return m.validate(true)
}

func (m *Server) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetHttp() == nil {
		err := ServerValidationError{
			field:  "Http",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetHttp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Http",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetGrpc() == nil {
		err := ServerValidationError{
			field:  "Grpc",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGrpc()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Grpc",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServerValidationError{
					field:  "Grpc",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpc()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Grpc",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServerMultiError(errors)
	}

	return nil
}

// ServerMultiError is an error wrapping multiple validation errors returned by
// Server.ValidateAll() if the designated constraints aren't met.
type ServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerMultiError) AllErrors() []error { return m }

// ServerValidationError is the validation error returned by Server.Validate if
// the designated constraints aren't met.
type ServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerValidationError) ErrorName() string { return "ServerValidationError" }

// Error satisfies the builtin error interface
func (e ServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Service) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServiceMultiError, or nil if none found.
func (m *Service) ValidateAll() error {
	return m.validate(true)
}

func (m *Service) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ServiceMultiError(errors)
	}

	return nil
}

// ServiceMultiError is an error wrapping multiple validation errors returned
// by Service.ValidateAll() if the designated constraints aren't met.
type ServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceMultiError) AllErrors() []error { return m }

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DataMultiError, or nil if none found.
func (m *Data) ValidateAll() error {
	return m.validate(true)
}

func (m *Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPostgres() == nil {
		err := DataValidationError{
			field:  "Postgres",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPostgres()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Postgres",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataValidationError{
					field:  "Postgres",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPostgres()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Postgres",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DataMultiError(errors)
	}

	return nil
}

// DataMultiError is an error wrapping multiple validation errors returned by
// Data.ValidateAll() if the designated constraints aren't met.
type DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMultiError) AllErrors() []error { return m }

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on Server_HTTP with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server_HTTP) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_HTTP with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Server_HTTPMultiError, or
// nil if none found.
func (m *Server_HTTP) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_HTTP) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	if utf8.RuneCountInString(m.GetAddr()) < 1 {
		err := Server_HTTPValidationError{
			field:  "Addr",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimeout() == nil {
		err := Server_HTTPValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Server_HTTPMultiError(errors)
	}

	return nil
}

// Server_HTTPMultiError is an error wrapping multiple validation errors
// returned by Server_HTTP.ValidateAll() if the designated constraints aren't met.
type Server_HTTPMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_HTTPMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_HTTPMultiError) AllErrors() []error { return m }

// Server_HTTPValidationError is the validation error returned by
// Server_HTTP.Validate if the designated constraints aren't met.
type Server_HTTPValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_HTTPValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_HTTPValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_HTTPValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_HTTPValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_HTTPValidationError) ErrorName() string { return "Server_HTTPValidationError" }

// Error satisfies the builtin error interface
func (e Server_HTTPValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_HTTP.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_HTTPValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_HTTPValidationError{}

// Validate checks the field values on Server_GRPC with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Server_GRPC) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Server_GRPC with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Server_GRPCMultiError, or
// nil if none found.
func (m *Server_GRPC) ValidateAll() error {
	return m.validate(true)
}

func (m *Server_GRPC) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	if utf8.RuneCountInString(m.GetAddr()) < 1 {
		err := Server_GRPCValidationError{
			field:  "Addr",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimeout() == nil {
		err := Server_GRPCValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Server_GRPCMultiError(errors)
	}

	return nil
}

// Server_GRPCMultiError is an error wrapping multiple validation errors
// returned by Server_GRPC.ValidateAll() if the designated constraints aren't met.
type Server_GRPCMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Server_GRPCMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Server_GRPCMultiError) AllErrors() []error { return m }

// Server_GRPCValidationError is the validation error returned by
// Server_GRPC.Validate if the designated constraints aren't met.
type Server_GRPCValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_GRPCValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_GRPCValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_GRPCValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_GRPCValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_GRPCValidationError) ErrorName() string { return "Server_GRPCValidationError" }

// Error satisfies the builtin error interface
func (e Server_GRPCValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_GRPC.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_GRPCValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_GRPCValidationError{}

// Validate checks the field values on Data_Postgres with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Data_Postgres) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Postgres with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Data_PostgresMultiError, or
// nil if none found.
func (m *Data_Postgres) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Postgres) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDsn()) < 1 {
		err := Data_PostgresValidationError{
			field:  "Dsn",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPool() == nil {
		err := Data_PostgresValidationError{
			field:  "Pool",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPool()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_PostgresValidationError{
					field:  "Pool",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_PostgresValidationError{
					field:  "Pool",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPool()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_PostgresValidationError{
				field:  "Pool",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetInit() == nil {
		err := Data_PostgresValidationError{
			field:  "Init",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetInit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Data_PostgresValidationError{
					field:  "Init",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Data_PostgresValidationError{
					field:  "Init",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Data_PostgresValidationError{
				field:  "Init",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Data_PostgresMultiError(errors)
	}

	return nil
}

// Data_PostgresMultiError is an error wrapping multiple validation errors
// returned by Data_Postgres.ValidateAll() if the designated constraints
// aren't met.
type Data_PostgresMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_PostgresMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_PostgresMultiError) AllErrors() []error { return m }

// Data_PostgresValidationError is the validation error returned by
// Data_Postgres.Validate if the designated constraints aren't met.
type Data_PostgresValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_PostgresValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_PostgresValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_PostgresValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_PostgresValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_PostgresValidationError) ErrorName() string { return "Data_PostgresValidationError" }

// Error satisfies the builtin error interface
func (e Data_PostgresValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Postgres.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_PostgresValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_PostgresValidationError{}

// Validate checks the field values on Data_Postgres_Pool with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Data_Postgres_Pool) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Postgres_Pool with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Data_Postgres_PoolMultiError, or nil if none found.
func (m *Data_Postgres_Pool) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Postgres_Pool) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMaxIdleConns() <= 0 {
		err := Data_Postgres_PoolValidationError{
			field:  "MaxIdleConns",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMaxOpenConns() <= 0 {
		err := Data_Postgres_PoolValidationError{
			field:  "MaxOpenConns",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetConnMaxLifetime() == nil {
		err := Data_Postgres_PoolValidationError{
			field:  "ConnMaxLifetime",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Data_Postgres_PoolMultiError(errors)
	}

	return nil
}

// Data_Postgres_PoolMultiError is an error wrapping multiple validation errors
// returned by Data_Postgres_Pool.ValidateAll() if the designated constraints
// aren't met.
type Data_Postgres_PoolMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_Postgres_PoolMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_Postgres_PoolMultiError) AllErrors() []error { return m }

// Data_Postgres_PoolValidationError is the validation error returned by
// Data_Postgres_Pool.Validate if the designated constraints aren't met.
type Data_Postgres_PoolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_Postgres_PoolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_Postgres_PoolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_Postgres_PoolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_Postgres_PoolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_Postgres_PoolValidationError) ErrorName() string {
	return "Data_Postgres_PoolValidationError"
}

// Error satisfies the builtin error interface
func (e Data_Postgres_PoolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Postgres_Pool.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_Postgres_PoolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_Postgres_PoolValidationError{}

// Validate checks the field values on Data_Postgres_Init with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Data_Postgres_Init) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data_Postgres_Init with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Data_Postgres_InitMultiError, or nil if none found.
func (m *Data_Postgres_Init) ValidateAll() error {
	return m.validate(true)
}

func (m *Data_Postgres_Init) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTimeout() == nil {
		err := Data_Postgres_InitValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Data_Postgres_InitMultiError(errors)
	}

	return nil
}

// Data_Postgres_InitMultiError is an error wrapping multiple validation errors
// returned by Data_Postgres_Init.ValidateAll() if the designated constraints
// aren't met.
type Data_Postgres_InitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Data_Postgres_InitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Data_Postgres_InitMultiError) AllErrors() []error { return m }

// Data_Postgres_InitValidationError is the validation error returned by
// Data_Postgres_Init.Validate if the designated constraints aren't met.
type Data_Postgres_InitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_Postgres_InitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_Postgres_InitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_Postgres_InitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_Postgres_InitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_Postgres_InitValidationError) ErrorName() string {
	return "Data_Postgres_InitValidationError"
}

// Error satisfies the builtin error interface
func (e Data_Postgres_InitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Postgres_Init.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_Postgres_InitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_Postgres_InitValidationError{}
