package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func open() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "new file content\n")
	io.Copy(os.Stdout, file)

	file, err = os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, file)
	os.Remove("text.txt")

}

func benchmark() {
	f, _ := os.Create("file.txt")
	a := time.Now()
	f.Write([]byte("緑の怪獣 "))
	b := time.Now()
	f.Sync()
	c := time.Now()
	f.Close()
	d := time.Now()
	fmt.Printf("Write: %v\n", b.Sub(a))
	fmt.Printf("Sync: %v\n", c.Sub(b))
	fmt.Printf("Close: %v\n", d.Sub(c))
}

func directory() {
	os.Mkdir("setting", 0755)
	os.MkdirAll("setting/myapp/networksettings", 0755)
	os.RemoveAll("setting")

}

func rename() {
	os.Rename("file.txt", "tmp/sample.txt")
}

func stat() {
	info, err := os.Stat("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Name())
	fmt.Println(info.Size())
}

func traverse()  {
	dir, err := os.Open("/")
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("dir %s\n", fileInfo.Name())
		} else {
			fmt.Printf("file %s\n", fileInfo.Name())
		}
	}
}

func main() {
	// benchmark()
	// directory()
	// rename()
	// stat()
	traverse()
}
