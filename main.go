package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/masterZSH/grayImage/pkg/fs"
	"github.com/masterZSH/grayImage/pkg/img"
)

var (
	h                     bool
	inputName, outputName string
)

func init() {
	flag.BoolVar(&h, "h", false, "输出帮助")
	flag.StringVar(&inputName, "i", "", "需要转换的图片")
	flag.StringVar(&outputName, "o", "", "输出的图片名称")
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if len(inputName) == 0 {
		log.Fatal("使用-i加图片输入灰色处理的图片\n")
	}
	if _, err := img.CheckImage(inputName); err != nil {
		log.Fatal(err)
	}
	if len(outputName) == 0 {
		var buffer bytes.Buffer
		buffer.WriteString("gray")
		buffer.WriteString(inputName)
		outputName = buffer.String()
	}

	imageReader := fs.ReadFile(inputName)
	image := img.Gray(imageReader)
	img.WriteImage(outputName, image)
}
