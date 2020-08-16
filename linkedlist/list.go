package main

import (
	"errors"
)

type List struct {
	head *Node
}

type Node struct {
	value int
	next *Node
}

func initList() *List {
	return &List{nil}
}

func (list *List) add(value int) {
	if list.head == nil {
		list.head = &Node{value, nil}
	} else {
		walkNode := list.head
		for walkNode.next != nil {
			walkNode = walkNode.next
		}
		walkNode.next = &Node{value, nil}
	}
}

func (list *List) get(position int) (*Node, error) {
	if position < 0 {
		return nil, errors.New("ops")
	} else if position == 0 {
		return list.head, nil
	}

	walkNode := list.head
	cont := 0
	for walkNode != nil {
		if position == cont {
			return walkNode, nil
		}
		walkNode = walkNode.next
		cont++
	}

	return nil, errors.New("ops")
}

func (list *List) remove(position int) error{
	if position == 0 {
		list.head = list.head.next
	} else {
		prev, err := list.get(position - 1)
		if err != nil {
			return errors.New("ops")
		}
		prev.next = prev.next.next
	}

	return nil
}

func main() {

	lista := initList()
	lista.add(1)
	lista.add(2)
	lista.remove(1)
	lista.add(3)

}
