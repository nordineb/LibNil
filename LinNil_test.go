package LibNil

import "testing"

func TestNil(t *testing.T) {
	want := (*int)(nil)
	if got := Nil(); got != want {
		t.Errorf("Nil() should return nil")
	}
}
