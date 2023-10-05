package goslices

import (
	"cmp"
	"slices"
)

type Counter[V cmp.Ordered] struct {
	Total               int
	Set                 map[V]int
	ordered             bool
	orderedList         []CounterItem[V]
	olLen               int
	reversed            bool
	reversedOrderedList []CounterItem[V]
	rolLen              int
}

type CounterItem[V cmp.Ordered] struct {
	Value V
	Count int
}

// Initialize a Counter[K] with a slice of K.
func NewCounter[T ~[]V, V cmp.Ordered](s T) *Counter[V] {
	m := make(map[V]int)
	for _, v := range s {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return &Counter[V]{
		Total: len(s),
		Set:   m,
	}
}

// Get the total number of element k in the Counter[K].
func (c *Counter[V]) Count(v V) int {
	i, ok := c.Set[v]
	if ok {
		return i
	}
	return 0
}

// Delete element k from the Counter[K].
func (c *Counter[V]) Delete(v V) {
	_, ok := c.Set[v]
	if ok {
		c.Total = c.Total - c.Set[v]
		delete(c.Set, v)
	}
}

func (c *Counter[V]) order() {
	if c.ordered {
		return
	}
	c.orderedList = make([]CounterItem[V], len(c.Set))
	i := 0
	for k, v := range c.Set {
		c.orderedList[i] = CounterItem[V]{k, v}
		i++
	}

	slices.SortFunc(c.orderedList, func(a, b CounterItem[V]) int {
		if n := cmp.Compare(a.Count, b.Count); n != 0 {
			return n
		}
		return cmp.Compare(a.Value, b.Value)
	})

	c.olLen = len(c.orderedList)
	c.ordered = true
}

func (c *Counter[V]) reverseOrder() {
	if c.reversed {
		return
	}
	c.reversedOrderedList = make([]CounterItem[V], len(c.Set))
	copy(c.reversedOrderedList, c.orderedList)
	slices.Reverse(c.reversedOrderedList)
	c.rolLen = len(c.reversedOrderedList)
	c.reversed = true
}

func (c *Counter[V]) MostCommon(n int) []CounterItem[V] {
	c.order()

	if n <= 0 || n > c.olLen {
		return c.orderedList
	}

	return c.orderedList[c.olLen-n:]
}

func (c *Counter[V]) LeastCommon(n int) []CounterItem[V] {
	c.order()

	if n <= 0 || n > c.olLen {
		return c.orderedList
	}

	return c.orderedList[:n]
}

func (c *Counter[V]) MostCommonValues(n int) []V {
	res := c.MostCommon(n)
	var values []V
	for _, v := range res {
		values = append(values, v.Value)
	}
	return values
}

func (c *Counter[V]) LeastCommonValues(n int) []V {
	res := c.LeastCommon(n)
	var values []V
	for _, v := range res {
		values = append(values, v.Value)
	}
	return values
}

func (c *Counter[V]) ASC() []CounterItem[V] {
	c.order()
	return c.orderedList
}

func (c *Counter[V]) ASCValues() []V {
	res := c.ASC()
	var values []V
	for _, v := range res {
		values = append(values, v.Value)
	}
	return values
}

func (c *Counter[V]) DESC() []CounterItem[V] {
	c.reverseOrder()
	return c.reversedOrderedList
}

func (c *Counter[V]) DESCValues() []V {
	res := c.DESC()
	var values []V
	for _, v := range res {
		values = append(values, v.Value)
	}
	return values
}

//todo optimize []V
