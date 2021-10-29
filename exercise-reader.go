package main

import (
	"io"
	"golang.org/x/tour/reader"
)

type MyReader struct{}


func (r MyReader) Read(b []byte) (int, error) {
	if b == nil || len(b) == 0 {
		return 0, io.EOF	
	}
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
