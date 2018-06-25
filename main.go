package main

import (
	"io"
	"os"
	"fmt"
)

type alphaReader struct {
	reader io.Reader
}

func newAlphaReader(reader io.Reader) alphaReader {
	return alphaReader{reader:reader}
}

func onlyAlpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r == 32) {
		return r
	}

	return 0
}

func (a alphaReader) Read(bs []byte) (int, error) {
	n, err := a.reader.Read(bs)

	if err != nil {
		return n, err
	}

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		if char := onlyAlpha(bs[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(bs, buf)
	return n, nil
}

func main() {
	//read in the arguments provided to the program
	fileName := os.Args[1]

	//example of using an os.File as source for alphaReader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//example of using a string.Reader as source for alphaReader
	//reader := newAlphaReader(strings.NewReader("I can eat up to 10000 sushi!"))

	reader := newAlphaReader(file)
	bs := make([]byte, 4)
	s := ""

	for {
		n, err := reader.Read(bs)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		s = s + string(bs[:n])
	}

	fmt.Println(s)
}