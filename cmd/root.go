package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	_ "github.com/fatih/color"
	"github.com/iJanki/google/browser"
)

var rootCmd = &cobra.Command{
	Use:   "google [args]",
	Short: "Open Google results in default browser from command line",
	Long:  "Open Google results in default browser from command line",
	Args:  cobra.MinimumNArgs(1),
	Run:   open,
}

func Execute() {
	rootCmd.Execute()
}

func open(_ *cobra.Command, args []string) {

	query := strings.Join(args, " ")

	browser.Start(query)
}

