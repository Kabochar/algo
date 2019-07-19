package main

import (
	"errors"
	"fmt"
)

//数组的数据插入，指定下标访问，指定下标删除

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (r *Array) Len() uint {
	return r.length
}

func (r *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(r.data)) {
		return true
	}
	return false
}

func (r *Array) FindFromIndex(index uint) (int, error) {
	if r.isIndexOutOfRange(index) {
		return 0, errors.New("Out of range!")
	}

	return r.data[index], nil
}

func (r *Array) InsertToIndex(index uint, value int) error {
	if r.Len() == uint(cap(r.data)) {
		return errors.New("Full array")
	}
	if index != r.length && r.isIndexOutOfRange(index) {
		return errors.New("Out of range")
	}

	for i := r.length; i > index; i-- {
		r.data[i] = r.data[i-1]
	}
	r.data[index] = value
	r.length++

	return nil
}

func (r *Array) InsertToTail(v int) error {
	return r.InsertToIndex(r.Len(), v)
}

func (r *Array) DeleteFromIndex(index uint) (int, error) {
	if r.isIndexOutOfRange(index) {
		return 0, errors.New("Out of range")
	}
	v := r.data[index]
	for i := index; i < r.Len()-1; i++ {
		r.data[i] = r.data[i+1]
	}
	r.length--

	return v, nil
}

func (r *Array) Print() {
	var format string
	for i := uint(0); i < r.Len(); i++ {
		format += fmt.Sprintf("|%+v", r.data[i])
	}
	fmt.Println(format)
}
