package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
	var myReader MyReader
	buffer := make([]byte, 1000)
	n, err := myReader.Read(buffer)
	fmt.Println(n, err, string(buffer))

}

func (MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}
