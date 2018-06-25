package main

import "io"

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
