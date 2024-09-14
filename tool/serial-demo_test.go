package tool

import (
	"testing"
)

func TestSer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"the one"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ser()
		})
	}
}
