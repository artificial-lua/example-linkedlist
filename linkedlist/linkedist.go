package linkedlist

import (
	"errors"
	"fmt"
)

type singleLinkedNode struct {
	value any
	next  *singleLinkedNode
}

type singleLinkedList struct {
	head *singleLinkedNode
	tail *singleLinkedNode
}

type doubleLinkedNode struct {
	value any
	next  *doubleLinkedNode
	prev  *doubleLinkedNode
}

type doubleLinkedList struct {
	head *doubleLinkedNode
	tail *doubleLinkedNode
}

func (s *singleLinkedNode) GetValue() any {
	return s.value
}

func (d *doubleLinkedNode) GetValue() any {
	return d.value
}

func (s *singleLinkedNode) getNextLink() *singleLinkedNode {
	if s.next == nil {
		return nil
	}
	return s.next
}

func (d *doubleLinkedNode) getNextLink() *doubleLinkedNode {
	if d.next == nil {
		return nil
	}
	return d.next
}

func (d *doubleLinkedNode) getPrevLink() *doubleLinkedNode {
	if d.prev == nil {
		return nil
	}
	return d.prev
}

func (s *singleLinkedList) GetHead() *singleLinkedNode {
	if s.head == nil {
		return nil
	}
	return s.head
}

func (s *singleLinkedList) GetTail() *singleLinkedNode {
	if s.tail == nil {
		return nil
	}
	return s.tail
}

func (d *doubleLinkedList) GetHead() *doubleLinkedNode {
	if d.head == nil {
		return nil
	}
	return d.head
}

func (d *doubleLinkedList) GetTail() *doubleLinkedNode {
	if d.tail == nil {
		return nil
	}
	return d.tail
}

func (s *singleLinkedList) Append(value any) {
	node := &singleLinkedNode{value: value}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = node
	}
}

func (d *doubleLinkedList) Append(value any) {
	node := &doubleLinkedNode{value: value}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		d.tail.next = node
		node.prev = d.tail
		d.tail = node
	}
}

func (s *singleLinkedList) Prepend(value any) {
	node := &singleLinkedNode{value: value}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		node.next = s.head
		s.head = node
	}
}

func (d *doubleLinkedList) Prepend(value any) {
	node := &doubleLinkedNode{value: value}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
}

var errIndexOutOfRange = errors.New("index out of range")

func (s *singleLinkedNode) traverse(index int) (*singleLinkedNode, error) {
	if index == 0 {
		return s, nil
	}
	if s.next == nil {
		return nil, errIndexOutOfRange
	}
	return s.next.traverse(index - 1)
}

func (d *doubleLinkedNode) traverse(index int) (*doubleLinkedNode, error) {
	if index == 0 {
		return d, nil
	}
	if d.next == nil {
		return nil, errIndexOutOfRange
	}
	return d.next.traverse(index - 1)
}

func (s *singleLinkedList) SearchByIndex(index int) (*singleLinkedNode, error) {
	if s.head == nil {
		return nil, errIndexOutOfRange
	}
	return s.head.traverse(index)
}

func (d *doubleLinkedList) SearchByIndex(index int) (*doubleLinkedNode, error) {
	if d.head == nil {
		return nil, errIndexOutOfRange
	}
	return d.head.traverse(index)
}

func (s *singleLinkedList) Insert(index int, value any) error {
	node := &singleLinkedNode{value: value}
	if index == 0 {
		s.Prepend(node)
		return nil
	}
	prevNode, err := s.SearchByIndex(index - 1)
	if err != nil {
		return err
	}
	nextNode := prevNode.getNextLink()
	prevNode.next = node
	node.next = nextNode
	return nil
}

func (d *doubleLinkedList) Insert(index int, value any) error {
	node := &doubleLinkedNode{value: value}
	if index == 0 {
		d.Prepend(node)
		return nil
	}
	prevNode, err := d.SearchByIndex(index - 1)
	if err != nil {
		return err
	}
	nextNode := prevNode.getNextLink()
	prevNode.next = node
	node.prev = prevNode
	node.next = nextNode
	return nil
}

func (s *singleLinkedList) Pop(index int) (any, error) {
	if index == 0 {
		popped := s.head
		s.head = s.head.next
		return popped.value, nil
	}
	prevNode, err := s.SearchByIndex(index - 1)
	if err != nil {
		return nil, err
	}
	popped := prevNode.getNextLink()
	nextNode := popped.getNextLink()
	prevNode.next = nextNode
	return popped.value, nil
}

func (d *doubleLinkedList) Pop(index int) (any, error) {
	if index == 0 {
		popped := d.head
		d.head = d.head.next
		d.head.prev = nil
		return popped.value, nil
	}
	prevNode, err := d.SearchByIndex(index - 1)
	if err != nil {
		return nil, err
	}
	popped := prevNode.getNextLink()
	nextNode := popped.getNextLink()
	prevNode.next = nextNode
	if nextNode != nil {
		nextNode.prev = prevNode
	}
	return popped.value, nil
}

func (s *singleLinkedList) Update(index int, value any) error {
	node, err := s.SearchByIndex(index)
	if err != nil {
		return err
	}
	node.value = value
	return nil
}

func (d *doubleLinkedList) Update(index int, value any) error {
	node, err := d.SearchByIndex(index)
	if err != nil {
		return err
	}
	node.value = value
	return nil
}

func (s *singleLinkedList) String() string {
	if s.head == nil {
		return ""
	}
	result := "["
	for node := s.head; node != nil; node = node.next {
		result += fmt.Sprintf("%v", node.value)
		if node.next != nil {
			result += ", "
		}
	}
	result += "]"
	return result
}

func (d *doubleLinkedList) String() string {
	if d.head == nil {
		return ""
	}
	result := "["
	for node := d.head; node != nil; node = node.next {
		result += fmt.Sprintf("%v", node.value)
		if node.next != nil {
			result += ", "
		}
	}
	result += "]"
	return result
}

func NewSingleLinkedList() *singleLinkedList {
	return &singleLinkedList{}
}

func NewDoubleLinkedList() *doubleLinkedList {
	return &doubleLinkedList{}
}
