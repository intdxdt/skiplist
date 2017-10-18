package skiplist

//item precedes value by insert support stable insert for duplicates
func (skp *SkipList) precedes(cur *Node, v interface{}, first bool) bool {
	var comparison = skp.cmp(cur, v)
	if !first && skp.duplicates {
		return comparison <= 0
	}
	return comparison < 0
}
