package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseIndexTxt(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/index.txt")
	if err != nil {
		panic(err)
	}
	_, err = ParseIndexTxt(data, "SP")
	if err != nil {
		panic(err)
	}
}
