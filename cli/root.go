package cli

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "imagecli",
	Short: "A CLI tool for image processing",
	Long:  "A CLI tool for various image processing tasks like grayscale conversion, resizing, etc.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(grayscaleCmd)
}
