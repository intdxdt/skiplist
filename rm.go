package skiplist


//Remove item from skiplist.
func (skp *SkipList) Remove(value interface{}) {
    var vlist, _ = skp.vertlist(value, true, true) //first occurrance
    var cur = vlist[0]
    var l int
    cur = cur.next(0)
    if skp.equal(cur, value) {
        //drop from the list
        l = 0   //bottom up
        for l <= skp.curlevel {
            if vlist[l].list[l] == cur {
                vlist[l].list[l] = cur.next(l)
            }
            l += 1
        }
        /*
         check if we have to lower level after removing 6 top row is empty
        |----------------------------------------------(6)--------------------------------|
        |----------------------------------------------(6)-----------(6)------------------|
        |----(1)----(2)----(2)----(3)----(3)----(5)----(6)----(6)----(6)----(6)----(9)----|
         */
        for skp.curlevel > 0 && skp.ismax(skp.head.list[skp.curlevel]) {
            skp.curlevel += -1
        }
    }
}
