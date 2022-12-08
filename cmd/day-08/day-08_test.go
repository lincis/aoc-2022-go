package main

import (
	"aoc-2022-go/internal/utils"
	"reflect"
	"testing"
)

func Test_parseForest(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test parsing",
			args: args{
				in: utils.ReadInputFile("test-08.txt"),
			},
			want: [][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseForest(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseForest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transpose(t *testing.T) {
	type args struct {
		x [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test transpose",
			args: args{
				x: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
			},
			want: [][]int{
				{3, 2, 6, 3, 3},
				{0, 5, 5, 3, 5},
				{3, 5, 3, 5, 3},
				{7, 1, 3, 4, 9},
				{3, 2, 2, 9, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_visible(t *testing.T) {
	type args struct {
		x    int
		line []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is visible 1",
			args: args{
				x:    5,
				line: []int{1, 3, 0},
			},
			want: true,
		},
		{
			name: "Is visible 2",
			args: args{
				x:    3,
				line: []int{1},
			},
			want: true,
		},
		{
			name: "Is visible 3",
			args: args{
				x:    3,
				line: []int{},
			},
			want: true,
		},
		{
			name: "Is not visible 1",
			args: args{
				x:    5,
				line: []int{1, 3, 5},
			},
			want: false,
		},
		{
			name: "Is not visible 2",
			args: args{
				x:    3,
				line: []int{1, 2, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visible(tt.args.x, tt.args.line); got != tt.want {
				t.Errorf("visible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeVisible(t *testing.T) {
	type args struct {
		forest  [][]int
		forestT [][]int
		i       int
		j       int
	}
	forest := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	forestT := transpose(forest)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tree visible 1",
			args: args{
				forest:  forest,
				forestT: forestT,
				i:       0,
				j:       3,
			},
			want: true,
		},
		{
			name: "tree visible 2",
			args: args{
				forest:  forest,
				forestT: forestT,
				i:       1,
				j:       1,
			},
			want: true,
		},
		{
			name: "tree visible 3",
			args: args{
				forest:  forest,
				forestT: forestT,
				i:       1,
				j:       2,
			},
			want: true,
		},
		{
			name: "tree visible 4",
			args: args{
				forest:  forest,
				forestT: forestT,
				i:       3,
				j:       1,
			},
			want: false,
		},
		{
			name: "tree visible 5",
			args: args{
				forest:  forest,
				forestT: forestT,
				i:       2,
				j:       3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeVisible(tt.args.forest, tt.args.forestT, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("treeVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOne(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test One",
			args: args{
				in: utils.ReadInputFile("test-08.txt"),
			},
			want: 21,
		},
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
		in []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Two",
			args: args{
				in: utils.ReadInputFile("test-08.txt"),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Two(tt.args.in); got != tt.want {
				t.Errorf("Two() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestScenic(t *testing.T) {
	type args struct {
		forest  [][]int
		forestT [][]int
		i       int
		j       int
	}
	forest := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	forestT := transpose(forest)
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tree scenic 2 1",
			args: args{
				forest:  forest,
				forestT: forestT,
				i:       2,
				j:       1,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scenicScore(tt.args.forest, tt.args.forestT, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("scenicScore() = %v, want %v", got, tt.want)
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
