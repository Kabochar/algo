package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 5
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.InsertToIndex(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	arr.InsertToIndex(uint(1), 999)
	arr.Print()

	arr.InsertToTail(666)
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.InsertToIndex(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.DeleteFromIndex(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.InsertToIndex(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.FindFromIndex(0))
	t.Log(arr.FindFromIndex(9))
	t.Log(arr.FindFromIndex(11))
}
