# VitaoDemolay collections
This is a simple pack of data structs for collections to help implements lists and other structs for generics objects.

### For using
```shell
    go get github.com/vitaodemolay/collections
``` 

## ReadOnlyList Interface
It`s an interface that provides methods to access a list without modifying it. This contains this methods:
-  Get
-  GetAll
-  Size
-  Find

### Get
Get returns the element at the specified index and a boolean indicating if it exists.

Ex: 
``` GO
    Get(index int) (T, bool)
```

### GetAll
GetAll returns all elements in the list.

Ex:
``` GO
    GetAll() []T
```

### Size
Size returns the number of elements in the list.

Ex:
``` GO
    Size() int
```

### Find
Find returns the first element that matches the given query function and its index. This method uses a parameter as predicate taht determines if an element matches the search criteria. 

Ex:
``` GO
    Find(query func(T) bool) (T, int)
```

## List Interface
It`s an interface that extends ReadOnlyList and adds methods to modify the list. This add this methods:
- Add
- Remove
- Clear

### Add
Add appends a new element to the list.

Ex:
``` GO
    Add(value T)
```

### Remove
Remove deletes an element from the list by index and returns true if it was removed, and false if not found.

Ex:
``` GO
    Remove(index int) bool
```

### Clear
Clear removes all elements from the list.

Ex:
``` GO
    Clear()
```

## Sample of using pack
``` GO
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
```