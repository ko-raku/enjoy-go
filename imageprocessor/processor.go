package imageprocessor

import (
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func ConvertToGray(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, _ := image.Decode(file)

	bound := img.Bounds()
	newImage := image.NewGray(bound)

	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			newImage.Set(x, y, grayColor)
		}
	}

	// 出力ファイルに保存
	fileOut, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	// JPEG形式で保存
	err = jpeg.Encode(fileOut, newImage, nil)
	if err != nil {
		return err
	}
	return nil
}
