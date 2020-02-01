package roper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilOr(t *testing.T) {
	var ni interface{}
	nn := "chosen"
	actual := NilOr(ni, nn)
	assert.Equal(t, "chosen", actual)
	actual = NilOr(nn, ni)
	assert.Equal(t, "chosen", actual)
}

func TestEmptyOr(t *testing.T) {
	ni := ""
	nn := "chosen"
	actual := EmptyOr(ni, nn)
	assert.Equal(t, "chosen", actual)
	actual = EmptyOr(nn, ni)
	assert.Equal(t, "chosen", actual)
}

func TestZeroOr(t *testing.T) {
	ni := 0
	nn := -1
	actual := ZeroOr(ni, nn)
	assert.Equal(t, -1, actual)
	actual = ZeroOr(nn, ni)
	assert.Equal(t, -1, actual)
}

func TestFirstStr(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"both non empty", args{[]string{"foo", "bar"}}, "foo"},
		{"first empty", args{[]string{"", "bar"}}, "bar"},
		{"second empty", args{[]string{"foo", ""}}, "foo"},
		{"middle empty", args{[]string{"foo", "", "baz"}}, "foo"},
		{"middle only", args{[]string{"", "foo", ""}}, "foo"},
		{"all empty", args{[]string{"", "", ""}}, ""},
		{"no args", args{[]string{}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstStr(tt.args.s...); got != tt.want {
				t.Errorf("FirstStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"both non empty", args{[]int{-1, 1}}, -1},
		{"first empty", args{[]int{0, 1}}, 1},
		{"second empty", args{[]int{-1, 0}}, -1},
		{"middle empty", args{[]int{-1, 0, 1}}, -1},
		{"middle only", args{[]int{0, 5, 0}}, 5},
		{"all empty", args{[]int{0, 0, 0}}, 0},
		{"no args", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstInt(tt.args.s...); got != tt.want {
				t.Errorf("FirstInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
