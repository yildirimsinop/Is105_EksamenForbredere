package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Panicln(err)
	}
	p.Signal(syscall.SIGSTOP)
}
