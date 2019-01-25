// package log defines an interface for Go programs that log.
//
// *Philosophy*
//
// - structured logging underlies the operation of this logger.
// - levels, prefixes, context, and even the log message itself
//   can be modeled as a chain (linked list) of key value pairs.
// - logging performance or allocations are not a design point.
//   If you need allocation free logging then you are either logging
//   too much, or are not in a position to use a general purpose
//   logger.
//
// *Usage*
//
// - this logger package contains exported types, only interfaces.
// - to construct a logger, known as the base logger because other
//   loggers extend from it, use something like log/stderr.New.
package log

// Log defines an interface for a logger which works with
// structured, leveled, or a combination of both, styles.
type Log interface {
	Println(vals ...interface{})
	WithValue(key, val interface{}) Log
}
