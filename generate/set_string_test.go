package generate

import (
	"testing"
)

func TestUnion(t *testing.T) {
	a := make(StringSet)
	a.Add("a", "b", "c")
	b := make(StringSet)
	b.Add("b", "c", "d")
	unioned := a.Union(b)
	expected := make(StringSet)
	expected.Add("a", "b", "c", "d")
	if !unioned.Equals(expected) {
		t.Error("unioned set does not meet expectations")
	}
}
