package cli_test

import (
	"bytes"
	"enjoy/cli"
	"testing"
)

func TestExecute_Success(t *testing.T) {
	// モックの準備
	var exitCode int
	cli.SetExitFunc(func(code int) {
		exitCode = code // exitCode に値を保存（os.Exitを呼ばない）
	})
	defer cli.SetExitFunc(nil) // テスト終了後にデフォルトを復元

	// 標準出力をキャプチャ
	buf := &bytes.Buffer{}
	cli.RootCmd.SetOut(buf)

	// コマンド実行
	cli.Execute()

	// 正常終了の検証
	if exitCode != 0 {
		t.Errorf("expected exit code 0, but got %d", exitCode)
	}

	// 出力の確認
	output := buf.String()
	if len(output) == 0 {
		t.Errorf("expected output from Execute, but got empty output")
	}
}

func TestExecute_Error(t *testing.T) {
	// モックの準備
	var exitCode int
	cli.SetExitFunc(func(code int) {
		exitCode = code // 終了コードを確認
	})
	defer cli.SetExitFunc(nil)

	// 標準出力・エラー出力のキャプチャ
	buf := &bytes.Buffer{}
	cli.RootCmd.SetOut(buf)
	cli.RootCmd.SetErr(buf)

	// 存在しないコマンドを設定（エラーケースを作り出す）
	cli.RootCmd.SetArgs([]string{"non-existent-command"})

	// Execute を実行
	cli.Execute()

	// Exit のコードを検証
	if exitCode != 1 {
		t.Errorf("expected exit code 1, but got %d", exitCode)
	}
}
