package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	startupTime := time.Now()

	for {

		fmt.Fprintf(os.Stdout, "Hello World! Alive since %ds\n", int(time.Since(startupTime).Seconds()))
		time.Sleep(time.Second * 30)
	}
}
