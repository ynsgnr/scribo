package parser

import (
	"reflect"
	"strconv"
)

type IntParser struct{}

func (IntParser) Parse(value string) (reflect.Value, error) {
	if value == "" {
		return reflect.ValueOf(0), nil
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(v), nil
}
