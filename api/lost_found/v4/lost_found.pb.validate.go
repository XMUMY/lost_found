// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/lost_found/v4/lost_found.proto

package v4

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

// Validate checks the field values on GetBriefsReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetBriefsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBriefsReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetBriefsReqMultiError, or
// nil if none found.
func (m *GetBriefsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBriefsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBefore()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetBriefsReqValidationError{
					field:  "Before",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetBriefsReqValidationError{
					field:  "Before",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetBriefsReqValidationError{
				field:  "Before",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetBriefsReqMultiError(errors)
	}

	return nil
}

// GetBriefsReqMultiError is an error wrapping multiple validation errors
// returned by GetBriefsReq.ValidateAll() if the designated constraints aren't met.
type GetBriefsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBriefsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBriefsReqMultiError) AllErrors() []error { return m }

// GetBriefsReqValidationError is the validation error returned by
// GetBriefsReq.Validate if the designated constraints aren't met.
type GetBriefsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBriefsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBriefsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBriefsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBriefsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBriefsReqValidationError) ErrorName() string { return "GetBriefsReqValidationError" }

// Error satisfies the builtin error interface
func (e GetBriefsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBriefsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBriefsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBriefsReqValidationError{}

// Validate checks the field values on GetBriefsResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetBriefsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBriefsResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetBriefsRespMultiError, or
// nil if none found.
func (m *GetBriefsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBriefsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBriefs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetBriefsRespValidationError{
						field:  fmt.Sprintf("Briefs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetBriefsRespValidationError{
						field:  fmt.Sprintf("Briefs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetBriefsRespValidationError{
					field:  fmt.Sprintf("Briefs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetBriefsRespMultiError(errors)
	}

	return nil
}

// GetBriefsRespMultiError is an error wrapping multiple validation errors
// returned by GetBriefsResp.ValidateAll() if the designated constraints
// aren't met.
type GetBriefsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBriefsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBriefsRespMultiError) AllErrors() []error { return m }

// GetBriefsRespValidationError is the validation error returned by
// GetBriefsResp.Validate if the designated constraints aren't met.
type GetBriefsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBriefsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBriefsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBriefsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBriefsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBriefsRespValidationError) ErrorName() string { return "GetBriefsRespValidationError" }

// Error satisfies the builtin error interface
func (e GetBriefsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBriefsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBriefsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBriefsRespValidationError{}

// Validate checks the field values on LostAndFoundBrief with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LostAndFoundBrief) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LostAndFoundBrief with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LostAndFoundBriefMultiError, or nil if none found.
func (m *LostAndFoundBrief) ValidateAll() error {
	return m.validate(true)
}

func (m *LostAndFoundBrief) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Uid

	// no validation rules for Type

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LostAndFoundBriefValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LostAndFoundBriefValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LostAndFoundBriefValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Location

	if len(errors) > 0 {
		return LostAndFoundBriefMultiError(errors)
	}

	return nil
}

// LostAndFoundBriefMultiError is an error wrapping multiple validation errors
// returned by LostAndFoundBrief.ValidateAll() if the designated constraints
// aren't met.
type LostAndFoundBriefMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LostAndFoundBriefMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LostAndFoundBriefMultiError) AllErrors() []error { return m }

// LostAndFoundBriefValidationError is the validation error returned by
// LostAndFoundBrief.Validate if the designated constraints aren't met.
type LostAndFoundBriefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LostAndFoundBriefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LostAndFoundBriefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LostAndFoundBriefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LostAndFoundBriefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LostAndFoundBriefValidationError) ErrorName() string {
	return "LostAndFoundBriefValidationError"
}

// Error satisfies the builtin error interface
func (e LostAndFoundBriefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLostAndFoundBrief.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LostAndFoundBriefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LostAndFoundBriefValidationError{}

// Validate checks the field values on GetDetailReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetDetailReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDetailReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetDetailReqMultiError, or
// nil if none found.
func (m *GetDetailReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDetailReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 24 {
		err := GetDetailReqValidationError{
			field:  "Id",
			reason: "value length must be 24 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return GetDetailReqMultiError(errors)
	}

	return nil
}

// GetDetailReqMultiError is an error wrapping multiple validation errors
// returned by GetDetailReq.ValidateAll() if the designated constraints aren't met.
type GetDetailReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDetailReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDetailReqMultiError) AllErrors() []error { return m }

// GetDetailReqValidationError is the validation error returned by
// GetDetailReq.Validate if the designated constraints aren't met.
type GetDetailReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDetailReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDetailReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDetailReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDetailReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDetailReqValidationError) ErrorName() string { return "GetDetailReqValidationError" }

