package component

import "fmt"

type Folder struct {
	Components []IComponent
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c IComponent) {
	f.Components = append(f.Components, c)
}
