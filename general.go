package goslices

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

// copy a slice -> return a new slice
func Copy[T ~[]V, V any](s *T) T {
	c := make(T, len(*s))
	copy(c, *s)
	return c
}

// copy a slice -> return a pointer to it
func CopyRef[T ~[]V, V any](s *T) *T {
	c := make(T, len(*s))
	copy(c, *s)
	return &c
}

// get the last element of a slice
func Last[T ~[]V, V any](s *T) V {
	return (*s)[len(*s)-1]
}

// check if a slice is empty
func IsEmpty[T ~[]V, V any](s *T) bool {
	return len(*s) == 0
}

// reverse a slice -> return a new slice
func Reverse[T ~[]V, V any](s *T) T {
	l := len(*s)
	r := make(T, l)
	if l <= 1 {
		return r
	}
	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		r[i] = (*s)[j]
	}
	return r
}

// reverse a slice -> return a pointer to it
func ReverseRef[T ~[]V, V any](s *T) *T {
	l := len(*s)
	r := make(T, l)
	if l <= 1 {
		return &r
	}
	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		r[i] = (*s)[j]
	}
	return &r
}

// reverse a slice -> modify the original slice
func ReverseSelf[T ~[]V, V any](s *T) {
	l := len(*s)
	if l <= 1 {
		return
	}
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// remove the last element of a slice -> return a new slice with the removed element
func Pop[T ~[]V, V any](s *T) (T, V) {
	l := len(*s)
	if l == 0 {
		return *s, make(T, 1)[0]
	}
	return (*s)[:l-1], (*s)[l-1]
}

// remove the last element of a slice -> return a pointer to it with the removed element
func PopRef[T ~[]V, V any](s *T) (*T, V) {
	l := len(*s)
	if l == 0 {
		return s, make(T, 1)[0]
	}
	r := (*s)[:l-1]
	return &r, (*s)[l-1]
}

// remove the last element of a slice -> modify the original slice
func PopSelf[T ~[]V, V any](s *T) V {
	l := len(*s)
	if l == 0 {
		return make(T, 1)[0]
	}
	r := (*s)[l-1]
	*s = (*s)[:l-1]
	return r
}

// remove the first element of a slice -> return a new slice with the removed element
func Shift[T ~[]V, V any](s *T) (T, V) {
	l := len(*s)
	if l == 0 {
		return *s, make(T, 1)[0]
	}
	return (*s)[1:], (*s)[0]
}

// remove the first element of a slice -> return a pointer to it with the removed element
func ShiftRef[T ~[]V, V any](s *T) (*T, V) {
	l := len(*s)
	if l == 0 {
		return s, make(T, 1)[0]
	}
	r := (*s)[1:]
	return &r, (*s)[0]
}

// remove the first element of a slice -> modify the original slice
func ShiftSelf[T ~[]V, V any](s *T) V {
	l := len(*s)
	if l == 0 {
		return make(T, 1)[0]
	}
	r := (*s)[0]
	*s = (*s)[1:]
	return r
}

// add an element to the beginning of a slice -> return a new slice
func UnShift[T ~[]V, V any](s *T, v V) T {
	l := len(*s)
	r := make(T, l+1)
	r[0] = v
	copy(r[1:], *s)
	return r
}

// add an element to the beginning of a slice -> return a pointer to it
func UnShiftRef[T ~[]V, V any](s *T, v V) *T {
	l := len(*s)
	r := make(T, l+1)
	r[0] = v
	copy(r[1:], *s)
	return &r
}

// add an element to the beginning of a slice -> modify the original slice
func UnShiftSelf[T ~[]V, V any](s *T, v V) {
	*s = append(*s, make(T, 1)[0])
	copy((*s)[1:], *s)
	(*s)[0] = v

}

// map a slice -> return a new map of map[index]element
func Entries[T ~[]V, V any](s *T) map[int]V {
	r := make(map[int]V)
	for i, v := range *s {
		r[i] = v
	}
	return r
}

// sum the elements of a slice -> return the sum
func Sum[T ~[]N, N number](s *T) N {
	var r N
	for _, v := range *s {
		r += v
	}
	return r
}

// creates a new slice populated with the results of calling a provided function on every element in the slice.
func Map[T ~[]V, V any](s *T, f func(v V) V) T {
	r := make(T, len(*s))
	for i, v := range *s {
		r[i] = f(v)
	}
	return r
}

// return the maximum value of a slice -> return the zero value of that type if the slice is empty
func Max[T ~[]O, O ordered](s *T) O {
	if len(*s) == 0 {
		return make(T, 1)[0]
	}

	r := (*s)[0]
	for _, v := range *s {
		if v > r {
			r = v
		}
	}
	return r
}

// return the minimum value of a slice -> return the zero value of that type if the slice is empty
func Min[T ~[]O, O ordered](s *T) O {
	if len(*s) == 0 {
		return make(T, 1)[0]
	}

	r := (*s)[0]
	for _, v := range *s {
		if v < r {
			r = v
		}
	}
	return r
}

// creates a new slice of a portion of a given slice,
// filtered down to just the elements from the given slice that pass the test implemented by the provided function.
func Filter[T ~[]V, V any](s *T, f func(v V) bool) T {
	r := make(T, 0)
	for _, v := range *s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// tests whether all elements in the slice pass the test implemented by the provided function.
func Every[T ~[]V, V any](s *T, f func(v V) bool) bool {
	for _, v := range *s {
		if !f(v) {
			return false
		}
	}
	return true
}

// tests whether at least one element in the slice passes the test implemented by the provided function.
func Some[T ~[]V, V any](s *T, f func(v V) bool) bool {
	for _, v := range *s {
		if f(v) {
			return true
		}
	}
	return false
}

// check if a slice contains an element
func Includes[T ~[]V, V comparable](s *T, v V) bool {
	for _, i := range *s {
		if i == v {
			return true
		}
	}
	return false
}

// check if a slice contains an element that satisfies a given function
func IncludesFunc[T ~[]V, V any](s *T, f func(v V) bool) bool {
	for _, i := range *s {
		if f(i) {
			return true
		}
	}
	return false
}

// get the first index of an element in a slice -> return -1 if not found
func Index[T ~[]V, V comparable](s *T, v V) int {
	for i, j := range *s {
		if j == v {
			return i
		}
	}
	return -1
}

//
// func IndexOf[T ~[]V, V comparable](s *T, v V) int {
// 	return Index(s, v)
// }
//

// get the last index of an element in a slice -> return -1 if not found
func LastIndex[T ~[]V, V comparable](s *T, v V) int {
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i] == v {
			return i
		}
	}
	return -1
}

