package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsExistFile(t *testing.T) {
	tempDir := os.TempDir()
	fp, err := ioutil.TempFile(tempDir, "PecaRecoderGo-TestExistFile")
	path := fp.Name()

	defer fp.Close()
	if err != nil {
		t.Error(err)
	}
	fp.WriteString("Hello World!\n")
	if !IsExistFile(path) {
		os.Remove(path)
		t.Error("ファイルが存在しているがExistFileでは存在しないことになっています.")
		return
	}
	os.Remove(path)
}
