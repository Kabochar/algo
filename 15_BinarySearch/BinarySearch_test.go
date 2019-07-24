package main

import "testing"

func TestBinarySearch(t *testing.T) {
	var arr []int

	arr = []int{1, 3, 5, 6, 8}
	if BinarySearch(arr, 8) != 4 {
		t.Fatal(BinarySearch(arr, 3))
	}
	if BinarySearch(arr, 4) != -1 {
		t.Fatal(BinarySearch(arr, 4))
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	var arr []int

	arr = []int{1, 3, 5, 6, 8}
	if BinarySearchRecursive(arr, 8) != 4 {
		t.Fatal(BinarySearch(arr, 3))
	}
	if BinarySearchRecursive(arr, 4) != -1 {
		t.Fatal(BinarySearch(arr, 4))
	}
}

func TestBinarySearchFirst(t *testing.T) {
	var arr []int

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirst(arr, 2) != 1 {
		t.Fatal(BinarySearchFirst(arr, 2))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirst(arr, 3) != 4 {
		t.Fatal(BinarySearchFirst(arr, 3))
	}
}

func TestBinarySearchLast(t *testing.T) {
	var arr []int

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLast(arr, 2) != 3 {
		t.Fatal(BinarySearchLast(arr, 2))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLast(arr, 3) != 4 {
		t.Fatal(BinarySearchLast(arr, 3))
	}
}

func TestBinarySearchFirstGT(t *testing.T) {
	var arr []int

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(arr, 2) != 4 {
		t.Fatal(BinarySearchFirstGT(arr, 2))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(arr, 1) != 1 {
		t.Fatal(BinarySearchFirstGT(arr, 1))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(arr, 3) != 5 {
		t.Fatal(BinarySearchFirstGT(arr, 3))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(arr, 4) != -1 {
		t.Fatal(BinarySearchFirstGT(arr, 4))
	}
}

func TestBinarySearchLastLT(t *testing.T) {
	var arr []int

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(arr, 2) != 0 {
		t.Fatal(BinarySearchLastLT(arr, 2))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(arr, 1) != -1 {
		t.Fatal(BinarySearchLastLT(arr, 1))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(arr, 3) != 3 {
		t.Fatal(BinarySearchLastLT(arr, 3))
	}

	arr = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(arr, 4) != 4 {
		t.Fatal(BinarySearchLastLT(arr, 4))
	}
}
