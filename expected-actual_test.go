package compare_test

import (
	"testing"

	"github.com/readwritepro/compare-test-results"
)

func TestCompare(t *testing.T) {
	compare.ExpectedActual(t, "testdata/expected.txt", "testdata/actual.txt")
}
