package stderr_test

import (
	"libgo.io/log"
	"libgo.io/log/stderr"
)

func ExamplePrintln() {
	log := stderr.New()

	log.Println("a simple log message without prefix")
}

type level struct{}

func ExampleWithValue() {
	log := stderr.New()

	info := log.WithValue(level{}, "INFO")
	info.Println("everything's fine")

	err := log.WithValue(level{}, "ERROR")
	err.Println("everything's not fine")
}

type prefix struct{}

func ExampleWithPrefix() {
	fn := func(log log.Log) {
		// do some work
		log.WithValue(level{}, "INFO").Println("everythings cool")
	}

	log := stderr.New()
	fn(log.WithValue(prefix{}, "important function"))
}
