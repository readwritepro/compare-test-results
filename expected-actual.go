//=============================================================================
// File:     expected-actual.go
//=============================================================================

// Package compare-test-results uses Google's classic diff-match-patch algorithm
// to compare two files. It sends the color highlighted output to *testing.T.
// Use this package to automatically display file-based diffs from within "go test".
package compare

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// The ExpectedActual function compares two files formatting them with red and green
// highlights to show subtractions and additions. The output is sent to *testing.T
// which will cause the calling test to FAIL when there is a difference and to PASS
// when they are the same.
func ExpectedActual(t *testing.T, expectedFile string, actualFile string) {
	expectedBytes, _ := ioutil.ReadFile(expectedFile)
	actualBytes, _ := ioutil.ReadFile(actualFile)

	if !bytes.Equal(expectedBytes, actualBytes) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(expectedBytes), string(actualBytes), false)
		diffs = dmp.DiffCleanupSemantic(diffs)
		t.Errorf("difference between expected file '%s' and actual file '%s'", expectedFile, actualFile)
		t.Errorf(dmp.DiffPrettyText(diffs))
	}
}
