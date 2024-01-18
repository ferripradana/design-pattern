package node

import "fmt"

type Folder struct {
	Children []INode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, child := range f.Children {
		child.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{
		Name: f.Name + "_clone",
	}
	var tempChildren []INode
	for _, child := range f.Children {
		copy := child.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
