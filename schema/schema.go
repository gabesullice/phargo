package schema

import (
	"encoding/json"
)

// Schema represents a JSON Schema.
// More info at: http://json-schema.org/documentation.html
type Schema struct {
	Type                 string                      `json:type`
	Id                   string                      `json:id`
	Properties           map[string]Schema           `json:properties`
	PatternProperties    map[string]Schema           `json:patternProperties`
	Dependencies         map[string]SchemaDependency `json:dependencies` // This is a placeholder and is likely to change.
	RequiredProperties   []string                    `json:requiredProperties`
	AdditionalProperties bool                        `json:additionalProperties`
	MinProperties        uint                        `json:minProperties`
	MaxProperties        uint                        `json:maxProperties`
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
