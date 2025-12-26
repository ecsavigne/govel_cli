package main

import (
	"fmt"
	"os"
)

var projectName string = "ecs_govel"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
