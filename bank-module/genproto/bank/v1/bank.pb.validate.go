// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bank/v1/bank.proto

package bankv1

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

// Validate checks the field values on Bank with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Bank) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bank with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BankMultiError, or nil if none found.
func (m *Bank) ValidateAll() error {
	return m.validate(true)
}

func (m *Bank) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return BankMultiError(errors)
	}

	return nil
}

// BankMultiError is an error wrapping multiple validation errors returned by
// Bank.ValidateAll() if the designated constraints aren't met.
type BankMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BankMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BankMultiError) AllErrors() []error { return m }

// BankValidationError is the validation error returned by Bank.Validate if the
// designated constraints aren't met.
type BankValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BankValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BankValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BankValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BankValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BankValidationError) ErrorName() string { return "BankValidationError" }

// Error satisfies the builtin error interface
func (e BankValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBank.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BankValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BankValidationError{}

// Validate checks the field values on GetBankRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetBankRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBankRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetBankRequestMultiError,
// or nil if none found.
func (m *GetBankRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBankRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetBankRequestMultiError(errors)
	}

	return nil
}

// GetBankRequestMultiError is an error wrapping multiple validation errors
// returned by GetBankRequest.ValidateAll() if the designated constraints
// aren't met.
type GetBankRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBankRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBankRequestMultiError) AllErrors() []error { return m }

// GetBankRequestValidationError is the validation error returned by
// GetBankRequest.Validate if the designated constraints aren't met.
type GetBankRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBankRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBankRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBankRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBankRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBankRequestValidationError) ErrorName() string { return "GetBankRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetBankRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBankRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBankRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBankRequestValidationError{}

// Validate checks the field values on GetBankResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetBankResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBankResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetBankResponseMultiError, or nil if none found.
func (m *GetBankResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBankResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetBankResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetBankResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetBankResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetBankResponseMultiError(errors)
	}

	return nil
}

// GetBankResponseMultiError is an error wrapping multiple validation errors
// returned by GetBankResponse.ValidateAll() if the designated constraints
// aren't met.
type GetBankResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBankResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBankResponseMultiError) AllErrors() []error { return m }

// GetBankResponseValidationError is the validation error returned by
// GetBankResponse.Validate if the designated constraints aren't met.
type GetBankResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBankResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBankResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBankResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBankResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBankResponseValidationError) ErrorName() string { return "GetBankResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetBankResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBankResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBankResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBankResponseValidationError{}
