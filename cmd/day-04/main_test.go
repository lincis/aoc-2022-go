package main

import (
	"aoc-2022-go/internal/utils"
	"fmt"
	"os"
	"testing"
)

func TestOne(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test part One",
			args: args{
				input: utils.ReadInputFile("test-data.txt"),
			},
			want: 2,
		},
	}
	fmt.Println(os.Getwd())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := One(tt.args.input); got != tt.want {
				t.Errorf("One() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test part Two",
			args: args{
				input: utils.ReadInputFile("test-data.txt"),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two(tt.args.input); got != tt.want {
				t.Errorf("Two() = %v, want %v", got, tt.want)
			}
		})
	}
}
