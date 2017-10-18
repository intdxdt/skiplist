package skiplist

//Insert item
func (skp *SkipList) Insert(value interface{}) *SkipList {

	var cutlist, _ = skp.vertlist(value, false, true) // last occurrance
	var cur = cutlist[0]
	var l int
	//cur.next
	cur = cur.list[0]

	//if not dups, cur == value
	if skp.equal(cur, value) {
		cur.value = value
	} else {
		// generate random level
		var rlevel = skp.randomlevel()
		// random level higher than current level
		if rlevel > skp.curlevel {
			// fill above current level with head
			l = skp.curlevel + 1
			for l <= rlevel {
				cutlist[l] = skp.head
				l += 1
			}
			// update level
			skp.curlevel = rlevel
		}
		// update levels from 0 to random level
		l = 0
		cur = NewNode(rlevel, value, skp.compare)
		for l <= rlevel {
			// cur.next -> prev.next
			cur.list[l] = cutlist[l].list[l]
			// prev.next -> cur
			cutlist[l].list[l] = cur
			// next level up
			l += 1
		}
	}
	return skp
}
