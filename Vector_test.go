package DataStructure

import (
	"testing"
)

func TestVectorAdd_OK(t *testing.T) {
	v := NewVector()
	ret := v.Add(1)
	if ret != true {
		t.Errorf("Add return value should have been true but it was %b\n", ret)
	}
	if len(v.Items) != 1 {
		t.Errorf("The size of the vector should have been 1 and it was %d\n", len(v.Items))
	}
}

func TestVectorAddElement_OK(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	if len(v.Items) != 1 {
		t.Errorf("The size of the vector should have been 1 and it was %d\n", len(v.Items))
	}
}

func TestVectorClear(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.Clear()
	if len(v.Items) != 0 {
		t.Errorf("The size of the vector should have been 0 and it was %d\n", len(v.Items))
	}
}

func TestVectorSize(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	s := v.Size()
	if s != 1 {
		t.Errorf("The size of the vector should have been 1 and it was %d\n", s)
	}
}

func TestVectorGet_OK(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	obj, _ := v.Get(1)
	if obj.(int) != 2 {
		t.Errorf("Get should have return int (2) but instead was: %+v\n", obj)
	}
}

func TestVectorGet_OutOfRange(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	obj, err := v.Get(3)
	if err == nil {
		t.Errorf("Get should have return an error but it was nil and the object was: %+v\n", obj)
	}
}

func TestVectorElementAt_OK(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	obj, _ := v.Get(1)
	if obj.(int) != 2 {
		t.Errorf("Get should have return int (2) but instead was: %+v\n", obj)
	}
}

func TestVectorElementAt_OutOfRange(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	obj, err := v.ElementAt(3)
	if err == nil {
		t.Errorf("Get should have return an error but it was nil and the object was: %+v\n", obj)
	}
}

func TestVectorSet_OK(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	obj, _ := v.Set(1, 6)
	if obj.(int) != 2 {
		t.Errorf("Get should have return int (2) but instead was: %+v\n", obj)
	}
	obj, _ = v.Get(1)
	if obj.(int) != 6 {
		t.Errorf("Get should have return int (6) but instead was: %+v\n", obj)
	}
}

func TestVectorSet_OutOfRange(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	obj, err := v.Set(3, 6)
	if err == nil {
		t.Errorf("Get should have return an error but it was nil and the object was: %+v\n", obj)
	}
}
