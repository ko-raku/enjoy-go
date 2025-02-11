# imagetool - 画像処理とOCR CLIツール

`imagetool`は、画像をグレースケールに変換したり、OCRを使用して画像からテキストを抽出するためのコマンドラインツールです。

## 特徴
- 画像をグレースケールに変換
- 前処理（グレースケール化・解像度調整）を行い、OCRでテキストを抽出
- Tesseract OCRエンジンを利用
- 簡単なCLIコマンド操作

## 必要な環境
- Tesseract OCRエンジン（[公式サイト](https://github.com/tesseract-ocr/tesseract)からインストールしてください）

### 各OSでのTesseractインストール方法
#### macOS
```bash
brew install tesseract
```

#### Linux (Ubuntu)
```bash
sudo apt update
sudo apt install tesseract-ocr
```

#### Windows
- [Tesseract公式インストーラー](https://github.com/tesseract-ocr/tesseract) をダウンロードしてインストール

## インストール手順
1. アーカイブをダウンロードします（例: `imagetool-package.zip`）。
2. 解凍後、ターミナルまたはコマンドプロンプトで実行可能ファイルを実行します。

例：
```bash
./imagetool --help
```

## 使い方
### グレースケール変換
```bash
imagetool grayscale --input example.jpg --output example_gray.jpg
```

### 画像からOCRテキスト抽出
```bash
imagetool textfromimage --input example.jpg --output output.txt
```

## 注意事項
- 現在`.jpeg`、`.jpg`、`.png`形式の画像にのみ対応。
- 複数言語のOCRには適切なTesseract言語モデルが必要です。

## 著者
このプロジェクトは [yunaction](https://github.com/ko-raku) によって開発されました。

## サポート
問題が発生した場合は、リポジトリの[Issueセクション](https://github.com/ko-raku/enjoy-go/issues)にご報告ください。