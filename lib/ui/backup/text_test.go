package backup

import (
	"testing"

	"github.com/restic/restic/lib/errors"
	"github.com/restic/restic/lib/test"
	"github.com/restic/restic/lib/ui"
)

func createTextProgress() (*ui.MockTerminal, ProgressPrinter) {
	term := &ui.MockTerminal{}
	printer := NewTextProgress(term, 3)
	return term, printer
}

func TestError(t *testing.T) {
	term, printer := createTextProgress()
	test.Equals(t, printer.Error("/path", errors.New("error \"message\"")), nil)
	test.Equals(t, []string{"error: error \"message\"\n"}, term.Errors)
}

func TestScannerError(t *testing.T) {
	term, printer := createTextProgress()
	test.Equals(t, printer.ScannerError("/path", errors.New("error \"message\"")), nil)
	test.Equals(t, []string{"scan: error \"message\"\n"}, term.Errors)
}
