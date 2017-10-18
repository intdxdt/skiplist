package skiplist

// vertical list of nodes
func (skp *SkipList) vertlist(value interface{}, first, memo bool) ([]*Node, *Node) {
	var cur = skp.head
	var l = skp.curlevel
	var list []*Node
	if memo {
		list = skp.vcut //make([]*Node, len(cur.list))
	}

	var next *Node
	for l >= 0 {
		next = cur.next(l)
		for skp.precedes(next, value, first) {
			cur = next
			next = cur.next(l)
		}

		if memo {
			list[l] = cur
		}
		l += -1
	}

	return list, cur
}
