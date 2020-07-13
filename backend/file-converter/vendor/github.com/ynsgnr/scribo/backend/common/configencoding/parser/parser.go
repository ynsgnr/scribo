package parser

import (
	"reflect"
)

type Parser interface {
	Parse(string) (reflect.Value, error)
}
