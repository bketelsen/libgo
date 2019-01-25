// stderr implements the libgo/log.Log interface that writes
// to all log information to os.Stderr.
package stderr

import (
	"libgo.io/log"
	internallog "libgo.io/log/internal/log"
	"os"
)

// New returns a new log.Log implementation that outputs log
// information to os.Stderr
func New() log.Log {
	return &internallog.Base{
		Writer: os.Stderr,
	}
}
