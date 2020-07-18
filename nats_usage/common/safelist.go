package common

import (
	"errors"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type ISafeList interface {
	Push(data interface{})
	Pop() (interface{}, error)
	Signal() chan bool
}

type _SafeListNode struct {
	next  unsafe.Pointer
	value interface{}
}

func newNode(data interface{}) unsafe.Pointer {
	return unsafe.Pointer(&_SafeListNode{
		next:  nil,
		value: data,
	})
}

type _SafeList struct {
	head unsafe.Pointer
	tail unsafe.Pointer

	C chan bool
}

func NewSafeList() ISafeList {
	// at least one node ( head is never nil )
	node := unsafe.Pointer(newNode(nil))
	return &_SafeList{
		head: node,
		tail: node,
		C:    make(chan bool, 1),
	}
}

func (sl *_SafeList) Push(data interface{}) {
	newNode := newNode(data)
	var tail unsafe.Pointer

	for {
		tail = sl.tail
		next := (*_SafeListNode)(tail).next

		if next != nil {
			atomic.CompareAndSwapPointer(&sl.tail, tail, next)
		} else {
			if atomic.CompareAndSwapPointer(&(*_SafeListNode)(sl.tail).next, nil, newNode) {
				break
			}
		}
		runtime.Gosched()
	}

	atomic.CompareAndSwapPointer(&sl.tail, tail, newNode)

	if len(sl.C) == 0 {
		sl.C <- true
	}
}

var errNoNode = errors.New("safe list. no node")

func (sl *_SafeList) Pop() (interface{}, error) {
	for {
		head := sl.head
		tail := sl.tail

		next := (*_SafeListNode)(head).next

		if head == tail {
			if next == nil {
				return nil, errNoNode
			}
			atomic.CompareAndSwapPointer(&sl.tail, tail, next)
		} else {
			if atomic.CompareAndSwapPointer(&sl.head, head, next) {
				return (*_SafeListNode)(next).value, nil
			}
		}

		runtime.Gosched()
	}
}

func (sl *_SafeList) Signal() chan bool {
	return sl.C
}

func (sl *_SafeList) IsEmpty() bool {
	head := sl.head
	tail := sl.tail

	next := (*_SafeListNode)(head).next
	if head == tail {
		if next == nil {
			return true
		}
	}

	return false
}
