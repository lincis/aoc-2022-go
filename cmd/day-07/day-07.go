package main

import (
	"aoc-2022-go/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type directory struct {
	parent   string
	children []string
	files    map[string]int
}

func fullPath(name string, parent string, tree map[string]directory) string {
	path := name
	return path
}

func parseTree(in []string) map[string]directory {
	tree := make(map[string]directory)
	// readStdout := false
	cwd := ""
	cwdStruct := directory{parent: "", children: []string{}, files: make(map[string]int)}
	for _, line := range in {
		// fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == '$' { //command
			if cwd != "" {
				tree[cwd] = cwdStruct
			}
			if line[2:4] == "cd" {
				// readStdout = false
				if line[5:] == ".." {
					cwd = cwdStruct.parent
					cwdStruct = tree[cwd]
					continue
				}
				cwdStruct = directory{parent: cwd, children: []string{}, files: make(map[string]int)}
				newCwd := line[5:]
				if cwd != "" {
					cwd = cwd + newCwd + "/"
				} else {
					cwd = newCwd
				}
				// fmt.Println(cwd, ": ", line)
			} // else if line[2:3] == "ls" {
			// readStdout = true
			// }
		} else {
			parts := strings.Split(line, " ")
			// fmt.Println(parts)
			if parts[0] == "dir" {
				cwdStruct.children = append(cwdStruct.children, cwd+parts[1]+"/")
			} else {
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Println("Error parsing line to size / file: ", line, " (", err, ")")
					continue
				}
				// fmt.Println(parts[1], ": ", size)
				// fmt.Println(cwdStruct.files)
				cwdStruct.files[parts[1]] = size
			}
		}
	}
	return tree
}

func directorySize(tree map[string]directory) map[string]int {
	allDirSizes := make(map[string]int)
	for dir := range tree {
		dirSize := 0
		for dir2, dirStruct := range tree {
			if len(dir) > len(dir2) {
				continue
			}
			if dir == dir2[:len(dir)] {
				for _, fileSize := range dirStruct.files {
					dirSize += fileSize
				}
			}
		}
		allDirSizes[dir] = dirSize
	}
	return allDirSizes
}

func One(in []string) int {
	tree := parseTree(in)
	// fmt.Println(tree)
	// fmt.Println(len(tree))
	sizes := directorySize(tree)
	fmt.Println(len(sizes))
	sizeBelow10k := 0
	for _, size := range sizes {
		if size > 100000 {
			continue
		}
		sizeBelow10k += size
	}
	return sizeBelow10k
}

func Two(in []string) int {
	tree := parseTree(in)
	sizes := directorySize(tree)
	totalSize := 0
	for _, dir := range tree {
		for _, size := range dir.files {
			totalSize += size
		}
	}
	// fmt.Println("Total size = ", totalSize)
	requiredSize := 30000000 - (70000000 - sizes["/"])
	// fmt.Println("Used ", sizes["/"], " required ", requiredSize)
	// fmt.Println(requiredSize)
	// nameToDelete := "/"
	sizeToDelete := sizes["/"]
	// fmt.Println(sizeToDelete)
	for _, size := range sizes {
		if size > requiredSize && size < sizeToDelete {
			// nameToDelete = dir
			sizeToDelete = size
			// fmt.Println(sizeToDelete)
		}
	}
	return sizeToDelete
}

func main() {
	raw_data := utils.ReadInputFile("input/day-07.txt")
	// fmt.Println(raw_data)
	fmt.Println("Part One = ", One(raw_data))
	fmt.Println("Part Two = ", Two(raw_data))
}
