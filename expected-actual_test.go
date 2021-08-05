package compare_test

import (
	"testing"

	"github.com/joehonton/compare-files"
)

func TestCompare(t *testing.T) {
	compare.ExpectedActual(t, "testdata/expected.txt", "testdata/actual.txt")
}
