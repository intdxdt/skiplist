package skiplist

//Each iterates over each item in skiplist with call to
// func (item Item, i int)
func (skp *SkipList) Each(fn func(interface{}, int)) {
	var cur = skp.head
	var i int
	cur = cur.next(0)
	for !skp.ismax(cur) {
		fn(cur.value, i)
		cur = cur.next(0)
		i += 1
	}
}

//Filters items based on predicate : func (item Item, i int) bool
func (skp *SkipList) Filter(fn func(item interface{}, i int) bool) []interface{} {
	var items = make([]interface{}, 0)
	skp.Each(func(v interface{}, i int) {
		if fn(v, i) {
			items = append(items, v)
		}
	})
	return items
}
