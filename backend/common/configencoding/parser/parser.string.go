package parser

import "reflect"

type StringParser struct{}

func (StringParser) Parse(value string) (reflect.Value, error) {
	return reflect.ValueOf(value), nil
}
