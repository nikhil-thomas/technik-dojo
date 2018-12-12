package log

import "testing"

func TestLog(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Log()
		})
	}
}
