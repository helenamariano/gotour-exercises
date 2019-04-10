package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read emits an infinite stream of the ASCII character 'A'
func (r MyReader) Read(b []byte) (n int, err error) {

	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
