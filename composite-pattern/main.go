package main

import (
	"composite-pattern/component"
	"fmt"
)

func main() {
	fmt.Println("RUN...")
	file1 := &component.File{Name: "File1"}
	file2 := &component.File{Name: "File2"}
	file3 := &component.File{Name: "File3"}

	folder1 := &component.Folder{Name: "Folder1"}
	folder1.Add(file1)
	folder2 := &component.Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
