package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gosuri/uilive"
)

func worker() {
	for {
		time.Sleep(1 * time.Millisecond)
	}
}

func systemThreadCount() int {
	return 0
}

func showStats() {
	term := uilive.New()
	term.RefreshInterval = 1 * time.Second
	term.Start()
	n := 0
	for {
		fmt.Fprintln(term, n, "s")

		g := runtime.NumGoroutine()
		fmt.Fprintln(term.Newline(), g, "goroutines")

		t := systemThreadCount()
		fmt.Fprintln(term.Newline(), t, "threads")

		c := runtime.NumCPU()
		fmt.Fprintln(term.Newline(), c, "CPUs")

		term.Flush()
		n++

		time.Sleep(1 * time.Second)
	}
}

func main() {
	go showStats()

	for i := 0; i < 100000; i++ {
		go worker()
	}

	time.Sleep(10 * time.Second)
}
