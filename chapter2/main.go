package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmtPrint()
}

func OsWrite() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))

	os.Stdout.Write([]byte("os.Stdout example\n"))
	os.Stderr.Write([]byte("os.Stdout example error\n"))
	file.Close()
}

func ByteBuffer() {
	var buffer bytes.Buffer

	// buffer.Write([]byte("bytes.Buffer  example\n"))
	buffer.WriteString("bytes.Buffer  example\n")
	fmt.Println(buffer.String())
	fmt.Println(buffer.Len())
}

func BuilderWrite() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}

func NetWrite() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET /HTTP/1.0\r\nHOST: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
	// http.ResponseWriter inherit io.Writer
}

func HttpRequest() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func MultiWrite() {
	var filename = "multiwriter.txt"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
	os.Remove(filename)
}

func GzipWrite() {
	var filename = "test.txt.gz"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// writer is compreessed
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()

	os.Remove(filename)
}

func BufioWrite() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.WriteString("example\n")
	buffer.Flush()
}

func fmtPrint() {
	fmt.Fprintf(os.Stdout, "Write %v", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello": "world",
	})
}
