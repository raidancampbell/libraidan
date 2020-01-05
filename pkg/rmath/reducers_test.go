package rmath

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"happy", args{[]int{1, 2, 3}}, 3},
		{"big", args{[]int{1, 2, 30000}}, 30000},
		{"negative", args{[]int{-1, -2, -3}}, -1},
		{"abuse", args{[]int{}}, math.MinInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := Max(tt.args.n...); gotMax != tt.wantMax {
				t.Errorf("Max() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
	}{
		{"happy", args{[]int{1, 2, 3}}, 1},
		{"big", args{[]int{1, 2, 30000}}, 1},
		{"negative", args{[]int{-1, -2, -3}}, -3},
		{"abuse", args{[]int{}}, math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := Min(tt.args.n...); gotMin != tt.wantMin {
				t.Errorf("Min() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"happy", args{[]int{1, 2, 3}}, 6},
		{"big", args{[]int{1, 2, 30000}}, 30003},
		{"negative", args{[]int{-1, -2, -3}}, -6},
		{"abuse", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.n...); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
