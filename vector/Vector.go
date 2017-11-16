package vector

import (
	"errors"
)

type Vector struct {
	Items []interface{}
}

func NewVector() *Vector {
	vector := &Vector{
		Items: make([]interface{}, 0),
	}
	return vector
}

func (v *Vector) Add(item interface{}) bool {
	v.Items = append(v.Items, item)
	return true
}

func (v *Vector) AddAll(items []interface{}) bool {
	for item := range items {
		v.Items = append(v.Items, item)
	}
	return true
}

func (v *Vector) AddElement(item interface{}) {
	v.Items = append(v.Items, item)
}

func (v *Vector) Clear() {
	v.Items = make([]interface{}, 0)
}

func (v *Vector) Clone() *Vector {
	newV := NewVector()
	for _, item := range v.Items {
		newV.Add(item)
	}
	return newV
}

func (v *Vector) Size() int {
	return len(v.Items)
}

func (v *Vector) Remove(i int) {
	v.Items = append(v.Items[:i], v.Items[i+1:]...)
}

func (v *Vector) RemoveRange(i, j int) {
	v.Items = append(v.Items[:i], v.Items[j+1:]...)
}

func (v *Vector) Get(index int) (interface{}, error) {
	if index < 0 || index > len(v.Items)-1 {
		return nil, errors.New("Vector Error: Index out of range.")
	}
	return v.Items[index], nil
}

func (v *Vector) ElementAt(index int) (interface{}, error) {
	if index < 0 || index > len(v.Items)-1 {
		return nil, errors.New("Vector Error: Index out of range.")
	}
	return v.Items[index], nil
}

func (v *Vector) InsertElementAt(item interface{}, index int) {
	sl := make([]interface{}, 0)
	sl = append(sl, item)

	if index < 0 || index > len(v.Items)-1 {
		return
	}
	v.Items = append(v.Items[:index], append(sl, v.Items[index:]...)...)
}

func (v *Vector) Set(index int, item interface{}) (interface{}, error) {
	if index < 0 || index > len(v.Items)-1 {
		return nil, errors.New("Vector Error: Index out of range.")
	}
	response := v.Items[index]
	v.Items[index] = item
	return response, nil
}
