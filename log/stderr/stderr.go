// stderr implements the libgo/log.Log interface that writes
// to all log information to os.Stderr.
package stderr

import (
	"os"

	"github.com/bketelsen/libgo/log"
	internallog "github.com/bketelsen/libgo/log/internal/log"
)

// New returns a new log.Log implementation that outputs log
// information to os.Stderr
func New() log.Log {
	return &internallog.Base{
		Writer: os.Stderr,
	}
}
