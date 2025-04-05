package main

import (
	"fmt"
	"strconv"
)

var FuncMapMath = map[string]any{
	"add": add,
	"sub": sub,
	"mul": mul,
	"div": div,
}

func convertToInt64(a any) (int64, error) {
	switch v := a.(type) {
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case float32:
		// floor? or round?
		return int64(v), nil
	case float64:
		// floor? or round?
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	}
	return 0, fmt.Errorf("Unable to parse %q", a)
}

func intType(a any) (int64, bool) {
	switch v := a.(type) {
	case int:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return int64(v), true
	case uint:
		return int64(v), true
	case uint32:
		return int64(v), true
	case uint64:
		return int64(v), true
	case string:
		val, err := strconv.ParseInt(v, 10, 64)
		return val, err == nil
	}
	return 0, false
}
func intPair(a, b any) (int64, int64, bool) {
	if ai, aIsInt := intType(a); aIsInt {
		if bi, bIsInt := intType(b); bIsInt {
			return ai, bi, true
		}
	}
	return 0, 0, false
}

func floatType(a any) (float64, bool) {
	switch v := a.(type) {
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case string:
		val, err := strconv.ParseFloat(v, 64)
		return val, err == nil
	}
	return 0, false
}

// convert any into a float64 or error
func num(a any) (float64, error) {
	switch v := a.(type) {
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	}
	return 0, fmt.Errorf("Unable to parse %q", a)
}
func numPair(a any, b any) (float64, float64, error) {
	af, err := num(a)
	if err != nil {
		return 0, 0, err
	}
	bf, err := num(b)
	if err != nil {
		return 0, 0, err
	}
	return af, bf, nil
}
func add(a, b any) (any, error) {
	if aInt, bInt, isInt := intPair(a, b); isInt {
		return aInt + bInt, nil
	}
	af, bf, err := numPair(a, b)
	if err != nil {
		return 0, err
	}
	return af + bf, nil
}
func sub(a, b any) (any, error) {
	if aInt, bInt, isInt := intPair(a, b); isInt {
		return aInt - bInt, nil
	}
	af, bf, err := numPair(a, b)
	if err != nil {
		return 0, err
	}
	return af - bf, nil
}
func mul(a, b any) (any, error) {
	if aInt, bInt, isInt := intPair(a, b); isInt {
		return aInt * bInt, nil
	}
	af, bf, err := numPair(a, b)
	if err != nil {
		return 0, err
	}
	return af * bf, nil
}
func div(a, b any) (any, error) {
	if aInt, bInt, isInt := intPair(a, b); isInt {
		if bInt == 0 {
			return 0, fmt.Errorf("Divide by zero...")
		}
		return aInt / bInt, nil
	}
	af, bf, err := numPair(a, b)
	if err != nil {
		return 0, err
	}
	if bf == 0 {
		return 0, fmt.Errorf("Divide by zero...")
	}
	return af / bf, nil
}
