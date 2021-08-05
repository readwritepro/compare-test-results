//=============================================================================
// File:     expected-actual.go
//=============================================================================

// Package compare-files uses Google's classic diff-match-patch algorithm to compare two files.
// It sends the color highlighted output to *testing.T for use when testing expected
// versus actual results.
package compare

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Test utility function
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
