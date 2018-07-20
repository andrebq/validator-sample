package main

import (
	"bytes"
	"fmt"
)

type (
	Violation interface {
		error

		Level() Level
	}

	violation struct {
		err error

		level Level
	}

	// Validator is called to check if a violation happens for a given object
	Validator interface {
		Validate() Violation
	}

	validatorList []Validator

	errorList []error

	Level int
)

const (
	Critical = Level(-1)
	Warning  = Level(-2)
	noError  = Level(0)
)

func (l Level) ExitCode() int {
	return int(l)
}

func (l Level) String() string {
	switch l {
	case Critical:
		return "CRITICAL"
	case Warning:
		return "WARNING"
	}
	return ""
}

func (errs errorList) Error() string {
	buf := bytes.Buffer{}
	for i, e := range errs {
		if i != 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintf(&buf, "%v", e)
	}
	return buf.String()
}

func (v violation) Error() string {
	return v.err.Error()
}

// Implements a violation that returns itself if validated, this allows
// for some simple validations to be easly implemented
func (v violation) Validate() Violation {
	if v.err == nil {
		return nil
	}
	return v
}

func (v violation) Level() Level {
	return v.level
}

func CombineValidation(items ...Validator) Validator {
	return validatorList(items)
}

// FailedValidator will always return the same violation
func FailedValidator(err error, level Level) Validator {
	return violation{err: err, level: level}
}

// AlwaysValid returns a validator that is always valid
func AlwaysValid() Validator {
	return violation{}
}

func (vl validatorList) Validate() Violation {
	if len(vl) == 0 {
		return nil
	}
	// create a validation with the worst level of all
	// and the error is the list of all errors
	var errs errorList
	ret := violation{
		level: noError, // every other level will be less than this one
	}
	for _, validation := range vl {
		v := validation.Validate()
		if v == nil {
			continue
		}
		if v.Level() < ret.level {
			ret.level = v.Level()
		}
		errs = append(errs, v)
	}
	ret.err = errs
	if ret.level == noError {
		return nil
	}
	return ret
}
