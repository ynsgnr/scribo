package parser

import (
	"reflect"
	"strconv"
)

type BoolParser struct{}

func (BoolParser) Parse(value string) (reflect.Value, error) {
	if value == "" {
		return reflect.ValueOf(false), nil
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(v), nil
}
