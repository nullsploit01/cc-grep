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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccgrep [pattern] [file]",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			os.Exit(1)
		}

		pattern := args[0]
		fileName := args[1]

		g := internal.NewGrep()

		if recursive {
			g.RecursiveGrep()
			os.Exit(0)
		}

		matches, err := g.Grep(pattern, fileName, caseInsensetive)
		if err != nil {
			cmd.ErrOrStderr().Write([]byte(err.Error() + "\n"))
			os.Exit(1)
		}

		output := strings.Join(matches, "\n")
		if len(matches) > 0 {
			output += "\n"
		}

		cmd.OutOrStdout().Write([]byte(output))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cc-grep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&caseInsensetive, "case-insensetive", "i", false, "Case insensetive")
}
