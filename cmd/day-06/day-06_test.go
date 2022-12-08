package main

import (
	"aoc-2022-go/internal/utils"
	"testing"
)

func Test_allDifferent(t *testing.T) {
	type args struct {
		in []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Different 4",
			args: args{in: []rune{'A', 'B', 'V', 'C'}},
			want: true,
		},
		{
			name: "Different 6",
			args: args{in: []rune{'A', 'B', 'V', 'C', '6', ' '}},
			want: true,
		},
		{
			name: "Equal 4",
			args: args{in: []rune{'A', 'B', 'V', 'A'}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allDifferent(tt.args.in); got != tt.want {
				t.Errorf("allDifferent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartMarker(t *testing.T) {
	type args struct {
		in     string
		length int
	}
	test_inputs := utils.ReadInputFile("test-06.txt")
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test line 1",
			args: args{in: test_inputs[0], length: 4},
			want: 7,
		},
		{
			name: "Test line 2",
			args: args{in: test_inputs[1], length: 4},
			want: 5,
		},
		{
			name: "Test line 3",
			args: args{in: test_inputs[2], length: 4},
			want: 6,
		},
		{
			name: "Test line 4",
			args: args{in: test_inputs[3], length: 4},
			want: 10,
		},
		{
			name: "Test line 5",
			args: args{in: test_inputs[4], length: 4},
			want: 11,
		},
		{
			name: "Test line 1 14",
			args: args{in: test_inputs[0], length: 14},
			want: 19,
		},
		{
			name: "Test line 2 14",
			args: args{in: test_inputs[1], length: 14},
			want: 23,
		},
		{
			name: "Test line 3 14",
			args: args{in: test_inputs[2], length: 14},
			want: 23,
		},
		{
			name: "Test line 4 14",
			args: args{in: test_inputs[3], length: 14},
			want: 29,
		},
		{
			name: "Test line 5 14",
			args: args{in: test_inputs[4], length: 14},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartMarker(tt.args.in, tt.args.length); got != tt.want {
				t.Errorf("StartMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOne(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := One(tt.args.in); got != tt.want {
				t.Errorf("One() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two(tt.args.in); got != tt.want {
				t.Errorf("Two() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
