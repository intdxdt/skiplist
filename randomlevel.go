package skiplist

//random level
func (skp *SkipList) randomlevel() int {
    var skpl = skp
    // 0 based indexing
    var index = skpl.maxlevel - 1
    var level = 0
    // see fig  5 of orig. paper
    for skp.rand.Float64() > skpl.p && level < index {
        level += 1
    }
    return level
}
