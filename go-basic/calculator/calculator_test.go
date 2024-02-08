package calculator

import (
	"testing"
)

func TestMyCalculator_IsOdd(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		m    CalculatorInterface
		args args
		want bool
	}{
		{
			name: "Test 90 IsOdd expect false",
			m:    NewTheBestCalculator("My Calculator"),
			args: args{90},
			want: false,
		},
		{
			name: "Test 0 IsOdd expect false",
			m:    NewTheBestCalculator("My Calculator"),
			args: args{0},
			want: false,
		},
		{
			name: "Test 7 IsOdd expect true",
			m:    NewTheBestCalculator("My Calculator"),
			args: args{7},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsOdd(tt.args.x); got != tt.want {
				t.Errorf("MyCalculator.IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCalculator_Power(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		m    CalculatorInterface
		args args
		want float64
	}{
		{
			name: "Test Power(3,2) expect 9",
			m:    NewTheBestCalculator("My calculator"),
			args: args{3, 2},
			want: 9,
		},
		{
			name: "Test Power(2,4) expect 16",
			m:    NewTheBestCalculator("My calculator"),
			args: args{2, 4},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Power(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("MyCalculator.Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCalculator_Square(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		m    CalculatorInterface
		args args
		want float64
	}{
		{
			name: "Test Sqrt(16) expect 4",
			m:    NewTheBestCalculator("My Cal"),
			args: args{16},
			want: 4,
		},
		{
			name: "Test Sqrt(100) expect 10",
			m:    NewTheBestCalculator("My Cal"),
			args: args{100},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Square(tt.args.input); got != tt.want {
				t.Errorf("MyCalculator.Square() = %v, want %v", got, tt.want)
			}
		})
	}
}
