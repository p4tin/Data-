package DataStructure

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

func (v *Vector)Add(item interface{}) bool {
	v.Items = append(v.Items, item)
	return true
}

func (v *Vector)AddElement(item interface{}) {
	v.Items = append(v.Items, item)
}

func (v *Vector)Clear() {
	v.Items = make([]interface{}, 0)
}


func (v *Vector)Size() int {
	return len(v.Items)
}

func (v *Vector)Get(index int) (interface{}, error) {
	if index < 0 || index > len(v.Items)-1 {
		return nil, errors.New("Vector Error: Index out of range.")
	}
	return v.Items[index], nil
}

func (v *Vector)ElementAt(index int) (interface{}, error) {
	if index < 0 || index > len(v.Items)-1 {
		return nil, errors.New("Vector Error: Index out of range.")
	}
	return v.Items[index], nil
}

func (v *Vector)Set(index int, item interface{}) (interface{}, error) {
	if index < 0 || index > len(v.Items)-1 {
		return nil, errors.New("Vector Error: Index out of range.")
	}
	response := v.Items[index]
	v.Items[index] = item
	return response, nil
}


//AddAll([]interface{})
//Clone()
//elementAt(int index)
//insertElementAt(E obj, int index)
//remove(int index)


