package types

import (
	"testing"
)

func TestNewStringType(t *testing.T) {
	types := []map[string]interface{}{
		{
			"length":  uint(10),
			"pattern": `^\w`,
			"format":  "date-time",
		},
	}

	for _, specifics := range types {
		st, err := NewStringType(specifics)
		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		len, _ := st.Length()
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
