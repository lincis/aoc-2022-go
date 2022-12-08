package main

import (
	"aoc-2022-go/internal/utils"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	type args struct {
		raw_input []string
	}
	tests := []struct {
		name string
		args args
		want input
	}{
		{
			name: "Test input parser",
			args: args{
				raw_input: utils.ReadInputFile("test-data.txt"),
			},
			want: input{
				stacks: [][]rune{
					{'N', 'Z'},
					{'D', 'C', 'M'},
					{'P'},
				},
				instructions: [][3]int{
					{1, 2, 1},
					{3, 1, 3},
					{2, 2, 1},
					{1, 1, 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInput(tt.args.raw_input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestOperate(t *testing.T) {
	type args struct {
		in input
	}
	tests := []struct {
		name string
		args args
		want input
	}{
		{
			name: "Test input Op",
			args: args{
				in: input{
					stacks: [][]rune{
						{'N', 'Z'},
						{'D', 'C', 'M'},
						{'P'},
					},
					instructions: [][3]int{
						{1, 2, 1},
						{3, 1, 3},
						{2, 2, 1},
						{1, 1, 2},
					},
				},
			},
			want: input{
				stacks: [][]rune{
					{'D', 'N', 'Z'},
					{'C', 'M'},
					{'P'},
				},
				instructions: [][3]int{
					{3, 1, 3},
					{2, 2, 1},
					{1, 1, 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Operate(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestOperateMultiple(t *testing.T) {
	type args struct {
		in input
	}
	tests := []struct {
		name string
		args args
		want input
	}{
		{
			name: "Test input Op one",
			args: args{
				in: input{
					stacks: [][]rune{
						{'N', 'Z'},
						{'D', 'C', 'M'},
						{'P'},
					},
					instructions: [][3]int{
						{1, 2, 1},
						{3, 1, 3},
						{2, 2, 1},
						{1, 1, 2},
					},
				},
			},
			want: input{
				stacks: [][]rune{
					{'D', 'N', 'Z'},
					{'C', 'M'},
					{'P'},
				},
				instructions: [][3]int{
					{3, 1, 3},
					{2, 2, 1},
					{1, 1, 2},
				},
			},
		},
		{
			name: "Test input multipleop",
			args: args{
				in: input{
					stacks: [][]rune{
						{'D', 'N', 'Z'},
						{'C', 'M'},
						{'P'},
					},
					instructions: [][3]int{
						{3, 1, 3},
						{2, 2, 1},
						{1, 1, 2},
					},
				},
			},
			want: input{
				stacks: [][]rune{
					{},
					{'C', 'M'},
					{'D', 'N', 'Z', 'P'},
				},
				instructions: [][3]int{
					{2, 2, 1},
					{1, 1, 2},
				},
			},
		},
		{
			name: "Test input multipleop",
			args: args{
				in: input{
					stacks: [][]rune{
						{},
						{'C', 'M'},
						{'D', 'N', 'Z', 'P'},
					},
					instructions: [][3]int{
						{2, 2, 1},
						{1, 1, 2},
					},
				},
			},
			want: input{
				stacks: [][]rune{
					{'C', 'M'},
					{},
					{'D', 'N', 'Z', 'P'},
				},
				instructions: [][3]int{
					{1, 1, 2},
				},
			},
		},
		{
			name: "Test input multipleop",
			args: args{
				in: input{
					stacks: [][]rune{
						{'C', 'M'},
						{},
						{'D', 'N', 'Z', 'P'},
					},
					instructions: [][3]int{
						{1, 1, 2},
					},
				},
			},
			want: input{
				stacks: [][]rune{
					{'M'},
					{'C'},
					{'D', 'N', 'Z', 'P'},
				},
				instructions: [][3]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OperateMultiple(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OperateMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOne(t *testing.T) {
	type args struct {
		input input
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test One",
			args: args{
				input: ParseInput(utils.ReadInputFile("test-data.txt")),
			},
			want: "CMZ",
		},
	}
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
		input input
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Two",
			args: args{
				input: ParseInput(utils.ReadInputFile("test-data.txt")),
			},
			want: "MCD",
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
