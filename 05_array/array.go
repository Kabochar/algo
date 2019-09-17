package main

import (
	"errors"
	"fmt"
)

// 数组的数据插入，指定下标访问，指定下标删除

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

// 是否超出容量
func (r *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(r.data)) {
		return true
	}
	return false
}

// 根据下标寻找
func (r *Array) FindFromIndex(index uint) (int, error) {
	if r.isIndexOutOfRange(index) {
		return 0, errors.New("Out of range!")
	}

	return r.data[index], nil
}

// 插入数据到指定下标
func (r *Array) InsertToIndex(index uint, value int) error {
	// 数组无空余位置
	if r.Len() == uint(cap(r.data)) {
		return errors.New("Full array")
	}
	// 插入下标超限
	if index != r.length && r.isIndexOutOfRange(index) {
		return errors.New("Out of range")
	}

	// 移动数据，将指定下标以后的数据往后挪一位
	for i := r.length; i > index; i-- {
		r.data[i] = r.data[i-1]
	}
	r.data[index] = value
	r.length++

	return nil
}

// 数据插入到数组末尾
func (r *Array) InsertToTail(v int) error {
	return r.InsertToIndex(r.Len(), v)
}

// 从指定下标位置删除
func (r *Array) DeleteFromIndex(index uint) (int, error) {
	if r.isIndexOutOfRange(index) {
		return 0, errors.New("Out of range")
	}
	value := r.data[index]
	for i := index; i < r.Len()-1; i++ {
		r.data[i] = r.data[i+1]
	}
	r.length--

	return value, nil
}

// 打印
func (r *Array) Print() {
	var format string
	for i := uint(0); i < r.Len(); i++ {
		format += fmt.Sprintf("|%+v", r.data[i])
	}
	fmt.Println(format)
}
