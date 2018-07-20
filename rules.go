package main

import (
	"errors"
	"fmt"
	"strings"
)

func MissingCommandPrefix(cmd string) Validator {
	if len(cmd) == 0 {
		return FailedValidator(errors.New("you are missing the command prefix"), Critical)
	}
	return AlwaysValid()
}

func OnlyTwoCommandPrefix(cmd string) Validator {
	if len(strings.Split(cmd, " ")) > 2 {
		return FailedValidator(fmt.Errorf("cmd [%v] has more than two elements", cmd), Warning)
	}
	return AlwaysValid()
}

func MissingDepsCommand(cmd string) Validator {
	if len(cmd) == 0 {
		return FailedValidator(errors.New("you are missing the deps command"), Critical)
	}
	return AlwaysValid()
}
