package skiplist

//Search skiplist for item.
func (skp *SkipList) Search(value interface{}) interface{} {
	var _, cur = skp.vertlist(value, true, false)
	cur = cur.next(0)
	if skp.equal(cur, value) {
		return cur.value
	}
	return nil
}
