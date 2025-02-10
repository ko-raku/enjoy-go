# 画像処理CLI手順

1. プロジェクト構造を作成し、`go mod init` を実行。
2. 必要なライブラリ（Cobra）をインストール。
3. 画像処理ロジック（`imageprocessor`）を実装。
4. CLIのベース（`cli/root.go` と `cli/grayscale.go`）を実装。
5. `main.go` でCLIを呼び出し、動作確認。
6. 機能拡張や別の画像処理コマンドを追加。

### **基本的な画像処理関数の作成**
- imageprocessor/processor.goに画像処理ロジックを実装
### **CLIのベースを作成**
- CLIエントリーポイント(cli/root.go)を実装
- グレースケール変換コマンド(cli/grayscale.go)を実装
### main.goの作成
### 動作確認
go run main.go grayscale --input m38.jpg


1. Tesseract OCRを動作させるために、TesseractエンジンとGoラッパーライブラリをインストールします。
```shell
   brew install tesseract
   go get -u github.com/otiai10/gosseract/v2
```
