package util

import "os"

// IsExistFile は指定パスにファイルが存在するかチェックする関数
// true: ファイルが存在する false: ファイルが存在しない
func IsExistFile(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
