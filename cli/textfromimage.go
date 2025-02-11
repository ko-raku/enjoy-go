package cli

import (
	"fmt"
	"os"
	"path/filepath"

	img "enjoy/imgprocessor" // imgprocessor パッケージをインポート
	"github.com/spf13/cobra"
)

var ocrInputPath string
var ocrOutputPath string

// OCRProcessor と Extractor の依存性注入
var OCRProcessor img.Processor = &img.DefaultProcessor{}
var Extractor img.Extractor = &img.TesseractExtractor{}

var TextFromImageCmd = &cobra.Command{
	Use:   "textfromimage",
	Short: "Extract text from an image using OCR",
	Long:  "A CLI tool that preprocesses an image and extracts text using the Tesseract OCR engine.",
	Run: func(cmd *cobra.Command, args []string) {
		// 入力ファイルの検証
		if ocrInputPath == "" {
			cmd.Println("Error: --input flag is required.")
			return
		}

		// 出力パスが指定されていない場合のデフォルト設定
		if ocrOutputPath == "" {
			base := filepath.Base(ocrInputPath)
			ext := filepath.Ext(base)
			name := base[0 : len(base)-len(ext)]
			ocrOutputPath = filepath.Join(filepath.Dir(ocrInputPath), name+"_output.txt")
			cmd.Printf("Output path not specified. Using %s\n", ocrOutputPath)
		}

		// 前処理画像の一時保存先
		tempImagePath := filepath.Join(filepath.Dir(ocrInputPath), "temp_processed.png")

		// 1. 画像の前処理
		if OCRProcessor == nil {
			cmd.Println("Error: No processor is configured.")
			return
		}
		err := OCRProcessor.OptimizeImageForOCR(ocrInputPath, tempImagePath)
		if err != nil {
			cmd.Printf("Error during image preprocessing: %v\n", err)
			return
		}

		// 2. OCR 実行
		if Extractor == nil {
			cmd.Println("Error: No OCR extractor is configured.")
			return
		}
		text, err := Extractor.ExtractText(tempImagePath)
		if err != nil {
			cmd.Printf("Error during OCR processing: %v\n", err)
			return
		}

		// 3. テキストの保存または表示
		err = saveTextToFile(cmd, text, ocrOutputPath)
		if err != nil {
			cmd.Printf("Error saving OCR output: %v\n", err)
			return
		}

		cmd.Printf("OCR text extraction successful. Output saved to: %s\n", ocrOutputPath)

		// 一時ファイル削除
		_ = os.Remove(tempImagePath)
	},
}

func init() {
	TextFromImageCmd.Flags().StringVar(&ocrInputPath, "input", "", "Path to the input image (required)")
	TextFromImageCmd.Flags().StringVar(&ocrOutputPath, "output", "", "Path to save the extracted text (optional)")
}

// saveTextToFile saves the OCR-extracted text to a file or displays it on the console.
func saveTextToFile(cmd *cobra.Command, text, filePath string) error {
	if filePath == "" {
		cmd.Println("Extracted text:")
		cmd.Println(text)
		return nil
	}

	// ファイルに書き込む
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("failed to write text to file: %w", err)
	}
	return nil
}
