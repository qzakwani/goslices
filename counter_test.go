package goslices

import "testing"

func TestCounter(t *testing.T) {
	// slice of string some of which are duplicates total of 20
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "a", "b", "c", "d", "e", "f", "g", "h", "i", "a", "b", "z"}

	// create a new Counter[string] from the slice
	c := NewCounter(s)
	// check the total is 21
	if c.Total != 21 {
		t.Errorf("Expected 20 got %v", c.Total)
	}
	// check the count of "a" is 3
	if c.Count("a") != 3 {
		t.Errorf("Expected 3 got %v", c.Count("a"))
	}
	// check the count of "b" is 3
	if c.Count("b") != 3 {
		t.Errorf("Expected 3 got %v", c.Count("b"))
	}
	// check most common is "a" or "b"
	if c.MostCommon(1)[0].Value != "a" && c.MostCommon(1)[0].Value != "b" {
		t.Errorf("Expected a or b got %v", c.MostCommon(1)[0].Value)
	}
	// check least common is "z"
	if c.LeastCommon(1)[0].Value != "z" {
		t.Errorf("Expected z got %v", c.LeastCommon(1)[0].Value)
	}
	// check the count of "z" is 1
	if c.Count("z") != 1 {
		t.Errorf("Expected 1 got %v", c.Count("z"))
	}
	// check least common values is "z"
	if c.LeastCommonValues(1)[0] != "z" {
		t.Errorf("Expected z got %v", c.LeastCommonValues(1)[0])
	}
	// check most common values is "a" or "b"
	if c.MostCommonValues(1)[0] != "a" && c.MostCommonValues(1)[0] != "b" {
		t.Errorf("Expected a or b got %v", c.MostCommonValues(1)[0])
	}
	// get ascending order
	if c.ASCValues()[0] != "z" && (c.ASCValues()[len(c.ASCValues())-1] != "a" || c.ASCValues()[len(c.ASCValues())-1] != "b") {
		t.Errorf("Expected asc order got %v", c.ASCValues()[0])
	}
	// get descending order
	if c.DESCValues()[len(c.DESCValues())-1] != "z" && (c.DESCValues()[0] != "a" || c.DESCValues()[0] != "b") {
		t.Errorf("Expected desc order got %v", c.DESCValues()[0])
	}
}
