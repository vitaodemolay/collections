package main

import (
	"fmt"
	"strings"

	"github.com/vitaodemolay/collections/list"
)

func main() {
	lst := list.CreateList[myStructDummy]()
	lst.Add(myStructDummy{ID: 1, Name: "John"})
	lst.Add(myStructDummy{ID: 2, Name: "Jane"})

	fmt.Println("List size:", lst.Size())
	fmt.Println("List contents:", lst.GetAll())
	lst.Remove(1)
	fmt.Println("List after removal:", lst.GetAll())
	vl, id := lst.Find(func(v myStructDummy) bool {
		return strings.HasPrefix(v.Name, "Jo")
	})

	fmt.Println("Found with prefix [Jo]:", vl, "at index", id)
	lst.Clear()
	fmt.Println("List after clear:", lst.GetAll())
	fmt.Println("List size after clear:", lst.Size())
}

type myStructDummy struct {
	ID   int
	Name string
}
