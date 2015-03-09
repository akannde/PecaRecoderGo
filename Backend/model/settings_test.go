package model

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/PyYoshi/PecaRecoderGo/Backend/util"
)

func TestEncodeSettings(t *testing.T) {
	setting := NewSettings()
	buffer, err := setting.Encode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", buffer.String())
}

func TestSaveSettings(t *testing.T) {
	tempDir := os.TempDir()
	path := filepath.Join(tempDir, "PecaRecoderGo.tml")
	setting := NewSettings()
	setting.Save(path)
	if !util.IsExistFile(path) {
		os.Remove(path)
		t.Error("設定ファイルが正しく保存されていません.")
		return
	}
	os.Remove(path)
}
