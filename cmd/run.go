/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/DaveSaah/some/compile"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run some source code",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		if ext := strings.Split(filepath, ".")[1]; ext != "sm" {
			return fmt.Errorf("%s is not some file", filepath)
		}

		compile.Start(filepath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
