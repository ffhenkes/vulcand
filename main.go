package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ffhenkes/vulcand/plugin/registry"
	"github.com/ffhenkes/vulcand/service"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := service.Run(registry.GetRegistry()); err != nil {
		fmt.Printf("Service exited with error: %s\n", err)
		os.Exit(service.ExitCodeError)
	} else {
		fmt.Println("Service exited gracefully")
	}
}
