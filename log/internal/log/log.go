package log

import (
	"fmt"
	"io"
	"runtime"
	"time"

	"libgo.io/log"
)

type Base struct {
	io.Writer
}

func (b *Base) WithValue(key, val interface{}) log.Log {
	return b.withValue(key, val)
}

func (b *Base) withValue(key, val interface{}) *Context {
	return &Context{
		logger: b,
		key:    key,
		val:    val,
	}
}

func (b *Base) output(c *Context) {

}

func (b *Base) Println(vals ...interface{}) {
	b.withValue(message{}, fmt.Sprintln(vals...)).log(1)
}

type logger interface {
	output(*Context)
	withValue(key, val interface{}) *Context
}

type Context struct {
	logger
	key, val interface{}
}

type (
	level     struct{}
	message   struct{}
	timestamp struct{}
	pc        struct{}
)

func (c *Context) WithValue(key, val interface{}) log.Log {
	return c.withValue(key, val)
}

func (c *Context) withValue(key, val interface{}) *Context {
	return &Context{
		logger: c,
		key:    key,
		val:    val,
	}
}

func (c *Context) Println(vals ...interface{}) {
	c.withValue(message{}, fmt.Sprintln(vals...)).log(1)
}

func (c *Context) log(skip int) {
	var callstack [1]uintptr
	runtime.Callers(skip+1, callstack[:])
	l := c.withValue(pc{}, callstack[0]).withValue(timestamp{}, time.Now())
	l.output(l)
}

// func (l *log) Format(f State, c rune)
