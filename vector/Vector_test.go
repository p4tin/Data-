package vector

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

func TestVectorAddAll_OK(t *testing.T) {
	v := NewVector()

	ints := []int{1, 2, 3, 4}
	faces := make([]interface{}, len(ints))
	for i, s := range ints {
		faces[i] = s
	}

	ret := v.AddAll(faces)
	if ret != true {
		t.Errorf("Add return value should have been true but it was %b\n", ret)
	}
	if len(v.Items) != 4 {
		t.Errorf("The size of the vector should have been 4 and it was %d\n", len(v.Items))
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

func TestVectorClone(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	nv := v.Clone()

	if len(v.Items) != len(nv.Items) {
		t.Errorf("The size of the vector should have been 3 and it was %d\n", len(nv.Items))
	}

	obj, _ := nv.Get(1)
	if obj.(int) != 2 {
		t.Errorf("Clone-Get should have return int (2) but instead was: %+v\n", obj)
	}
}

func TestVectorRemove_OK(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	v.Remove(1)
	obj, _ := v.Get(1)
	if obj.(int) != 3 {
		t.Errorf("Get should have return int (3) but instead was: %+v\n", obj)
	}
}

func TestVectorRemoveRange_OK(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	v.AddElement(4)
	v.RemoveRange(1, 2)
	obj, _ := v.Get(1)
	if obj.(int) != 4 {
		t.Errorf("Get should have return int (4) but instead was: %+v\n", obj)
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
	obj, _ := v.ElementAt(1)
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

func TestVectorInsertElementAt_OutOfRange(t *testing.T) {
	v := NewVector()
	v.AddElement(1)
	v.AddElement(2)
	v.AddElement(3)
	v.InsertElementAt(99, 1)
	if len(v.Items) != 4 {
		t.Errorf("The size of the vector should have been 4 and it was %d\n", len(v.Items))
	}
	obj, _ := v.Get(1)
	if obj.(int) != 99 {
		t.Errorf("Get should have return int (99) but instead was: %+v\n", obj)
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
