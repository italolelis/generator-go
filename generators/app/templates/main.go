package main

import (
	"fmt"
	"os"

	"<%=projectRoot%>/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
