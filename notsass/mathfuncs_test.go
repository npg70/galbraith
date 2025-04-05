package main

import (
	"reflect"
	"testing"
)

var testNumData = []struct {
	val  any
	want float64
	err  bool
}{
	{0, 0, false},
	{1, 1.0, false},
	{int(2), 2.0, false},
	{int32(3), 3.0, false},
	{int64(4), 4.0, false},
	{uint(5), 5.0, false},
	{uint32(6), 6.0, false},
	{uint64(7), 7.0, false},
	{float32(8.0), 8.0, false},
	{float64(9.0), 9.0, false},
	{"1.125", 1.125, false},
	{"1234", 1234, false},
	{nil, 0, true},
	{"junk", 0, true},
}

func TestNum(t *testing.T) {
	for i, tc := range testNumData {
		val, err := num(tc.val)
		if !tc.err && err != nil {
			t.Errorf("Got error in test case %d", i)
		}
		if val != tc.want {
			t.Errorf("Expected %f, got %f", tc.want, val)
		}
	}
}

var testIntData = []struct {
	val  any
	want int64
	err  bool
}{
	{0, 0, true},
	{int(1), 1, true},
	{int32(2), 2, true},
	{int64(3), 3, true},
	{uint(4), 4, true},
	{uint32(5), 5, true},
	{uint64(6), 6, true},
	{float32(7), 0, false},
	{float64(8), 0, false},
	{"9", 9, true},
	{"123.456", 0, false},
}

func TestIntData(t *testing.T) {
	for _, tt := range testIntData {
		val, ok := intType(tt.val)
		if tt.err != ok {
			t.Errorf("Expected conversion of %v to be %v, got %v", tt.val, tt.err, ok)
			continue
		}
		if val != tt.want {
			t.Errorf("Expected conversion of %v to %d, got %d", tt.val, tt.want, val)
		}
	}
}

var testAddData = []struct {
	a    any
	op   string
	b    any
	want any
	ok   bool
}{
	{0, "+", 0, int64(0), true},
	{int(1), "+", int(2), int64(3), true},
	{"123", "+", "junk", 0, false},
	{"10", "+", "20", int64(30), true},
	{"10", "+", "20.0", float64(30), true},
	{"10.0", "+", "20.0", float64(30), true},
	{10.0, "+", 20.0, float64(30), true},
	{10, "+", "20", int64(30), true},
	{10, "+", "20.0", float64(30), true},

	{12, "/", 3, int64(4), true},
	{12, "/", 3.0, float64(4), true},
	{12, "/", "0", 0, false},
	{12, "/", "0.0", 0, false},
	{12, "/", int(0), 0, false},
}

func TestAddData(t *testing.T) {
	for _, tt := range testAddData {
		var got any
		var err error
		switch tt.op {
		case "+":
			got, err = add(tt.a, tt.b)
		case "-":
			got, err = sub(tt.a, tt.b)
		case "*":
			got, err = mul(tt.a, tt.b)
		case "/":
			got, err = div(tt.a, tt.b)
		default:
			panic("Got unknown op " + tt.op)
		}
		if err == nil && !tt.ok {
			t.Errorf("For %v %s %v, expected no error but got %v", tt.a, tt.op, tt.b, err)
			continue
		}
		if err != nil && tt.ok {
			t.Errorf("For %v %s %v, expected error and didn't get one", tt.a, tt.op, tt.b)
			continue
		}
		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf("For %v %s %v, got type %v, want %v", tt.a, tt.op, tt.b, reflect.TypeOf(got), reflect.TypeOf(tt.want))
			continue
		}
	}
}
