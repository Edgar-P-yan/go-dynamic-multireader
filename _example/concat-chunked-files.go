package main

import (
	"io"
	"os"

	dynamicmultireader "github.com/Edgar-P-yan/go-dynamic-multireader"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	multiReader := dynamicmultireader.DynamicMultiReader(func(n int) io.Reader {
		switch n {
		case 0:
			r, err := os.Open(wd + "/chunk1.txt")
			if err != nil {
				panic(err)
			}
			return r
		case 1:
			r, err := os.Open(wd + "/chunk2.txt")
			if err != nil {
				panic(err)
			}
			return r
		case 2:
			r, err := os.Open(wd + "/chunk3.txt")
			if err != nil {
				panic(err)
			}
			return r
		default:
			return nil
		}
	})

	processed := UpperCase(multiReader)

	file, err := os.Create(wd + "/uppercased.txt")

	if err != nil {
		panic(err)
	}

	io.Copy(file, processed)
}

// This is just a simple io.Reader that uppercases all english letters.
func UpperCase(in io.Reader) (out io.Reader) {
	return &upper{in}
}

type upper struct {
	in io.Reader
}

func (u *upper) Read(p []byte) (n int, err error) {
	n, err = u.in.Read(p)

	for i, b := range p {
		if b > 97 && b < 122 {
			p[i] = b - 32
		}
	}

	return
}
