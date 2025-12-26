package main

import (
	"github.com/spf13/cobra"
)

var aliases = []string{"exe", "ex", "e"}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "govel_cli [flags] [command]",
	Short:        "govel_cli is a command line interface for govel",
	Long:         `govel_cli. Command line interface for govel, create scalfold for a project govel`,
	SilenceUsage: false,
	Args:         fValidateArgsOfRoot,
	Run:          fRunOfRoot,
}

var executeCmd = &cobra.Command{
	Use:          "execute",
	Short:        "execute: create or view info of project govel",
	Long:         `execute: create or view info of project govel, depending on the flags create or info. If no flag is specified, the default is create project govel`,
	SilenceUsage: false,
	Aliases:      aliases,
	Args:         validCantArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},
}

var cmds = []*cobra.Command{executeCmd}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	initFlags()
	rootCmd.AddCommand(cmds...)
}