// Error satisfies the builtin error interface
func (e GetDetailReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDetailReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDetailReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDetailReqValidationError{}

// Validate checks the field values on LostAndFoundDetail with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LostAndFoundDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LostAndFoundDetail with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LostAndFoundDetailMultiError, or nil if none found.
func (m *LostAndFoundDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *LostAndFoundDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Type

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LostAndFoundDetailValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LostAndFoundDetailValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LostAndFoundDetailValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Location

	// no validation rules for Description

	// no validation rules for Contacts

	if len(errors) > 0 {
		return LostAndFoundDetailMultiError(errors)
	}

	return nil
}

// LostAndFoundDetailMultiError is an error wrapping multiple validation errors
// returned by LostAndFoundDetail.ValidateAll() if the designated constraints
// aren't met.
type LostAndFoundDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LostAndFoundDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LostAndFoundDetailMultiError) AllErrors() []error { return m }

// LostAndFoundDetailValidationError is the validation error returned by
// LostAndFoundDetail.Validate if the designated constraints aren't met.
type LostAndFoundDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LostAndFoundDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LostAndFoundDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LostAndFoundDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LostAndFoundDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LostAndFoundDetailValidationError) ErrorName() string {
	return "LostAndFoundDetailValidationError"
}

// Error satisfies the builtin error interface
func (e LostAndFoundDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLostAndFoundDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LostAndFoundDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LostAndFoundDetailValidationError{}

// Validate checks the field values on AddItemReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddItemReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddItemReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddItemReqMultiError, or
// nil if none found.
func (m *AddItemReq) ValidateAll() error {
	return m.validate(true)
}

func (m *AddItemReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if utf8.RuneCountInString(m.GetName()) > 50 {
		err := AddItemReqValidationError{
			field:  "Name",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTime() == nil {
		err := AddItemReqValidationError{
			field:  "Time",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetLocation()) > 100 {
		err := AddItemReqValidationError{
			field:  "Location",
			reason: "value length must be at most 100 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) > 500 {
		err := AddItemReqValidationError{
			field:  "Description",
			reason: "value length must be at most 500 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetContacts()) > 10 {
		err := AddItemReqValidationError{
			field:  "Contacts",
			reason: "value must contain no more than 10 pair(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetContacts()))
		i := 0
		for key := range m.GetContacts() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetContacts()[key]
			_ = val

			if utf8.RuneCountInString(key) > 50 {
				err := AddItemReqValidationError{
					field:  fmt.Sprintf("Contacts[%v]", key),
					reason: "value length must be at most 50 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if utf8.RuneCountInString(val) > 100 {
				err := AddItemReqValidationError{
					field:  fmt.Sprintf("Contacts[%v]", key),
					reason: "value length must be at most 100 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return AddItemReqMultiError(errors)
	}

	return nil
}

// AddItemReqMultiError is an error wrapping multiple validation errors
// returned by AddItemReq.ValidateAll() if the designated constraints aren't met.
type AddItemReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddItemReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddItemReqMultiError) AllErrors() []error { return m }

// AddItemReqValidationError is the validation error returned by
// AddItemReq.Validate if the designated constraints aren't met.
type AddItemReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddItemReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddItemReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddItemReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddItemReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddItemReqValidationError) ErrorName() string { return "AddItemReqValidationError" }

// Error satisfies the builtin error interface
func (e AddItemReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddItemReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddItemReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddItemReqValidationError{}

// Validate checks the field values on DeleteItemReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteItemReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteItemReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteItemReqMultiError, or
// nil if none found.
func (m *DeleteItemReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteItemReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteItemReqMultiError(errors)
	}

	return nil
}

// DeleteItemReqMultiError is an error wrapping multiple validation errors
// returned by DeleteItemReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteItemReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteItemReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteItemReqMultiError) AllErrors() []error { return m }

// DeleteItemReqValidationError is the validation error returned by
// DeleteItemReq.Validate if the designated constraints aren't met.
type DeleteItemReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteItemReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteItemReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteItemReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteItemReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteItemReqValidationError) ErrorName() string { return "DeleteItemReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteItemReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteItemReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteItemReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteItemReqValidationError{}
