package img

import (
	"errors"
	"image"
	"image/color"

	// jpeg、png支持
	"image/jpeg"
	"image/png"
	"path"

	"io"
	"log"
	"os"
)

// WriteImage image写入图片文件
func WriteImage(outputName string, image image.Image) {
	f, err := os.Create(outputName)
	if err != nil {
		log.Fatal(err)
	}
	if err := Encode(outputName, f, image); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// Gray 输出的
func Gray(ir io.Reader) image.Image {
	// 读取img
	img, _, err := image.Decode(ir)
	if err != nil {
		log.Fatal(err)
	}
	// 图片宽度
	width := img.Bounds().Dx()
	// 图片高度
	height := img.Bounds().Dy()
	// 新的rgba图片
	newImage := image.NewRGBA(img.Bounds())
	// 循环图片像素
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// 获取图片像素颜色
			pixelColor := img.At(i, j)
			originalColor := color.RGBAModel.Convert(pixelColor).(color.RGBA)
			// 灰度化
			r := float64(originalColor.R) * 0.92020
			g := float64(originalColor.G) * 0.90404
			b := float64(originalColor.B) * 0.90200
			gray := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: gray, G: gray, B: gray, A: originalColor.A,
			}
			// 设置颜色
			newImage.Set(i, j, c)
		}
	}
	return newImage
}

// CheckImage 检查图片后缀
func CheckImage(imageName string) (string, error) {
	// 获取文件后缀
	ext := path.Ext(imageName)
	if ext == ".jpg" || ext == ".jpeg" {
		return "jpeg", nil
	}
	if ext == ".png" {
		return "png", nil
	}
	return "", errors.New("不支持的图片类型")
}

// Encode 编码图片到io.Writer
func Encode(imageName string, w io.Writer, m image.Image) error {
	imageType, err := CheckImage(imageName)
	if err != nil {
		return err
	}
	switch imageType {
	case "png":
		return png.Encode(w, m)
	case "jpeg":
		return jpeg.Encode(w, m, &jpeg.Options{100})
	default:
		return errors.New("不支持的类型")
	}

}
