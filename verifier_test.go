package errors

import (
	. "errors"
	"reflect"
	"testing"
)

func TestCheck(t *testing.T) {
	var err1, err2, err3 error
	err1 = New("error message 0")

	type args struct {
		err  *error
		cond bool
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				err:  &err1,
				cond: false,
				text: "error message 1",
			},
		},
		{
			name: "2",
			args: args{
				err:  &err2,
				cond: true,
				text: "error message 2",
			},
		},
		{
			name: "3",
			args: args{
				err:  &err3,
				cond: false,
				text: "error message 3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Check(tt.args.err, tt.args.cond, tt.args.text)
			switch tt.name {
			case "1":
				res := New("error message 0")
				if !reflect.DeepEqual(*tt.args.err, res) {
					t.Errorf("Check() error = '%v', wantErr '%v'", *tt.args.err, res)
				}
			case "2":
				if !reflect.DeepEqual(*tt.args.err, nil) {
					t.Errorf("Check() error = '%v', wantErr '%v'", *tt.args.err, nil)
				}
			case "3":
				res := New("error message 3")
				if !reflect.DeepEqual(*tt.args.err, res) {
					t.Errorf("Check() error = '%v', wantErr '%v'", *tt.args.err, res)
				}
			}
		})
	}
}

func TestCheckf(t *testing.T) {
	var err1, err2, err3 error
	err1 = New("error message 0")

	type args struct {
		err    *error
		cond   bool
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				err:    &err1,
				cond:   false,
				format: "error message 1: %s",
				a:      []interface{}{"reason 1"},
			},
		},
		{
			name: "2",
			args: args{
				err:    &err2,
				cond:   true,
				format: "error message 2: %s",
				a:      []interface{}{"reason 2"},
			},
		},
		{
			name: "3",
			args: args{
				err:    &err3,
				cond:   false,
				format: "error message 3: %s",
				a:      []interface{}{"reason 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Checkf(tt.args.err, tt.args.cond, tt.args.format, tt.args.a...)
			switch tt.name {
			case "1":
				res := New("error message 0")
				if !reflect.DeepEqual(*tt.args.err, res) {
					t.Errorf("Check() error = '%v', wantErr '%v'", *tt.args.err, res)
				}
			case "2":
				if !reflect.DeepEqual(*tt.args.err, nil) {
					t.Errorf("Check() error = '%v', wantErr '%v'", *tt.args.err, nil)
				}
			case "3":
				res := New("error message 3: reason 3")
				if !reflect.DeepEqual(*tt.args.err, res) {
					t.Errorf("Check() error = '%v', wantErr '%v'", *tt.args.err, res)
				}
			}
		})
	}
}
