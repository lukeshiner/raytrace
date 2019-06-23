package comparison

import (
	"testing"
)

func TestEpsilonEqual(t *testing.T) {
	a := 0.20000000000000007
	b := 0.2
	if EpsilonEqual(a, b) != true {
		t.Error("Equality check failed.")
	}
}
