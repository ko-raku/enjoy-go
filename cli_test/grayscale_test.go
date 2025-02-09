package cli_test

import (
	"bytes"
	"enjoy/cli"
	"enjoy/mock"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func TestGrayscaleCmd(t *testing.T) {
	// モック準備
	mockProcessor := &mock.MockImageProcessor{}
	cli.Processor = mockProcessor // モックを注入

	// コマンド実行の準備
	cmd := &cobra.Command{}
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	// 引数とフラグを設定
	cli.GrayscaleCmd.Flags().Set("input", "/path/to/input.jpg")
	cli.GrayscaleCmd.Flags().Set("output", "/path/to/output.jpg")

	// コマンド実行
	cli.GrayscaleCmd.Run(cmd, []string{})

	// 標準出力を確認
	output := buf.String()

	// メッセージの検証
	if !strings.Contains(output, "Grayscale image created successfully") {
		t.Errorf("expected success message, but got '%s'", output)
	}

	// モックが適切に呼び出されているか検証
	if mockProcessor.CalledInputPath != "/path/to/input.jpg" {
		t.Errorf("expected inputPath '/path/to/input.jpg', but got '%s'", mockProcessor.CalledInputPath)
	}
	if mockProcessor.CalledOutputPath != "/path/to/output.jpg" {
		t.Errorf("expected outputPath '/path/to/output.jpg', but got '%s'", mockProcessor.CalledOutputPath)
	}
}

func TestGrayscaleCmd_MockFailure(t *testing.T) {
	// モック準備
	mockProcessor := &mock.MockImageProcessor{ShouldFail: true}
	cli.Processor = mockProcessor

	// コマンド実行の準備
	cmd := &cobra.Command{}
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	// 引数とフラグを設定
	cli.GrayscaleCmd.Flags().Set("input", "/path/to/input.jpg")
	cli.GrayscaleCmd.Flags().Set("output", "/path/to/output.jpg")

	// コマンド実行
	cli.GrayscaleCmd.Run(cmd, []string{})

	// 標準出力を確認
	output := buf.String()

	// エラーメッセージの確認
	if !strings.Contains(output, "Error: mock conversion failed") {
		t.Errorf("expected error message 'mock conversion failed', but got '%s'", output)
	}
}
