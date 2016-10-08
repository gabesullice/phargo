package types

import (
	"github.com/pkg/errors"
)

// StringType provides the structure for representing a string schema.
// More info at: https://spacetelescope.github.io/understanding-json-schema/reference/string.html
type StringType struct {
	Type
	length     uint
	pattern    string
	format     string
	hasLength  bool
	hasPattern bool
	hasFormat  bool
}

// StringTypeDefiner provides type-specific keyword/value pairs for a string.
//
// One should always check the value of `set`. The first return will
// always be a valid value because the zero-value of the type is valid.
// Therefore, it is necessary to ensure it is actually set. For example:
//
//    if val, set := t.Length(); set {
//			foo := val
//		}
//
// Failure to check `set` will result in false schema invalidations.
type StringTypeDefiner interface {
	Length() (len uint, set bool)
	Pattern() (pattern string, set bool)
	Format() (format string, set bool)
}

// NewStringType receives type-specific keyword value pairs for a string and
// appropriately initializes a new StringType.
//
// Possible keys are: "length", "pattern" and "format". Other keys will be
// ignored.
func NewStringType(specifics map[string]interface{}) (StringType, error) {
	var err error
	var t StringType

	t.typeID = "string"

	for k, v := range specifics {
		switch k {
		case "length":
			if v, ok := v.(uint); ok {
				t.length = v
				t.hasLength = true
			} else {
				err = errors.New("Unexpected length type, needs uint.")
			}
		case "pattern":
			if v, ok := v.(string); ok {
				t.pattern = v
				t.hasPattern = true
			} else {
				err = errors.New("Unexpected pattern type, needs string.")
			}
		case "format":
			if v, ok := v.(string); ok {
				t.format = v
				t.hasFormat = true
			} else {
				err = errors.New("Unexpected length type, needs string.")
			}
		}
	}

	return t, err
}

// ToMap returns a normalized definition of a string type.
//
// The returned map will always have a "type" key, but may or may not include
// the "length", "pattern" or "format" keys.
func (t StringType) ToMap() map[string]interface{} {
	def := map[string]interface{}{"type": t.typeID}

	if t.hasLength {
		def["length"] = t.length
	}

	if t.hasPattern {
		def["pattern"] = t.pattern
	}

	if t.hasFormat {
		def["format"] = t.format
	}

	return def
}

// Length returns the string length if one exists
func (t StringType) Length() (len uint, set bool) {
	return t.length, t.hasLength
}

// Pattern returns the pattern if one exists
func (t StringType) Pattern() (pattern string, set bool) {
	return t.pattern, t.hasPattern
}

// Format returns the format if one exists
func (t StringType) Format() (format string, set bool) {
	return t.format, t.hasFormat
}
