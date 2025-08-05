package figaro

import (
	"reflect"
	"testing"
)

func TestConfiguration(t *testing.T) {
	if err := Init(); err != nil {
		t.Errorf("failed to load configuration")
	}

	var (
		value1 string
		value2 string
		value3 int
		value4 bool
		values []string
		err    error
	)

	value1, err = Ensure[string]("TEST_STRING")
	if err != nil {
		t.Errorf("failed to obtain value for key `TEST_STRING`")
	}

	if value1 != "test value" {
		t.Errorf("got %v, want `test string`", value1)
	}

	if reflect.TypeOf(value1).Kind() != reflect.String {
		t.Errorf("got %s, want string", reflect.TypeOf(value1).Kind())
	}

	value2, err = Obtain[string]("ANOTHER_STRING", "default value")
	if err != nil {
		t.Errorf("failed to obtain value for key `ANOTHER_STRING`")
	}

	if value2 != "default value" {
		t.Errorf("got %v, want `defail value`", value2)
	}

	if reflect.TypeOf(value2).Kind() != reflect.String {
		t.Errorf("got %s, want string", reflect.TypeOf(value2).Kind())
	}

	value3, err = Obtain[int]("TEST_NUMBER", 512)
	if err != nil {
		t.Errorf("failed to obtain value for key `TEST_NUMBER`")
	}

	if value3 != 1024 {
		t.Errorf("got %v, want 1024", value3)
	}

	if reflect.TypeOf(value3).Kind() != reflect.Int {
		t.Errorf("got %s, want int", reflect.TypeOf(value3).Kind())
	}

	value4, err = Obtain[bool]("TEST_BOOL", false)
	if err != nil {
		t.Errorf("failed to obtain value for key `TEST_BOOL`")
	}

	if value4 != true {
		t.Errorf("got %v, want true", value4)
	}

	if reflect.TypeOf(value4).Kind() != reflect.Bool {
		t.Errorf("got %s, want int", reflect.TypeOf(value4).Name())
	}

	values, err = Collect[string]("TEST_STRING", "TEST_NUMBER", "TEST_BOOL")
	if err != nil {
		t.Errorf("failed to obtain values for keys `TEST_STRING`. `TEST_NUMBER`, 'TEST_BOOL'")
	}

	if len(values) != 3 {
		t.Errorf("got %v elements, want 3 elements", value4)
	}

	if reflect.TypeOf(values).Kind() != reflect.Slice {
		t.Errorf("got %s, want slice", reflect.TypeOf(values).Name())
	}

	if reflect.TypeOf(values).Elem().Kind() != reflect.String {
		t.Errorf("got the slice of %s, want a slice of strings", reflect.TypeOf(values).Elem().Kind())
	}

	if values[0] != "test value" {
		t.Errorf("got %v, want `test value`", values[0])
	}

	if values[1] != "1024" {
		t.Errorf("got %v, want `1024`", values[1])
	}

	if values[2] != "TRUE" {
		t.Errorf("got %v, want `TRUE`", values[2])
	}
}
