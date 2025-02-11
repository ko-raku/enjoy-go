package imgprocessor

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Extractor インターフェイス定義
// OCR処理を行う関数を抽象化
type Extractor interface {
	ExtractText(inputPath string) (string, error)
}

// TesseractExtractor Tesseractを利用したOCRの実装
type TesseractExtractor struct{}

// ExtractText OCR実行 (Tesseractを利用)
func (t *TesseractExtractor) ExtractText(inputPath string) (string, error) {
	// OCR 出力用バッファ
	var outBuffer bytes.Buffer

	// OCR 実行コマンド
	cmd := exec.Command(
		"tesseract",
		inputPath,
		"stdout",        // OCR 結果を標準出力に
		"-l", "jpn+eng", // 言語モデル
		"--tessdata-dir", "/opt/homebrew/share/tessdata/",
		"--oem", "1", // LSTM OCR モード
		"--psm", "3", // 自由テキストレイアウトの設定
		"--user-words", "words.txt", // カスタム辞書の付加も効果的（任意）
	)
	cmd.Stdout = &outBuffer

	// 実行とエラーチェック
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("OCR 実行エラー: %w", err)
	}

	return outBuffer.String(), nil
}
