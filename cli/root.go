package cli

import (
	"github.com/spf13/cobra"
	"os"
)

// exitFunc を定義。デフォルトは os.Exit に設定。
var exitFunc = os.Exit

var RootCmd = &cobra.Command{
	Use:   "image",
	Short: "A CLI tool for image processing",
	Long:  "A CLI tool for various image processing tasks like grayscale conversion, resizing, etc.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		exitFunc(1)
	}
}

func init() {
	RegisterCommands()
}

func SetExitFunc(fn func(code int)) {
	if fn == nil {
		exitFunc = os.Exit
	} else {
		exitFunc = fn
	}
}
