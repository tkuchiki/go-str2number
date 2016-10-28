package str2number

import (
	"reflect"
	"testing"
)

func TestToNumber32(t *testing.T) {
	var i interface{}
	var err error

	stri := "1"

	i, err = ToNumber32(stri)
	if err != nil {
		t.Error(err)
	}

	v := reflect.ValueOf(i)

	if v.Type().Name() != "int32" {
		t.Error("failed to convert int32")
	}

	strf := "1.2345"
	i, err = ToNumber32(strf)
	if err != nil {
		t.Error(err)
	}

	v = reflect.ValueOf(i)

	if v.Type().Name() != "float32" {
		t.Error("failed to convert float32")
	}

	invalid := "foobar"
	i, err = ToNumber32(invalid)

	if err == nil {
		t.Error("invalid")
	}
}

func TestToNumber64(t *testing.T) {
	var i interface{}
	var err error

	stri := "1"

	i, err = ToNumber64(stri)
	if err != nil {
		t.Error(err)
	}

	v := reflect.ValueOf(i)

	if v.Type().Name() != "int64" {
		t.Error("failed to convert int64")
	}

	strf := "1.2345"
	i, err = ToNumber64(strf)
	if err != nil {
		t.Error(err)
	}

	v = reflect.ValueOf(i)

	if v.Type().Name() != "float64" {
		t.Error("failed to convert float64")
	}

	invalid := "foobar"
	i, err = ToNumber64(invalid)

	if err == nil {
		t.Error("invalid")
	}
}
