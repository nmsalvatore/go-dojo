package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var runTests = []struct {
		name    string
		in      []string
		out     string
		wantErr bool
	}{
		{
			name:    "one valid integer",
			in:      []string{"3"},
			out:     "Min: 3\nMax: 3\nMean: 3.00\nMedian: 3.00\n",
			wantErr: false,
		},
		{
			name:    "multiple valid integers",
			in:      []string{"1", "2", "3"},
			out:     "Min: 1\nMax: 3\nMean: 2.00\nMedian: 2.00\n",
			wantErr: false,
		},
		{
			name:    "no args",
			in:      []string{},
			out:     "",
			wantErr: true,
		},
		{
			name:    "includes non-integer",
			in:      []string{"3", "abc", "5"},
			out:     "",
			wantErr: true,
		},
	}

	for _, tt := range runTests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			err := run(&buf, tt.in)

			if tt.wantErr && err == nil {
				t.Error("want error, didn't get one")
				return
			}

			if !tt.wantErr && err != nil {
				t.Error("got error, didn't want one")
				return
			}

			if buf.String() != tt.out {
				t.Errorf("got %s, want %s", buf.String(), tt.out)
			}
		})
	}
}
