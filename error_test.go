package errors

import (
	"testing"
)

func TestError(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				text: "error string",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = '%v', wantErr '%v'", err, tt.wantErr)
			}
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				format: "error string: %s",
				a:      []interface{}{"reason"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Errorf(tt.args.format, tt.args.a...); (err != nil) != tt.wantErr {
				t.Errorf("Errorf() error = '%v', wantErr '%v'", err, tt.wantErr)
			}
		})
	}
}
