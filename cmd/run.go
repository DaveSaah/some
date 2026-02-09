package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DaveSaah/some/lexer"
	"github.com/DaveSaah/some/token"
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

		run(filepath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
