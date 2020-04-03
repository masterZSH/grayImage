package fs

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
)

// 文件处理包

// ReadFile 读取文件返回io.reader
func ReadFile(fileName string) io.Reader {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	bsReader := bytes.NewReader(bs)
	return bsReader
}
