package main

type Queue interface {
	EnQueue(v interface{}) bool
	DeQueue() interface{}
	Print() string
}
