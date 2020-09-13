package cubic

import (
	"testing"
)

func TestCubicComplexity(t *testing.T) {
	cubic := cubicComplexity()
	expected := 162
	if cubic != expected {
		t.Errorf("The computed value is equal %d, but received %d\n", expected, cubic)
	}
}
