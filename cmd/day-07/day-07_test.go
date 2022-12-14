package main

import (
	"aoc-2022-go/internal/utils"
	"reflect"
	"testing"
)

func Test_parseTree(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want map[string]directory
	}{
		{
			name: "Test parsing",
			args: args{
				in: utils.ReadInputFile("test-07.txt"),
			},
			want: map[string]directory{
				"/":     {parent: "", children: []string{"/a/", "/d/"}, files: map[string]int{"b.txt": 14848514, "c.dat": 8504156}},
				"/a/":   {parent: "/", children: []string{"/a/e/"}, files: map[string]int{"f": 29116, "g": 2557, "h.lst": 62596}},
				"/a/e/": {parent: "/a/", children: []string{}, files: map[string]int{"i": 584}},
				"/d/":   {parent: "/", children: []string{}, files: map[string]int{"j": 4060174, "d.log": 8033020, "d.ext": 5626152, "k": 7214296}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTree(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directorySize(t *testing.T) {
	type args struct {
		tree map[string]directory
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Test dir sizes",
			args: args{
				tree: map[string]directory{
					"/":     {parent: "", children: []string{"/a/", "/d/"}, files: map[string]int{"b.txt": 14848514, "c.dat": 8504156}},
					"/a/":   {parent: "/", children: []string{"/a/e/"}, files: map[string]int{"f": 29116, "g": 2557, "h.lst": 62596}},
					"/a/e/": {parent: "/a/", children: []string{}, files: map[string]int{"i": 584}},
					"/d/":   {parent: "/", children: []string{}, files: map[string]int{"j": 4060174, "d.log": 8033020, "d.ext": 5626152, "k": 7214296}},
				},
			},
			want: map[string]int{
				"/":     48381165,
				"/a/e/": 584,
				"/a/":   94853,
				"/d/":   24933642,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := directorySize(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("directorySize() = %v, want %v", got, tt.want)
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
				in: utils.ReadInputFile("test-07.txt"),
			},
			want: 95437,
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
			name: "Part Two",
			args: args{
				in: utils.ReadInputFile("test-07.txt"),
			},
			want: 24933642,
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
