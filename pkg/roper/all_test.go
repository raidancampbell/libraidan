package roper

import (
	"testing"
)

func TestAllEmpty(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"happy_false", args{[]string{"foo", "bar"}}, false},
		{"happy_true", args{[]string{"", ""}}, true},
		{"void", args{[]string{}}, true},
		{"single", args{[]string{"", "foo"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllEmpty(tt.args.a...); got != tt.want {
				t.Errorf("AllEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllNil(t *testing.T) {
	type args struct {
		a []interface{}
	}

	strs := []string{"", "foo", "bar"}
	nonNils := make([]interface{}, len(strs))
	oneNil := make([]interface{}, len(strs))
	allNil := make([]interface{}, len(strs))
	for i, s := range strs {
		nonNils[i] = s
		oneNil[i] = s
		allNil[i] = s
	}

	oneNil[0] = nil

	allNil[0] = nil
	allNil[1] = nil
	allNil[2] = nil

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"happy_false", args{allNil}, true},
		{"happy_true", args{nonNils}, false},
		{"one_nil", args{oneNil}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllNil(tt.args.a...); got != tt.want {
				t.Errorf("AllNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllZero(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"happy_false", args{[]int{1, 2, 3}}, false},
		{"happy_true", args{[]int{0, 0, 0, 0}}, true},
		{"void", args{[]int{}}, true},
		{"single", args{[]int{0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllZero(tt.args.a...); got != tt.want {
				t.Errorf("AllZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
