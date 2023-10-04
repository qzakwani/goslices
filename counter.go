package goslices

type Counter[K comparable] struct {
	Total   int
	Set     map[K]int
	Order   []K
	Ordered bool
}

// Initialize a Counter[K] with a slice of K.
func NewCounter[K comparable](s *[]K) *Counter[K] {
	m := make(map[K]int)
	for _, v := range *s {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return &Counter[K]{
		Total: len(*s),
		Set:   m,
	}
}

// Get the total number of element k in the Counter[K].
func (c *Counter[K]) Count(k K) int {
	v, ok := c.Set[k]
	if ok {
		return v
	}
	return 0
}

// Delete element k from the Counter[K].
func (c *Counter[K]) Delete(k K) {
	_, ok := c.Set[k]
	if ok {
		c.Total = c.Total - c.Set[k]
		delete(c.Set, k)
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
