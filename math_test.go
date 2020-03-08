package math

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		operation string
		first     float64
		second    float64
		want      float64
	}{
		{
			operation: "+",
			first:     12,
			second:    10,
			want:      22,
		},
		{
			operation: "-",
			first:     10,
			second:    4,
			want:      6,
		},
		{
			operation: "*",
			first:     6,
			second:    6,
			want:      36,
		},
		{
			operation: "/",
			first:     10,
			second:    2,
			want:      5,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("Checking calc for %s, %f, %f", tt.operation, tt.first, tt.second)
		t.Run(name, func(t *testing.T) {
			operation := Calc(tt.operation)
			got := operation(tt.first, tt.second)

			if got != tt.want {
				t.Errorf("Expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
