package str2number

import (
	"strconv"
	"strings"
)

// ToNumber32 string convert to int32 or float32
func ToNumber32(val string) (interface{}, error) {
	var i interface{}
	n, err := toNumber(val, 32)
	if err != nil {
		return i, err
	}

	switch v := n.(type) {
	case int64:
		i = int32(v)
	case float64:
		i = float32(v)
	}

	return i, err
}

// ToNumber64 string convert to int64 or float64
func ToNumber64(val string) (interface{}, error) {
	return toNumber(val, 64)
}

func isFloat(val string) bool {
	if strings.Index(val, ".") != -1 {
		return true
	}

	return false
}

func toNumber(val string, bitSize int) (interface{}, error) {
	var i interface{}
	var err error
	if isFloat(val) {
		i, err = strconv.ParseFloat(val, bitSize)
	} else {
		i, err = strconv.ParseInt(val, 10, bitSize)
	}

	return i, err
}
