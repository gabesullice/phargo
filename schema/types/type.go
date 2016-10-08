package types

// Type allows one to provide a JSON Schema type.
// More info at: https://spacetelescope.github.io/understanding-json-schema/reference/type.html
type Type struct {
	typeID string
}

// TypeDefiner can provide their schema type and also normalize themselves into
// a map based definition.
type TypeDefiner interface {
	Type() string
	TypeMapper
}

// TypeMapper is used to normalize any schema type into a map.
type TypeMapper interface {
	ToMap() map[string]interface{}
}

func (t Type) Type() string {
	return t.typeID
}
