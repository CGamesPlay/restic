package backup

import (
	"testing"

	"github.com/restic/restic/lib/errors"
	"github.com/restic/restic/lib/test"
	"github.com/restic/restic/lib/ui"
)

func createJSONProgress() (*ui.MockTerminal, ProgressPrinter) {
	term := &ui.MockTerminal{}
	printer := NewJSONProgress(term, 3)
	return term, printer
}

func TestJSONError(t *testing.T) {
	term, printer := createJSONProgress()
	test.Equals(t, printer.Error("/path", errors.New("error \"message\"")), nil)
	test.Equals(t, []string{"{\"message_type\":\"error\",\"error\":{\"message\":\"error \\\"message\\\"\"},\"during\":\"archival\",\"item\":\"/path\"}\n"}, term.Errors)
}

func TestJSONScannerError(t *testing.T) {
	term, printer := createJSONProgress()
	test.Equals(t, printer.ScannerError("/path", errors.New("error \"message\"")), nil)
	test.Equals(t, []string{"{\"message_type\":\"error\",\"error\":{\"message\":\"error \\\"message\\\"\"},\"during\":\"scan\",\"item\":\"/path\"}\n"}, term.Errors)
}
