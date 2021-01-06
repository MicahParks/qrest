package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mvo5/qrest-skeleton/backend"
	"github.com/mvo5/qrest-skeleton/mgrd"
)

func run() error {
	backend := backend.NewQuotaManager()
	mgrd := mgrd.New(":8080", backend)
	// support graceful shutdown
	go func() {
		cancelCh := make(chan os.Signal, 1)
		signal.Notify(cancelCh, syscall.SIGTERM, syscall.SIGINT)
		<-cancelCh
		fmt.Println("shuting down gracefully")
		mgrd.Stop()
	}()
	// and run
	if err := mgrd.Start(); err != nil {
		return err
	}
	return mgrd.Wait()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
