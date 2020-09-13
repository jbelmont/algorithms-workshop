package quadratic

import (
	"testing"
)

func TestQuadraticComplexity(t *testing.T) {
	quadratic := quadraticComplexity(10)
	if quadratic != 100 {
		t.Error("10 * 10 should equal 100")
	}
}
