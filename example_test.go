package compare_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/readwritepro/compare-test-results"
)

func ExampleExpectedActual() {
	fIn, _ := os.Open("testdata/fixture.txt")
	defer fIn.Close()

	fOut, _ := os.Create("testdata/actual.txt")
	defer fOut.Close()

	scanner := bufio.NewScanner(fIn)
	scanner.Split(bufio.ScanLines)
	writer := bufio.NewWriter(fOut)

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		fmt.Fprintf(writer, "%04d %s\n", lineNum, line)
	}
	writer.Flush()

	mockT := testing.T{}
	var t *testing.T = &mockT
	compare.ExpectedActual(t, "testdata/expected.txt", "testdata/actual.txt")
	// Output:
	// 0001 THE QUEEN'S TWIN AND OTHER STORIES
	// 0002 BY SARAH ORNE JEWETT
	// 0003 The coast of Maine was in former years brought so near to foreign
	// 0004 shores by its busy fleet of ships that among the older men and women
	// 0005 one still finds a surprising proportion of travelers.  Each
	// 0006 seaward-stretching headland with its high-set houses, each island of a
	// 0007 single farm, has sent its spies to view many a Land of Eshcol; one may
	// 0008 see plain, contented old faces at the windows, whose eyes have looked
	// 0009 at far-away ports and known the splendors of the Eastern world.  They
	// 0010 shame the easy voyager of the North Atlantic and the Mediterranean;
	// 0011 they have rounded the Cape of Good Hope and braved the angry seas of
	// 0012 Cape Horn in small wooden ships; they have brought up their hardy boys
	// 0013 and girls on narrow decks; they were among the last of the Northmen's
	// 0014 children to go adventuring to unknown shores.  More than this one
	// 0015 cannot give to a young State for its enlightenment; the sea captains
	// 0016 and the captains' wives of Maine knew something of the wide world, and
	// 0017 never mistook their native parishes for the whole instead of a part
	// 0018 thereof; they knew not only Thomaston and Castine and Portland, but
	// 0019 London and Bristol and Bordeaux, and the strange-mannered harbors of
	// 0020 the China Sea.
}
