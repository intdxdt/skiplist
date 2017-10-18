package skiplist
//Empty a skplist.
func (skp *SkipList) Empty() *SkipList{
  for skp.curlevel >= 0 {
    skp.head.list[skp.curlevel] = skp.tail
    skp.curlevel += -1
  }
  //reset level
  skp.curlevel = 0
  return skp
}
