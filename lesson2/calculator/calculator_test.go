package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Add 0, 0 Expect 0",
			args: args{0, 0},
			want: 0,
		},
		{
			name: "Test Add -10, -10 Expect -20",
			args: args{-10, -10},
			want: -20,
		},
		{
			name: "Test Add 10, -10 Expect 0",
			args: args{10, -10},
			want: 0,
		},
		{
			name: "Test Add 1, 2 Expect 3",
			args: args{1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrct(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Subtract 0, 0 Expect 0",
			args: args{0, 0},
			want: 0,
		},
		{
			name: "Test Subtract -1, 0 Expect -1",
			args: args{-1, 0},
			want: -1,
		},
		{
			name: "Test Subtract -1, -3 Expect 2",
			args: args{-1, -3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substrct(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Substrct() = %v, want %v", got, tt.want)
			}
		})
	}
}
