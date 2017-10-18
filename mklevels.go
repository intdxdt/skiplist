package skiplist

import "github.com/intdxdt/cmp"

//make levels
func mklevels(maxlevel int, minval, maxval interface{}, comparator cmp.Compare) *Node {
	maxlevel += -1 //zero index
	var head = NewNode(maxlevel, minval, comparator)
	var tail = NewNode(maxlevel, maxval, comparator)
	for maxlevel >= 0 {
		head.list[maxlevel] = tail
		maxlevel += -1
	}
	return head
}
