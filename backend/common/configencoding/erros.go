package configencoding

import "errors"

var (
	ParserNotFound       = errors.New("parser not found")
	RequiredValueMissing = errors.New("required value missing")
	RequiredPointer      = errors.New("pointer to object is required")
)
