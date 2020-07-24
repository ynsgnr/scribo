package configencoding

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/ynsgnr/scribo/backend/common/configencoding/parser"
	"github.com/ynsgnr/scribo/backend/common/configencoding/source"
)

const (
	DefaultTag  = "default"
	ValidateTag = "validate"

	RequiredValue = "required"
)

func Set(cfg interface{}) error {
	sources := map[string]source.Source{
		source.EnvTag: source.Environment{},
	}
	var b bool
	var s string
	var i int
	var d time.Duration
	parsers := map[reflect.Type]parser.Parser{
		reflect.TypeOf(b): parser.BoolParser{},
		reflect.TypeOf(s): parser.StringParser{},
		reflect.TypeOf(i): parser.IntParser{},
		reflect.TypeOf(d): parser.DurationParser{},
	}
	return SetWithCustomSources(cfg, sources, parsers)
}

func SetWithCustomSources(cfg interface{}, sources map[string]source.Source, parsers map[reflect.Type]parser.Parser) error {
	value := reflect.ValueOf(cfg)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return RequiredPointer
	}
	destination := value.Elem()
	for i := 0; i < destination.NumField(); i++ {
		field := destination.Type().Field(i)
		for tagID, source := range sources {
			if tag, ok := field.Tag.Lookup(tagID); ok {
				value, ok := source.GetValue(tag)
				if !ok {
					if tagValue, ok := destination.Type().Field(i).Tag.Lookup(DefaultTag); ok {
						value = tagValue
					} else if tagValue, ok := destination.Type().Field(i).Tag.Lookup(ValidateTag); ok && tagValue == RequiredValue {
						return errors.Wrapf(RequiredValueMissing, "for field %s with type %s (%s:%s)", field.Name, field.Type.String(), tagID, tag)
					}
				}
				parser, ok := parsers[field.Type]
				if !ok {
					return errors.Wrapf(ParserNotFound, "for field %s with type %s (%s:%s)", field.Name, field.Type.String(), tagID, tag)
				}
				parsedValue, err := parser.Parse(value)
				if err != nil {
					return errors.Wrap(err, "can not parse")
				}
				destination.Field(i).Set(parsedValue)
			}
		}
	}
	return nil
}
