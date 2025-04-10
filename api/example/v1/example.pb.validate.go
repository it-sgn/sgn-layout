// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: example/v1/example.proto

package v1

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

// Validate checks the field values on UpdateExampleReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateExampleReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateExampleReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateExampleReqMultiError, or nil if none found.
func (m *UpdateExampleReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateExampleReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateExampleReqValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 60 {
		err := UpdateExampleReqValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 60 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Status

	if len(errors) > 0 {
		return UpdateExampleReqMultiError(errors)
	}

	return nil
}

// UpdateExampleReqMultiError is an error wrapping multiple validation errors
// returned by UpdateExampleReq.ValidateAll() if the designated constraints
// aren't met.
type UpdateExampleReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateExampleReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateExampleReqMultiError) AllErrors() []error { return m }

// UpdateExampleReqValidationError is the validation error returned by
// UpdateExampleReq.Validate if the designated constraints aren't met.
type UpdateExampleReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateExampleReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateExampleReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateExampleReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateExampleReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateExampleReqValidationError) ErrorName() string { return "UpdateExampleReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateExampleReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateExampleReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateExampleReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateExampleReqValidationError{}

// Validate checks the field values on CreateExampleReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateExampleReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateExampleReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateExampleReqMultiError, or nil if none found.
func (m *CreateExampleReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateExampleReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 60 {
		err := CreateExampleReqValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 60 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Status

	if len(errors) > 0 {
		return CreateExampleReqMultiError(errors)
	}

	return nil
}

// CreateExampleReqMultiError is an error wrapping multiple validation errors
// returned by CreateExampleReq.ValidateAll() if the designated constraints
// aren't met.
type CreateExampleReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateExampleReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateExampleReqMultiError) AllErrors() []error { return m }

// CreateExampleReqValidationError is the validation error returned by
// CreateExampleReq.Validate if the designated constraints aren't met.
type CreateExampleReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateExampleReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateExampleReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateExampleReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateExampleReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateExampleReqValidationError) ErrorName() string { return "CreateExampleReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateExampleReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateExampleReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateExampleReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateExampleReqValidationError{}

// Validate checks the field values on Example with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Example) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Example with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ExampleMultiError, or nil if none found.
func (m *Example) ValidateAll() error {
	return m.validate(true)
}

func (m *Example) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Status

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return ExampleMultiError(errors)
	}

	return nil
}

// ExampleMultiError is an error wrapping multiple validation errors returned
// by Example.ValidateAll() if the designated constraints aren't met.
type ExampleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExampleMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExampleMultiError) AllErrors() []error { return m }

// ExampleValidationError is the validation error returned by Example.Validate
// if the designated constraints aren't met.
type ExampleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExampleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExampleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExampleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExampleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExampleValidationError) ErrorName() string { return "ExampleValidationError" }

// Error satisfies the builtin error interface
func (e ExampleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExample.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExampleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExampleValidationError{}

// Validate checks the field values on GetExampleListPageRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExampleListPageRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExampleListPageRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExampleListPageResMultiError, or nil if none found.
func (m *GetExampleListPageRes) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExampleListPageRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetExampleListPageResValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetExampleListPageResValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetExampleListPageResValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetExampleListPageResMultiError(errors)
	}

	return nil
}

// GetExampleListPageResMultiError is an error wrapping multiple validation
// errors returned by GetExampleListPageRes.ValidateAll() if the designated
// constraints aren't met.
type GetExampleListPageResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExampleListPageResMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExampleListPageResMultiError) AllErrors() []error { return m }

// GetExampleListPageResValidationError is the validation error returned by
// GetExampleListPageRes.Validate if the designated constraints aren't met.
type GetExampleListPageResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExampleListPageResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExampleListPageResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExampleListPageResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExampleListPageResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExampleListPageResValidationError) ErrorName() string {
	return "GetExampleListPageResValidationError"
}

// Error satisfies the builtin error interface
func (e GetExampleListPageResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExampleListPageRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExampleListPageResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExampleListPageResValidationError{}

// Validate checks the field values on GetExampleListReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetExampleListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExampleListReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExampleListReqMultiError, or nil if none found.
func (m *GetExampleListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExampleListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := GetExampleListReqValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() <= 0 {
		err := GetExampleListReqValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetExampleListReqValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetExampleListReqValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetExampleListReqValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetIsDeleted()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetExampleListReqValidationError{
					field:  "IsDeleted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetExampleListReqValidationError{
					field:  "IsDeleted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIsDeleted()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetExampleListReqValidationError{
				field:  "IsDeleted",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedAtStart

	// no validation rules for CreatedAtEnd

	if len(errors) > 0 {
		return GetExampleListReqMultiError(errors)
	}

	return nil
}

// GetExampleListReqMultiError is an error wrapping multiple validation errors
// returned by GetExampleListReq.ValidateAll() if the designated constraints
// aren't met.
type GetExampleListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExampleListReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExampleListReqMultiError) AllErrors() []error { return m }

// GetExampleListReqValidationError is the validation error returned by
// GetExampleListReq.Validate if the designated constraints aren't met.
type GetExampleListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExampleListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExampleListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExampleListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExampleListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExampleListReqValidationError) ErrorName() string {
	return "GetExampleListReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetExampleListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExampleListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExampleListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExampleListReqValidationError{}

// Validate checks the field values on ExampleIdReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExampleIdReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExampleIdReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExampleIdReqMultiError, or
// nil if none found.
func (m *ExampleIdReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ExampleIdReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ExampleIdReqMultiError(errors)
	}

	return nil
}

// ExampleIdReqMultiError is an error wrapping multiple validation errors
// returned by ExampleIdReq.ValidateAll() if the designated constraints aren't met.
type ExampleIdReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExampleIdReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExampleIdReqMultiError) AllErrors() []error { return m }

// ExampleIdReqValidationError is the validation error returned by
// ExampleIdReq.Validate if the designated constraints aren't met.
type ExampleIdReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExampleIdReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExampleIdReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExampleIdReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExampleIdReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExampleIdReqValidationError) ErrorName() string { return "ExampleIdReqValidationError" }

// Error satisfies the builtin error interface
func (e ExampleIdReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExampleIdReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExampleIdReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExampleIdReqValidationError{}
