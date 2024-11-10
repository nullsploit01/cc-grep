/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/nullsploit01/cc-grep/internal"
	"github.com/spf13/cobra"
)

var caseInsensetive bool
var recursive bool
var invert bool

var rootCmd = &cobra.Command{
	Use:   "ccgrep [pattern] [file]",
	Short: "Search for a pattern in a file or directory, with options for case insensitivity, recursion, and inverted matches.",
	Long: `ccgrep is a tool to search for a specified pattern in files, with support for advanced options.

	Flags:
	-i, --case-insensitive  Perform a case-insensitive search.
	-r, --recursive         Search recursively within directories.
	-v, --invert            Return lines that do NOT match the pattern.

	Examples:
	# Basic usage: search for "example" in file.txt
	ccgrep example file.txt

	# Case-insensitive search
	ccgrep -i example file.txt

	# Recursive search for "example" in all files under a directory
	ccgrep -r example /path/to/directory

	# Invert match: return lines that do NOT contain "example"
	ccgrep -v example file.txt
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			os.Exit(1)
		}

		pattern := args[0]
		fileName := args[1]
		g := internal.NewGrep()

		handleError := func(err error) {
			if err != nil {
				cmd.ErrOrStderr().Write([]byte(err.Error() + "\n"))
				os.Exit(1)
			}
		}

		printMatches := func(fileName string, matches []string) {
			for _, match := range matches {
				cmd.OutOrStdout().Write([]byte(fileName + ": " + match + "\n"))
			}
		}

		if recursive {
			recursiveMatches, err := g.RecursiveGrep(fileName, pattern, caseInsensetive, invert)
			handleError(err)

			for _, result := range recursiveMatches {
				printMatches(result.FileName, result.Matches)
			}
		} else {
			matches, err := g.Grep(pattern, fileName, caseInsensetive, invert)
			handleError(err)

			if len(matches) > 0 {
				cmd.OutOrStdout().Write([]byte(strings.Join(matches, "\n") + "\n"))
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&caseInsensetive, "case-insensetive", "i", false, "Case insensetive")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursive")
	rootCmd.Flags().BoolVarP(&invert, "invert", "v", false, "Invert")
}
