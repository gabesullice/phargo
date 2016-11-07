package types

import "testing"

func TestStringField_Values(t *testing.T) {
	t.Parallel()
	t.Run("Values should return a slice of strings", func(t *testing.T) {
		table := [][]string{
			[]string{},
			[]string{"foo"},
			[]string{"bar", "baz"},
		}
		for _, values := range table {
			f := &StringField{field{"test_field"}, values}
			v, _ := f.Values()
			if res, ok := v.([]string); !ok {
				t.Errorf("Expected slice of strings. Got: %T", res)
			}
		}
	})
	t.Run("Values should set a slice of strings if passed string arguments", func(t *testing.T) {
		f := &StringField{field{"test_field"}, []string{"foo", "bar", "baz"}}
		set := []string{"dog", "cat"}
		v, _ := f.Values(set[0], set[1])
		res, ok := v.([]string)
		if !ok {
			t.Errorf("Expected slice of strings. Got: %T", res)
		}
		for i := 0; i < len(set); i++ {
			if res[i] != set[i] {
				t.Errorf("Expected values to match. Expected: %s Got: %s", set[i], res[i])
			}
		}
	})
	t.Run("Values should return an error if passed non-string values", func(t *testing.T) {
		sets := [][]interface{}{
			[]interface{}{"dog", "cat", 0},
			[]interface{}{"dog", 0, 1},
			[]interface{}{true, "dog", "cat"},
		}
		for _, set := range sets {
			f := &StringField{field{"test_field"}, []string{"foo", "bar", "baz"}}
			if _, err := f.Values(set[0], set[1], set[2]); err == nil {
				t.Error("Expected error, got none.")
			}
		}
	})
}

func TestStrings(t *testing.T) {
	t.Run("Strings should return a slice of strings from a StringField", func(t *testing.T) {
		strings := [][]string{
			[]string{},
			[]string{"foo"},
			[]string{"foo", "bar"},
		}
		for _, expected := range strings {
			f := &StringField{field{"test_field"}, expected}
			actual, _ := Strings(f)
			for i := 0; i < len(expected); i++ {
				if expected[i] != actual[i] {
					t.Errorf("Expected %s to equal %s", expected[i], actual[i])
				}
			}
		}
	})
	t.Run("Strings should return an error if passed a non StringField", func(t *testing.T) {
		_, err := Strings(&BooleanField{field{"test_field"}, []bool{true}})
		if err == nil {
			t.Error("Expected error, go none.")
		}
	})
}

func TestField_Name(t *testing.T) {
	t.Parallel()
	t.Run("Fields should return their names", func(t *testing.T) {
		name := "test_field_name"
		fields := []Field{
			&StringField{field{name}, []string{}},
			&BooleanField{field{name}, []bool{}},
		}
		for _, f := range fields {
			if n := f.Name(); n != name {
				t.Errorf("Expected %T to return its name (%s). Got: %s", f, name, n)
			}
		}
	})
}
