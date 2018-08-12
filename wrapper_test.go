package errors

import (
	. "errors"
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		err  error
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
				err:  Error("error message 0"),
				text: "error message 1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = '%v', wantErr '%v'", err, tt.wantErr)
			} else {
				if res := New("error message 1-> error message 0"); !reflect.DeepEqual(err, res) {
					t.Errorf("Wrap() error = '%v', wantErr '%v'", err, res)
				}
			}
		})
	}
}

func TestWrapf(t *testing.T) {
	type args struct {
		err    error
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
				err:    Error("error message 0"),
				format: "error message 1 (%s)",
				a:      []interface{}{"reason"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrapf(tt.args.err, tt.args.format, tt.args.a...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wrapf() error = '%v', wantErr '%v'", err, tt.wantErr)
			} else {
				if res := New("error message 1 (reason)-> error message 0"); !reflect.DeepEqual(err, res) {
					t.Errorf("Wrapf() error = '%v', wantErr '%v'", err, res)
				}
			}
		})
	}
}
