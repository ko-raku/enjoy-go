package cli

import (
	img "enjoy/image"
	"github.com/spf13/cobra"
	"path/filepath"
)

var inputPath string
var outputPath string

// Processor 依存性を注入
var Processor img.Processor = &img.DefaultProcessor{}

var GrayscaleCmd = &cobra.Command{
	Use:   "grayscale",
	Short: "Convert an image to grayscale",
	Long:  "A CLI tool for various image processing tasks such as grayscale conversion.",
	Run: func(cmd *cobra.Command, args []string) {
		if inputPath == "" {
			cmd.Println("Error: --input flag is required.")
			return
		}
		if outputPath == "" {
			base := filepath.Base(inputPath)
			ext := filepath.Ext(inputPath)
			name := base[:len(base)-len(ext)]
			outputPath = filepath.Join(filepath.Dir(inputPath), name+"_gray"+ext)
			cmd.Printf("Output path not specified. Using %s\n", outputPath)
		}
		// Processor を利用して変換処理を実行
		if Processor == nil {
			cmd.Println("Error: No processor is set.")
			return
		}
		err := Processor.ConvertToGray(inputPath, outputPath)
		if err != nil {
			cmd.Printf("Error: %v\n", err)
			return
		}
		cmd.Printf("Grayscale image created successfully: %s\n", outputPath)
	},
}

func init() {
	GrayscaleCmd.Flags().StringVar(&inputPath, "input", "", "Path to the input image file (required)")
	GrayscaleCmd.Flags().StringVar(&outputPath, "output", "", "Path to save the output image file (optional)")
}
