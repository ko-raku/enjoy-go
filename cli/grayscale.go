package cli

import (
    "enjoy/imageprocessor"
    "fmt"
    "github.com/spf13/cobra"
    "path/filepath"
)

var inputPath string
var outputPath string

var grayscaleCmd = &cobra.Command{
    Use:   "grayscale",
    Short: "Convert an image to grayscale",
    Long:  "A CLI tool for various image processing tasks such as grayscale conversion.",
    Run: func(cmd *cobra.Command, args []string) {
        if inputPath == "" {
            fmt.Println("Error: --input flag is required.")
            return
        }
        if outputPath == "" {
            base := filepath.Base(inputPath)
            ext := filepath.Ext(inputPath)
            name := base[:len(base)-len(ext)]
            outputPath = filepath.Join(filepath.Dir(inputPath), name+"_gray"+ext)
            fmt.Printf("Output path not specified. Using %s\n", outputPath)
        }
        err := imageprocessor.ConvertToGray(inputPath, outputPath)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Printf("Grayscale image created successfully: %s\n", outputPath)
    },
}

func init() {
    grayscaleCmd.Flags().StringVar(&inputPath, "input", "", "Path to the input image file (required)")
    grayscaleCmd.Flags().StringVar(&outputPath, "output", "", "Path to save the output image file (optional)")
}
