package cmd

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkAndGenerateMessage(t *testing.T) {
	// Setup
	type args struct {
		target uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "target is a prime number(target=2)",
			args: args{target: 2},
			want: "Congrats!! The number 2 is a prime number.",
		},
		{
			name: "target is a prime number(target=3571)",
			args: args{target: 3571},
			want: "Congrats!! The number 3571 is a prime number.",
		},
		{
			name: "target is not a prime number(target=0)",
			args: args{target: 0},
			want: "The number 0 is not a prime number.\nNot even a natural number.",
		},
		{
			name: "target is not a prime number(target=1)",
			args: args{target: 1},
			want: "The number 1 is not a prime number.",
		},
		{
			name: "target is not a prime number(other)",
			args: args{target: 10},
			want: "The number 10 is not a prime number.",
		},
		{
			name: "Semi-normal:target is not a prime number(MaxUint64)",
			args: args{target: math.MaxUint64},
			want: "The number 18446744073709551615 is not a prime number.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute & Verify
			assert.Equal(t, checkAndGenerateMessage(tt.args.target), tt.want)
		})
	}
}
