package skiplist

import (
	"strings"
	"fmt"
)

//String returns skiplist as string
func (skp *SkipList) String() string {
	var levelindex = make([][]*Node, skp.curlevel+1)
	var head = skp.head
	var cur = head

	for i := 0; i <= skp.curlevel; i++ {
		levelindex[i] = make([]*Node, 0)

		for !skp.ismax(cur) {
			cur = cur.list[i]

			if !skp.ismax(cur) {

				if i == 0 {
					levelindex[i] = append(levelindex[i], cur)
				} else {
					//init
					if len(levelindex[i]) == 0 && cap(levelindex[i]) == 0 {
						levelindex[i] = make([]*Node, len(levelindex[0]), cap(levelindex[0]))
					}
					for j := 0; j < len(levelindex[0]); j++ {
						if levelindex[0][j] == cur {
							levelindex[i][j] = cur
							break
						}
					}
				}

			}

		}
		cur = head
	}

	var dashline = "----"

	//else
	var printlist = make([]string, 0)
	var printlvl []string
	for l := 0; l < len(levelindex); l++ {

		var lvl = make([]interface{}, len(levelindex[l]))
		for i := 0; i < len(levelindex[l]); i++ {
			cur = levelindex[l][i]
			if cur != nil {
				lvl[i] = cur.value
			}

		}

		if len(printlvl) == 0 {
			var width = make([]int, len(lvl))
			printlvl = make([]string, len(lvl))
			for i := 0; i < len(lvl); i++ {
				width[i] = len(fmt.Sprintf("%v", lvl[i])) + 2
			}
			for i := 0; i < len(width); i++ {
				printlvl[i] = strings.Repeat("-", width[i])
			}
		}

		var ln = make([]string, len(printlvl))
		copy(ln, printlvl)
		for j, v := range lvl {
			if v != nil {
				ln[j] = "(" + fmt.Sprintf("%v", v) + ")"
			}
		}

		ln = unshift(ln, "|")
		ln = append(ln, "|")
		printlist = append(printlist, strings.Join(ln, dashline))
	}

	printlist = reverse(printlist)
	return strings.Join(printlist, "\n")
}

//unshift string into list
func unshift(ln []string, x string) []string {
	ln = append([]string{x}, ln...)
	return ln
}

//reverses a string
func reverse(a []string) []string {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return a
}
