package util

import "fmt"

type CompressedList struct {
	startNode *ListNode
	length    int
	currID    int
}

type ListNode struct {
	nextNode        *ListNode
	prevNode        *ListNode
	size, index, id int
	allocated       bool
}

func (l *ListNode) Size() int {
	return l.size
}

func (l *ListNode) Index() int {
	return l.index
}

func (l *ListNode) ID() int {
	return l.id
}

func CreateCompressedList(length int) (*CompressedList, error) {
	if length < 0 {
		return nil, fmt.Errorf("cannot create compressed list with length < 0")
	}

	node := &ListNode{
		size:      length,
		index:     0,
		id:        0,
		allocated: false,
	}
	node.nextNode = node
	node.prevNode = node

	return &CompressedList{
		startNode: node,
		length:    length,
	}, nil
}

func (c *CompressedList) Append(size int) {
	c.length += size
	// If the list only contains a single node then just extend the size
	if c.startNode.prevNode == c.startNode {
		return
	}

	// Else create a new node containing the new space
	freeNode := &ListNode{
		nextNode:  c.startNode,
		prevNode:  c.startNode.prevNode,
		size:      size,
		index:     c.length - size,
		id:        c.currID,
		allocated: false,
	}

	c.currID++
	c.startNode.prevNode.nextNode = freeNode
}

func (c *CompressedList) Allocate(size int) (*ListNode, error) {
	if size > c.length {
		return nil, fmt.Errorf("cannot allocate for size: %v greater than length: %v", size, c.length)
	}

	// Iterate from the start of the list until we find a spot
	// which is not allocated and is of our required size
	totalLength := 0
	node := c.startNode
	for totalLength+size <= c.length {
		// Skip if this node is already allocated
		if node.allocated || node.size < size {
			totalLength += node.size
			node = node.nextNode
			continue
		}

		// We have found a free node with enough size
		// Create a free node with remaining space
		c.currID++
		freeNode := &ListNode{
			nextNode:  node.nextNode,
			prevNode:  node,
			size:      node.size - size,
			index:     totalLength + size,
			id:        c.currID,
			allocated: false,
		}

		// Update the existing node
		node.size = size
		node.nextNode = freeNode
		node.allocated = true
		// If we are the nextNode then freeNode is the prevNode
		if freeNode.nextNode == node {
			node.prevNode = freeNode
		}

		return node, nil
	}

	return nil, fmt.Errorf("unable to allocate for size: %v", size)
}

func (node *ListNode) Free() {
	// Re-attach us and the following node if it is not allocated
	// Detach the following node to be garbage collected
	if nextNode := node.nextNode; !nextNode.allocated {
		node.size += node.nextNode.size
		node.allocated = false
		node.nextNode = nextNode.nextNode

		nextNode.nextNode = nil
		nextNode.prevNode = nil
	} else if prevNode := node.prevNode; !prevNode.allocated {
		// If this isn't then case then check if the previous node is un-allocated and repeat
		prevNode.size += node.size
		prevNode.nextNode = nextNode

		node.nextNode = nil
		node.prevNode = nil
	} else {
		// Otherwise just remove our allocation
		node.allocated = false
	}
}

func (c *CompressedList) Clear() {
	l := 0
	node := c.startNode
	node.id = 0
	for l < c.length {
		l += node.size
		node.Free()
		node = node.nextNode
	}
	c.currID = 0
}
