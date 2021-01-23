package restic_test

import (
	"testing"
	"time"

	"github.com/restic/restic/lib/restic"
	rtest "github.com/restic/restic/lib/test"
)

func TestNewSnapshot(t *testing.T) {
	paths := []string{"/home/foobar"}

	_, err := restic.NewSnapshot(paths, nil, "foo", time.Now())
	rtest.OK(t, err)
}
