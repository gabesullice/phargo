package types

import "github.com/pkg/errors"

type Field interface {
	Values(v ...interface{}) (interface{}, error)
	Name() string
}

type field struct {
	name string
}

func (f field) Name() string {
	return f.name
}

type StringField struct {
	field
	values []string
}

func (f *StringField) Values(v ...interface{}) (interface{}, error) {
	if len(v) > 0 {
		var values []string
		for _, val := range v {
			str, err := toString(val)
			if err != nil {
				return f.values, err
			}
			values = append(values, str)
		}
		f.values = values
	}
	return f.values, nil
}

func (f *StringField) Strings() []string {
	return f.values
}

func Strings(f Field) ([]string, error) {
	if strf, ok := f.(*StringField); ok {
		return strf.Strings(), nil
	} else {
		return nil, errors.New("Tried to extract string values from non StringField.")
	}
}

func toString(v interface{}) (string, error) {
	if str, ok := v.(string); !ok {
		return "", errors.New("Tried to use a non-string value in a string context.")
	} else {
		return str, nil
	}
}

type BooleanField struct {
	field
	values []bool
}

func (f *BooleanField) Values(v ...interface{}) (interface{}, error) {
	if len(v) > 0 {
		var values []bool
		for _, val := range v {
			b, err := toBool(val)
			if err != nil {
				return f.values, err
			}
			values = append(values, b)
		}
		f.values = values
	}
	return f.values, nil
}

func Booleans(f Field) ([]bool, error) {
	if bf, ok := f.(*BooleanField); ok {
		return bf.values, nil
	} else {
		return nil, errors.New("Tried to extract boolean values from non BooleanField.")
	}
}

func toBool(v interface{}) (bool, error) {
	if b, ok := v.(bool); !ok {
		return false, errors.New("Tried to use a non-bool value in a bool context.")
	} else {
		return b, nil
	}
}
