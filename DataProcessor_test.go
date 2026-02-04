package main

import (
	"errors"
	"strconv"
	"testing"
)

func Test_parser(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want Input
		err  error
	}{
		{
			name: "valid, expected input",
			data: []byte("1\n*\n2\n4"),
			want: Input{
				Id:   "1",
				Op:   "*",
				Val1: 2,
				Val2: 4,
			},
		},
		{
			name: "can't parse num1",
			data: []byte("1\n*\na\n4"),
			want: Input{},
			err:  strconv.ErrSyntax,
		},
		{
			name: "can't parse num2",
			data: []byte("1\n*\n2\na"),
			want: Input{},
			err:  strconv.ErrSyntax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser(tt.data)
			if err != nil {
				if tt.err == nil {
					t.Fatal(err)
				}
				if !errors.Is(errors.Unwrap(err), tt.err) {
					t.Fatalf("parser() error = %v, wantErr %v", err, tt.err)
				}
			}
			if got != tt.want {
				t.Errorf("got: %v\nwant: %v", got, tt.want)
			}
		})
	}
}
