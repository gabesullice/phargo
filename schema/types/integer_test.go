package types

import (
	"testing"
)

func TestNewNumberType(t *testing.T) {
	types := []map[string]interface{}{
		{
			"minimum":          0,
			"exclusiveMinimum": false,
			"maximum":          10.0,
			"exclusiveMaximum": true,
			"multipleOf":       2,
		},
	}

	for _, specifics := range types {
		st, err := NewNumberType(specifics)
		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		len, _ := st.Minimum()
		if specifics["length"] != len {
			t.Errorf("Expected length %d, got %d", specifics["length"], len)
		}

		pattern, _ := st.Pattern()
		if specifics["pattern"] != pattern {
			t.Errorf("Expected pattern %s, got %s", specifics["pattern"], pattern)
		}

		format, _ := st.Format()
		if specifics["format"] != format {
			t.Errorf("Expected format %s, got %s", specifics["format"], format)
		}
	}
}
