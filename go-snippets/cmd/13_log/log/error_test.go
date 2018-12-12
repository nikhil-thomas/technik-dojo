package log

import "testing"

func TestOriginalError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OriginalError(); (err != nil) != test.wantErr {
				t.Errorf("OriginalError(0 erro = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestPassThroughError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := PassThroughError(); (err != nil) != test.wantErr {
				t.Errorf("OriginalError(0 erro = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Handler()
		})
	}
}
