package imageprocessor

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

// ImageProcessor インターフェイス定義
type ImageProcessor interface {
	ConvertToGray(inputPath, outputPath string) error
}

// DefaultImageProcessor デフォルト実装
type DefaultImageProcessor struct{}

func (p *DefaultImageProcessor) ConvertToGray(inputPath, outputPath string) error {
	// 拡張子を確認
	ext := strings.ToLower(filepath.Ext(inputPath))
	if ext != ".jpeg" && ext != ".jpg" && ext != ".png" {
		return errors.New("unsupported file format: only .jpeg, .jpg, and .png are supported")
	}

	// 入力ファイルを開く
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 画像をデコードする
	img, _, _ := image.Decode(file)

	// グレースケール変換
	bound := img.Bounds()
	newImage := image.NewGray(bound)
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			newImage.Set(x, y, grayColor)
		}
	}

	// 出力ファイルを作成する
	fileOut, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func(fileOut *os.File) {
		err := fileOut.Close()
		if err != nil {
			panic(err)
		}
	}(fileOut)

	// 出力形式に応じて保存する
	switch {
	case ext == ".png":
		if err := png.Encode(fileOut, newImage); err != nil {
			return fmt.Errorf("failed to encode PNG image: %w", err)
		}
	case isJPEG(ext):
		options := jpeg.Options{Quality: 90}
		if err := jpeg.Encode(fileOut, newImage, &options); err != nil {
			return fmt.Errorf("failed to encode JPEG image: %w", err)
		}
	default:
		return errors.New("unsupported file format: only .jpeg, .jpg, and .png are supported")
	}
	// JPEG形式で保存
	err = jpeg.Encode(fileOut, newImage, nil)
	if err != nil {
		return err
	}
	return nil
}

func isJPEG(ext string) bool {
	return ext == ".jpeg" || ext == ".jpg"
}
