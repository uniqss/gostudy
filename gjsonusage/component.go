package main

type Component struct {
	ComponentName string
	Properties []*ComponentProperty
	Tables []*ComponentTable
	Containers []*ComponentContainer
}
