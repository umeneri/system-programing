package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x0, 0x0, 0x27, 0x10}

	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("%b\n", data)

	for i := 0; i < len(data); i++ {
		fmt.Printf("%d, ", data[i])
	}
	fmt.Printf("\n")

	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", int16(i))
	fmt.Printf("%d\n", uint8(i))
	fmt.Printf("%d\n", byte(i))
	fmt.Printf("%d\n", byte(2))
	fmt.Printf("%d\n", byte(3))
	fmt.Printf("%d\n", byte(4))
	fmt.Printf("%d\n", byte(255))

}
