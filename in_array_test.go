package in_array

import "testing"

func TestInArray(t *testing.T) {
	tcs := []struct {
		summary    string
		inputArray interface{}
		input      interface{}
		output     bool
	}{
		{"int comparison true", []int{40, 50, 35, 42}, 42, true},
		{"int comparison false", []int{40, 50, 35, 42}, 24, false},
		{"int negative comparison true", []int{-42, 40, 50, 35, 42}, -42, true},
		{"int comparison against empty array false", []int{}, 42, false},
		{"int8 comparison false", []int8{40, 50, 35}, int8(42), false},
		{"int16 comparison true", []int16{41, 42, 43}, int16(42), true},
		{"int32 comparison true", []int32{41, 42, 43}, int32(42), true},
		{"int64 comparison false", []int64{14, 41, 43, 44, 45, 46}, int32(42), false},
		{"float32 comparison true", []float32{40.5, 41.9, 42.1, 42.0, 42.9, 50.60, 35.98}, float32(42), true},
		{"float64 comparison false", []float64{40.5, 41.9, 42.1, 42.0000001, 42.9, 50.60, 35.98}, float64(42), false},
		{"uint comparison true", []uint{40, 50, 35, 42}, uint(42), true},
		{"uint comparison against empty array false", []uint{}, uint(42), false},
		{"uint8 comparison false", []uint8{40, 50, 35}, uint8(42), false},
		{"uint16 comparison true", []uint16{41, 42, 43}, uint16(42), true},
		{"uint32 comparison true", []uint32{41, 42, 43}, uint32(42), true},
		{"uint64 comparison false", []uint64{14, 41, 43, 44, 45, 46}, uint32(42), false},
		{"string comparison true", []string{"ha", "bla", "cle"}, "bla", true},
		{"string comparison false", []string{"abc", "def", "ghi"}, "jkl", false},
		{"string case-difference false", []string{"", "abc", "def", "ghi"}, "DEF", false},
		{"string empty comparison true", []string{"", "abc", "def", "ghi"}, "", true},
		{"boolean comparison false", []bool{true, true, true}, false, false},
		{"boolean comparison true", []bool{true, true, true}, true, true},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := InArray(tc.inputArray, tc.input)

			if tr != tc.output {
				t.Errorf("Test has failed!\n\tExpected: %t, \n\tGot: %t, \n\tarray: %#v, \n\tItem: %v", tc.output, tr, tc.inputArray, tc.input)
			}
		})
	}
}