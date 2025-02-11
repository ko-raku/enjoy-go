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
```shell
go run main.go grayscale --input m38.jpg
go run main.go textfromimage --input image_1.png
```


1. Tesseract OCRを動作させるために、TesseractエンジンとGoラッパーライブラリをインストールします。
```shell
   brew install tesseract
   go get -u github.com/otiai10/gosseract/v2
```

## MEMO
1. Go モジュールの管理や依存関係に何らかの問題が生じた場合
#### 1. 定期的に `go mod tidy` を実行
開発中に必要のない依存関係が残ることを防ぐため、定期的に以下を実行すると良いでしょう。
``` bash
go mod tidy
```
#### 2. モジュールキャッシュを定期的にクリアする
長期間開発を続けるとモジュールキャッシュが大きくなり、壊れる場合があります。必要に応じて以下を実行してキャッシュをクリアしてください。
``` bash
go clean -modcache
```
#### 3. 依存関係の変更後は必ず再ビルドを行う
依存関係に変更が生じた場合、以下を実行し、問題がないかを確認します。
``` bash
go build ./...
```
#### 4. 依存関係の状態や問題を確認
依存関係の詳細を確認するには、以下を実行すると、モジュールの依存関係や問題点がわかります。
``` bash
go list -m all
```
問題がある場合には以下を試してください。
``` bash
go mod verify
```


