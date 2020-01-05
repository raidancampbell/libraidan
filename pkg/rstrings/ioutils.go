package rstrings

import (
	"io/ioutil"
	"os"
)

func FileToString(fname string) (result string, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}
	err = file.Close()

	result = string(b)
	return
}
