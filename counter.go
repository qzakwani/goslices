package goslices

type Counter[V comparable] struct {
	Total   int
	Set     map[V]int
	// Order   []K
	Ordered bool
}

// Initialize a Counter[K] with a slice of K.
func NewCounter[T ~[]V, V comparable](s T) *Counter[V] {
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

// todo: implement
// func (c *Counter[K]) MostCommon(n int) []K {
// 	if n <= 0 {
// 		return []K{}
// 	}
// 	m := jmaps.Values(c.Set)
// 	l := len(m)
// 	if l <= n {
// 		n = l
// 	}
// 	c.Order = make([]K, n)
// 	slices.Sort(m)
// 	j := 0
// 	for i := l-1; i < n; i-- {
// 		c.Order[j] = m[i]
// 		j++
// 	}

// }
