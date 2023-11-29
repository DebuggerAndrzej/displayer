package main

import (
	"fmt"
	"github.com/DebuggerAndrzej/displayer/backend"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("Start using bestest OS :)")
	}
	fmt.Println(os.Getenv("XDG_SESSION_TYPE"))

	fmt.Println(backend.GetMonitorData())
}
