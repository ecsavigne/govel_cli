package main

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	if executeFlag != "" && (!slices.Contains(aliases, executeFlag) && executeFlag != "execute") {
		projectName = executeFlag
	}
	fmt.Printf("Project Name: %s\n", projectName)

	switch cmd.Name() {
	case "execute":
		createProject(projectName)
		fmt.Printf("Comando ejecutado\n")
	}
}
