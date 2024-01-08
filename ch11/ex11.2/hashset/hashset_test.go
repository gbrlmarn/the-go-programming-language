package hashset

import "testing"

func InitHasSet() HashSet {
    return make(HashSet)
}

func TestHas(t *testing.T) {
    s := InitHasSet()
	s.Add(1)
	if !s.Has(1) {
		t.Errorf("!s.Has(1)")
	}
}

func TestAdd(t *testing.T) {
    s := InitHasSet()
	s.Add(1)
	if !s.Has(1) {
		t.Errorf("!s.Has(1)")
	}
	s.Add(2)
	if !s.Has(2) {
		t.Errorf("!s.Has(2)")
	}
}

func TestUnionWith(t *testing.T) {
    s1 := InitHasSet()
	s1.Add(1)
	s1.Add(2)
    s2 := InitHasSet()
	s2.Add(3)
	s2.Add(4)
	s1.UnionWith(s2)
	if !s1.Has(3) {
		t.Errorf("!s1.Has(3)")
	}
	if !s1.Has(4) {
		t.Errorf("!s1.Has(4)")
	}
}

func TestString(t *testing.T) {
    s := InitHasSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
    if s.String() != "{1 2 3}" {
        t.Errorf("got: %v\nwant: %v", s.String(), "{1 2 3}") 
    }
}
