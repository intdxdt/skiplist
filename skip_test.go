package skiplist

import (
	"fmt"
	"testing"
	"github.com/intdxdt/cmp"
	"github.com/franela/goblin"
)

func TestSkipList(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("SkipList - Sentinel , Min , Max Item", func() {
		var s = sentinel{}
		var max = maxItem{}
		var min = minItem{}
		g.Assert(s.Compare(100)).Equal(-1)
		g.Assert(min.Compare(100)).Equal(-1)
		g.Assert(max.Compare(100)).Equal(1)

		g.Assert(s.String()).Equal("")
		g.Assert(min.String()).Equal("")
		g.Assert(max.String()).Equal("")
	})

	g.Describe("SkipList - With Duplicates", func() {
		var array = []int{9, 5, 3, 2, 6, 6, 6, 6, 1, 2, 3}
		var obj = NewSkipList(len(array), true, cmp.Int)
		for j := 0; j < len(array); j++ {
			obj.Insert(array[j])
		}
		//print
		fmt.Println(obj)

		g.It("should test fist & last in list", func() {
			g.Assert(obj.First()).Equal(1)
			g.Assert(obj.Last()).Equal(9)
		})
		g.It("should test Delete & Search & Iter item from list", func() {
			var f = 6
			g.Assert(obj.Search(f)).Equal(6)
			obj.Remove(f)
			g.Assert(obj.Search(f)).Equal(6)
			obj.Remove(f)
			g.Assert(obj.Search(f)).Equal(6)
			obj.Remove(f)
			g.Assert(obj.Search(f)).Equal(6)
			f = 10
			g.Assert(obj.Search(f)).Equal(nil)

			//print
			fmt.Println(obj)

			var list = make([]int, 0) //
			obj.Each(func(o interface{}, i int) {
				list = append(list, o.(int))
			})
			g.Assert(list).Eql([]int{1, 2, 2, 3, 3, 5, 6, 9})
			var expect_evens = make([]int, 0)
			var isven = func(o interface{}, i int) bool {
				v := o.(int)
				bln := v%2 == 0
				if bln {
					expect_evens = append(expect_evens, v)
				}
				return bln
			}
			var evens = obj.Filter(isven)
			g.Assert(len(evens)).Equal(3)
			g.Assert(expect_evens).Eql([]int{2, 2, 6})

			obj.Empty()
			list = make([]int, 0)
			obj.Each(func(o interface{}, i int) {
				list = append(list, o.(int))
			})
			g.Assert(len(list)).Equal(0)
			g.Assert(obj.curlevel).Equal(0)
			g.Assert(obj.First()).Equal(nil)
			g.Assert(obj.Last()).Equal(nil)
			//print
			fmt.Println("afeter emtpy:\n", obj)

		})

	})

	g.Describe("SkipList - Set", func() {
		var array = []int{9, 5, 3, 2, 6, 6, 6, 6, 1, 2, 3}
		var obj = NewSkipList(len(array), false, cmp.Int)

		g.It("should test fist & last in list", func() {
			g.Assert(obj.First()).Equal(nil)
			g.Assert(obj.Last()).Equal(nil)
			for j := 0; j < len(array); j++ {
				obj.Insert(array[j])
			}
			//print
			fmt.Println(obj)
			g.Assert(obj.First()).Equal(1)
			g.Assert(obj.Last()).Equal(9)
		})

		g.It("should test Delete & Search & ToArray item from list", func() {
			var f = 6
			obj.Remove(f)
			g.Assert(obj.Search(f)).Equal(nil)
			f = 10
			g.Assert(obj.Search(f)).Equal(nil)
			//print
			fmt.Println(obj)

			var list = make([]int, 0) //
			obj.Each(func(o interface{}, i int) {
				list = append(list, o.(int))
			})
			g.Assert(list).Eql([]int{1, 2, 3, 5, 9})
			obj.Empty()
			list = make([]int, 0)
			obj.Each(func(o interface{}, i int) {
				list = append(list, o.(int))
			})
			g.Assert(len(list)).Equal(0)
			g.Assert(obj.curlevel).Equal(0)
			g.Assert(obj.IsEmpty()).Equal(true)

			//print
			fmt.Println("after emtpy:\n", obj)
		})

	})

	g.Describe("SkipList - Set -Special Case Delete", func() {
		for {
			var array = []int{9, 5, 3, 2, 6, 6, 6, 6, 1, 2, 3}
			var obj = NewSkipList(len(array), false, cmp.Int)
			for j := 0; j < len(array); j++ {
				obj.Insert(array[j])
			}
			top := obj.curlevel
			if top == 0 {
				continue
			}
			cur := obj.head.next(top)
			next := cur.next(top)
			if obj.ismax(next) {
				obj.Remove(cur.value)
				g.Assert(obj.curlevel <= (top - 1)).IsTrue()
				break
			}
		}
	})
}
