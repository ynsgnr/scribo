package parser

import (
	"reflect"
	"time"
)

type DurationParser struct{}

func (DurationParser) Parse(value string) (reflect.Value, error) {
	d, err := time.ParseDuration(value)
	return reflect.ValueOf(d), err
}
