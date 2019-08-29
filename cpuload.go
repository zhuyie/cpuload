package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 1, "thread count")
	flag.Parse()

	fmt.Printf("Thread count: %v\n", count)
	fmt.Printf("Press Ctrl+C to exit...")

	for i := 0; i < count; i++ {
		go cpuload()
	}

	qc := make(chan os.Signal, 1)
	signal.Notify(qc, syscall.SIGINT, syscall.SIGTERM)
	<-qc
}

var dummy int

func cpuload() {
	for {
		dummy = dummy + 1
	}
}
