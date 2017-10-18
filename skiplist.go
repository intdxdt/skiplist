package skiplist

import (
	"time"
	"math/rand"
	"github.com/intdxdt/cmp"
)

/*
 description insert, search , delete  is O( log n) with high probability
   |----(1)----(2)---------------------------------------(6)-----------|
   |----(1)----(2)------------------(3)-----------(6)----(6)-----------|
   |----(1)----(2)----(2)----(3)----(3)----(5)----(6)----(6)----(9)----|
 */

type SkipList struct {
	p          float64
	rand       *rand.Rand
	curlevel   int
	maxlevel   int
	duplicates bool
	minvalue   interface{}
	maxvalue   interface{}
	head       *Node
	tail       *Node
	vcut       []*Node
	compare    cmp.Compare
}

//NewSkipList - SkipList Constructor.
func NewSkipList(size int, duplicates bool, comparator cmp.Compare) *SkipList {
	var seed = rand.NewSource(time.Now().UnixNano())
	var rnd = rand.New(seed)
	var maxlevel = computeLevels(size)
	var minvalue = &minItem{}
	var maxvalue = &maxItem{}
	var head = mklevels(maxlevel, minvalue, maxvalue, comparator)
	var tail = head.list[0]
	var vcut = make([]*Node, len(head.list))

	return &SkipList{
		p:          0.5,
		rand:       rnd,
		curlevel:   0,
		duplicates: duplicates,
		maxlevel:   maxlevel,
		minvalue:   minvalue,
		maxvalue:   maxvalue,
		head:       head,
		tail:       tail,
		vcut:       vcut,
		compare:    comparator,
	}
}

//comparator
func (skp *SkipList) cmp(node *Node, value interface{}) int {
	if skp.ismax(node) || skp.ismin_item(value) {
		//max node > every value & every value > than min value
		return 1
	} else if skp.ismin(node) || skp.ismax_item(value) {
		//min node < every value & every value < than max value
		return -1
	}

	return node.cmp(node.value, value)
}

//computes the equlity of node and item
func (skp *SkipList) equal(cur *Node, v interface{}) bool {
	return skp.cmp(cur, v) == 0
}

//is min value
func (skp *SkipList) ismin(cur *Node) bool {
	return cur.value == skp.minvalue
}

//check if is min value.
func (skp *SkipList) ismin_item(cur interface{}) bool {
	return cur == skp.minvalue
}

//checks if is max value
func (skp *SkipList) ismax(cur *Node) bool {
	return cur.value == skp.maxvalue
}

//checks if is a max item.
func (skp *SkipList) ismax_item(cur interface{}) bool {
	return cur == skp.maxvalue
}

//First item in a skiplist.
func (skp *SkipList) First() interface{} {
	var cur = skp.head
	cur = cur.next(0)
	if skp.ismax(cur) {
		return nil
	}
	return cur.value
}

//Last value in skiplist.
func (skp *SkipList) Last() interface{} {
	var _, cur = skp.vertlist(skp.maxvalue, true, false)
	if skp.ismin(cur) {
		return nil
	}
	return cur.value
}

//IsEmpty checks if a skiplist is empty.
func (skp *SkipList) IsEmpty() bool {
	return skp.First() == nil
}
