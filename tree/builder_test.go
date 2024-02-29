package tree

import (
	"strings"
	"testing"
)

func TestPrefixUniqueness(t *testing.T) {
	symbols := New("beep boop beer!").Leaves()

	for _, a := range symbols {
		for _, b := range symbols {
			if a != b && strings.HasPrefix(a.String(), b.String()) {
				t.Fatalf("prefix uniqueness is broken")
			}
		}
	}
}
