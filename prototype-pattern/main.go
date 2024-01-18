package main

import (
	"fmt"
	"prototype-pattern/node"
)

func main() {
	file1 := &node.File{Name: "File 1"}
	file2 := &node.File{Name: "File 2"}
	file3 := &node.File{Name: "File 3"}

	folder1 := &node.Folder{
		Children: []node.INode{file1},
		Name:     "Folder 1",
	}

	folder2 := &node.Folder{
		Children: []node.INode{folder1, file2, file3},
		Name:     "Folder 2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print(" ")

	var cloneFolder node.INode
	cloneFolder = folder2.Clone()
	fmt.Println("\nPrinting hierarchy for cloneFolder")
	cloneFolder.Print(" ")

}
