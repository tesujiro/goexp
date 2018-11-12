package main

import (
	"testing"
)

func TestFlag(t *testing.T) {
	tests := []struct {
		prepare  func()
		cleanup  func()
		expected string
	}{
		{prepare: func() { *arg_bool = true }, cleanup: func() { *arg_bool = false }, expected: "bool:true"},
		{expected: "bool:false"},
	}
	for _, test := range tests {
		if test.prepare != nil {
			test.prepare()
		}
		actual := parse()
		if actual != test.expected {
			t.Errorf("got: %v\texpected: %v\n", actual, test.expected)
		}
		if test.cleanup != nil {
			test.cleanup()
		}
	}
}
