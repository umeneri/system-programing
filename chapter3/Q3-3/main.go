package main

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

func main() {
	dest, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	zipWriter := zip.NewWriter(dest)
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
}
