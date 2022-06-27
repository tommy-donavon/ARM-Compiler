package compiler

import (
	"fmt"
	"strconv"
)

type (
	iCompiler interface {
		ToString() string
	}
)

func BuildCompiler(flag string, input string) (iCompiler, error) {
	switch flag {
	case "-Blink":
		i, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		return newBlinkCompiler(i - 1)
	}
	return nil, fmt.Errorf("%s is not a supported program", flag)
}
