package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	_, err = rand.Reader.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println(buffer)

	file.Write(buffer)

}
