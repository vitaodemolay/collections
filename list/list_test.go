package list

import (
	"testing"
)

/*********************************************************
 *************Test Cases Using int as Type****************
 *********************************************************/

// TestListInt tests the list implementation with int type.
// It checks adding, retrieving, removing elements, and clearing the list.
// It also verifies the size of the list at different stages.
func TestListInt(t *testing.T) {
	l := newList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	if val, ok := l.Get(1); !ok || val != 2 {
		t.Errorf("Expected value 2, got %v", val)
	}

	if val := l.GetAll(); len(val) != 3 {
		t.Errorf("Expected 3 values, got %d", len(val))
	}

	if _, ok := l.Get(3); ok {
		t.Error("Expected no value at index 3")
	}

	if !l.Remove(1) {
		t.Error("Expected to remove value at index 1")
	}

	if val, ok := l.Get(1); !ok || val != 3 {
		t.Errorf("Expected value 3, got %v", val)
	}

	vl, id := l.Find(func(value int) bool {
		return value == 1
	})
	if id == -1 {
		t.Error("Expected to find value 1, got -1")
	} else if vl != 1 {
		t.Errorf("Expected value 1, got %v", vl)
	}

	l.Clear()
	if l.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", l.Size())
	}
}

// TestListIntWithValues tests the list implementation with int type and initial values.
// It checks adding, retrieving, removing elements, and clearing the list.
// It also verifies the size of the list at different stages.
func TestListIntWithValues(t *testing.T) {
	l := newListWithValues([]int{1, 2, 3})
	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	if val, ok := l.Get(1); !ok || val != 2 {
		t.Errorf("Expected value 2, got %v", val)
	}

	if val := l.GetAll(); len(val) != 3 {
		t.Errorf("Expected 3 values, got %d", len(val))
	}

	if _, ok := l.Get(3); ok {
		t.Error("Expected no value at index 3")
	}

	if !l.Remove(1) {
		t.Error("Expected to remove value at index 1")
	}

	if val, ok := l.Get(1); !ok || val != 3 {
		t.Errorf("Expected value 3, got %v", val)
	}

	vl, id := l.Find(func(value int) bool {
		return value == 1
	})
	if id == -1 {
		t.Error("Expected to find value 1, got -1")
	} else if vl != 1 {
		t.Errorf("Expected value 1, got %v", vl)
	}

	l.Clear()
	if l.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", l.Size())
	}
}

/*********************************************************
 ************Test Cases Using string as Type**************
 *********************************************************/

// TestListString tests the list implementation with string type.
// It checks adding, retrieving, removing elements, and clearing the list.
// It also verifies the size of the list at different stages.
func TestListString(t *testing.T) {
	l := newList[string]()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	if val, ok := l.Get(1); !ok || val != "b" {
		t.Errorf("Expected value b, got %v", val)
	}

	if val := l.GetAll(); len(val) != 3 {
		t.Errorf("Expected 3 values, got %d", len(val))
	}

	if _, ok := l.Get(3); ok {
		t.Error("Expected no value at index 3")
	}

	if !l.Remove(1) {
		t.Error("Expected to remove value at index 1")
	}

	if val, ok := l.Get(1); !ok || val != "c" {
		t.Errorf("Expected value c, got %v", val)
	}

	vl, id := l.Find(func(value string) bool {
		return value == "a"
	})
	if id == -1 {
		t.Error("Expected to find value a, got -1")
	} else if vl != "a" {
		t.Errorf("Expected value a, got %v", vl)
	}

	l.Clear()
	if l.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", l.Size())
	}
}

/*********************************************************
 ************Test Cases Using float as Type***************
 *********************************************************/

// TestListFloat tests the list implementation with float type.
// It checks adding, retrieving, removing elements, and clearing the list.
// It also verifies the size of the list at different stages.
func TestListFloat(t *testing.T) {
	l := newListWithValues([]float64{1.1, 2.2, 3.3})

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	if val, ok := l.Get(1); !ok || val != 2.2 {
		t.Errorf("Expected value 2.2, got %v", val)
	}

	if val := l.GetAll(); len(val) != 3 {
		t.Errorf("Expected 3 values, got %d", len(val))
	}

	if _, ok := l.Get(3); ok {
		t.Error("Expected no value at index 3")
	}

	if !l.Remove(1) {
		t.Error("Expected to remove value at index 1")
	}

	if val, ok := l.Get(1); !ok || val != 3.3 {
		t.Errorf("Expected value 3.3, got %v", val)
	}

	vl, id := l.Find(func(value float64) bool {
		return value == 1.1
	})
	if id == -1 {
		t.Error("Expected to find value 1.1, got -1")
	} else if vl != 1.1 {
		t.Errorf("Expected value 1.1, got %v", vl)
	}

	l.Clear()
	if l.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", l.Size())
	}
}

/*********************************************************
 ************Test Cases Using struct as Type*************
 *********************************************************/
type personDummy struct {
	Name string
	Age  int
}

// TestListStruct tests the list implementation with struct type.
// It checks adding, retrieving, removing elements, and clearing the list.
// It also verifies the size of the list at different stages.
func TestListStruct(t *testing.T) {
	l := newList[personDummy]()
	l.Add(personDummy{"Alice", 30})
	l.Add(personDummy{"Bob", 25})
	l.Add(personDummy{"Charlie", 35})

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	if val, ok := l.Get(1); !ok || val.Name != "Bob" {
		t.Errorf("Expected value Bob, got %v", val)
	}

	if val := l.GetAll(); len(val) != 3 {
		t.Errorf("Expected 3 values, got %d", len(val))
	}

	if _, ok := l.Get(3); ok {
		t.Error("Expected no value at index 3")
	}

	if !l.Remove(1) {
		t.Error("Expected to remove value at index 1")
	}

	if val, ok := l.Get(1); !ok || val.Name != "Charlie" {
		t.Errorf("Expected value Charlie, got %v", val)
	}

	vl, id := l.Find(func(value personDummy) bool {
		return value.Name == "Alice"
	})

	if id == -1 {
		t.Error("Expected to find value Alice, got -1")
	} else if vl.Name != "Alice" {
		t.Errorf("Expected value Alice, got %v", vl.Name)
	}

	vl, id = l.Find(func(value personDummy) bool {
		return value.Name == "NonExistent"
	})
	if id != -1 {
		t.Error("Expected not to find value NonExistent, got id", id)
	} else if vl.Name != "" {
		t.Errorf("Expected empty value, got %v", vl.Name)
	}

	l.Clear()
	if l.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", l.Size())
	}
}
