package main

import (
	"fmt"
	"os"

	xxhtmlcli "github.com/zulubit/xxhtml/cmd/xxhtml/cli"
)

func main() {
	if err := xxhtmlcli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

}
