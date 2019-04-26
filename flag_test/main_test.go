package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		options []string
		result  string
	}{
		{options: []string{"-s", "sss"}, result: "sss,,[]"},
		{options: []string{"-t", "ttt"}, result: ",ttt,[]"},
		{options: []string{}, result: ",,[]"},
		{options: []string{"-s", "sss", "-t", "ttt"}, result: "sss,ttt,[]"},
		{options: []string{"xxx"}, result: ",,[xxx]"},
	}

	for case_number, test := range tests {
		os.Args = []string{os.Args[0]}
		os.Args = append(os.Args, test.options...)
		result := _main()

		if result != test.result {
			t.Errorf("Case:[%v] received: %v - expected: %v - os.Args: %v", case_number, result, test.result, os.Args)
		} else {
			t.Logf("OK Case:[%v] received: %v - os.Args: %v", case_number, result, os.Args)
		}
	}
}
