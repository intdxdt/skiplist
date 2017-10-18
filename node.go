package skiplist

import "github.com/intdxdt/cmp"

type sentinel struct{}

func (stnl *sentinel) Compare(o interface{}) int {
	return -1
}

func (stnl *sentinel) String() string {
	return ""
}

type minItem struct {
	sentinel
}

//Max Item
type maxItem struct {
	sentinel
}

//Compare max item, max is greater than everything in the skiplist
func (mxitm *maxItem) Compare(o interface{}) int {
	return 1
}

//Node type for a skiplist (value and level as list)
type Node struct {
	value interface{}
	list  []*Node
	cmp   cmp.Compare
}

//NewNode constructor.
func NewNode(level int, value interface{}, comparator cmp.Compare) *Node {
	return &Node{
		value: value,
		list:  make([]*Node, level+1),
		cmp:   comparator,
	}
}

//next
func (n *Node) next(level int) *Node {
	return n.list[level]
}