// find the first element in a slice -> return the element if found, else return the zero value of that type
func Find[T ~[]V, V comparable](s *T, v V) V {
	if i := Index(s, v); i != -1 {
		return (*s)[i]
	} else {
		return make(T, 1)[0]
	}
}

// find the last element in a slice -> return the element if found, else return the zero value of that type
func FindLast[T ~[]V, V comparable](s *T, v V) V {
	if i := LastIndex(s, v); i != -1 {
		return (*s)[i]
	} else {
		return make(T, 1)[0]
	}
}

// find the first element in a slice that satisfies a given function -> return the element if found, else return the zero value of that type
func FindFirstFunc[T ~[]V, V any](s *T, f func(v V) bool) V {
	for _, v := range *s {
		if f(v) {
			return v
		}
	}
	return make(T, 1)[0]
}

// find the last element in a slice that satisfies a given function -> return the element if found, else return the zero value of that type
func FindLastFunc[T ~[]V, V any](s *T, f func(v V) bool) V {
	for i := len(*s) - 1; i >= 0; i-- {
		if f((*s)[i]) {
			return (*s)[i]
		}
	}
	return make(T, 1)[0]
}

// test if two slices are equal
func IsEqual[T ~[]V, V comparable](s *T, r *T) bool {
	if len(*s) != len(*r) {
		return false
	}
	for i, v := range *s {
		if v != (*r)[i] {
			return false
		}
	}
	return true
}

// test if two slices are equal using a given function
func IsEqualFunc[T ~[]V, V any](s *T, r *T, f func(v V, r V) bool) bool {
	if len(*s) != len(*r) {
		return false
	}
	for i, v := range *s {
		if !f(v, (*r)[i]) {
			return false
		}
	}
	return true
}

func IsUnique[T ~[]V, V comparable](s *T) bool {
	unique := make(map[V]bool)
	for _, v := range *s {
		if unique[v] {
			return false
		}
		unique[v] = true
	}
	return true
}

func JoinInt(array []int, sep string) string {
	var str strings.Builder
	for i, v := range array {
		if i != 0 {
			str.WriteString(sep)
		}
		str.WriteString(strconv.Itoa(v))
	}
	str.WriteString("")
	return str.String()
}

func Sort[T ~[]V, V cmp.Ordered](s T) T {
	m := make(T, len(s))
	copy(m, s)
	slices.Sort(m)
	return m
}
