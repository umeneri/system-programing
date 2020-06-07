package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	Stream()
}

func handler(w http.ResponseWriter, r *http.Request) {
	// https://stackoverflow.com/questions/46791169/create-serve-over-http-a-zip-file-without-writing-to-disk
	buf := CreateZip()
	fmt.Println(buf.Bytes())

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=new.zip")
	w.Write(buf.Bytes())
}

func HttpRequest() {
	http.HandleFunc("/", handler)
	fmt.Print("hoge")
	http.ListenAndServe(":8080", nil)
	fmt.Print("hoge")
}

func CreateZip() *bytes.Buffer {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	r := strings.NewReader("hogehoge")
	r.Read(buffer)
	fmt.Println(buffer)
	writer.Write(buffer)

	return buf
}

func CopyN(dest io.Writer, src io.Reader, length int) {
	lReader := io.LimitReader(src, int64(length))
	io.Copy(dest, lReader)
}

func CopyTest() {
	src := strings.NewReader("1234567890")
	CopyN(os.Stdout, src, 5)
}

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func Stream() {
	var stream io.Reader

	var bufs = map[string]*bytes.Buffer{
		"a":  new(bytes.Buffer),
		"s":  new(bytes.Buffer),
		"c":  new(bytes.Buffer),
		"i":  new(bytes.Buffer),
		"i1": new(bytes.Buffer),
		"i2": new(bytes.Buffer),
	}

	// a
	buffer := make([]byte, 1024)
	programming.ReadAt(buffer, 5)
	_, _ = io.CopyN(bufs["a"], bytes.NewBuffer(buffer), 1)
	// c
	_, _ = io.CopyN(bufs["c"], computer, 1)

	// s
	_, _ = io.CopyN(bufs["s"], system, 1)

	// i
	buffer2 := make([]byte, 1024)
	programming.ReadAt(buffer2, 8)
	_, _ = io.CopyN(bufs["i"], bytes.NewBuffer(buffer2), 1)
	io.Copy(io.MultiWriter(bufs["i1"], bufs["i2"]), bufs["i"])

	stream = io.MultiReader(bufs["a"], bufs["s"], bufs["c"], bufs["i1"], bufs["i2"])
	io.Copy(os.Stdout, stream)
}
