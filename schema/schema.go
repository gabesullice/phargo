package schema

import (
	"github.com/pkg/errors"
)

// Schema represents a JSON Schema.
// More info at: http://json-schema.org/documentation.html
type Schema struct {
	Type
	Id                   string `json:id`
	TypeSpecifics        map[string]string
	Properties           map[string]Schema           `json:properties`
	PatternProperties    map[string]Schema           `json:patternProperties`
	Dependencies         map[string]SchemaDependency `json:dependencies` // This is a placeholder and is likely to change.
	RequiredProperties   []string                    `json:requiredProperties`
	AdditionalProperties bool                        `json:additionalProperties`
	MinProperties        uint                        `json:minProperties`
	MaxProperties        uint                        `json:maxProperties`
}

// Type allows one to provide a JSON Schema type and type-specific keywords.
// More info at: https://spacetelescope.github.io/understanding-json-schema/reference/type.html
type Type interface {
	Definition() map[string]interface{}
}

// StringType provides the structure for representing a string schema.
// More info at: https://spacetelescope.github.io/understanding-json-schema/reference/string.html
type StringType struct {
	Type       string
	Length     uint
	Pattern    string
	Format     string
	hasLength  bool
	hasPattern bool
	hasFormat  bool
}

func NewStringType(specifics map[string]interface{}) (StringType, error) {
	var err error
	t := StringType{Type: "string"}

	for k, v := range specifics {
		switch k {
		case "length":
			if v, ok := v.(uint); ok {
				t.Length = v
				t.hasLength = true
			} else {
				err = errors.New("Unexpected length type, needs uint.")
			}
		case "pattern":
			if v, ok := v.(string); ok {
				t.Pattern = v
				t.hasPattern = true
			} else {
				err = errors.New("Unexpected pattern type, needs string.")
			}
		case "format":
			if v, ok := v.(string); ok {
				t.Format = v
				t.hasFormat = true
			} else {
				err = errors.New("Unexpected length type, needs string.")
			}
		}
	}

	return t, err
}

func (t StringType) Definition() map[string]interface{} {
	def := map[string]interface{}{"type": t.Type}

	if t.hasLength {
		def["length"] = t.Length
	}

	if t.hasPattern {
		def["pattern"] = t.Pattern
	}

	if t.hasFormat {
		def["format"] = t.Format
	}

	return def
}

// SchemaDependency allows one to specify a more dynamic schema.
//
// A SchemaDependency can be either a simple property requirement or an
// additional schema to validate against. One should call IsSchema() to
// determine if one should next call Schema() or Properties().
type SchemaDependency interface {
	IsSchema() bool
	Properties() ([]string, error)
	Schema() (Schema, error)
}
