// +build windows

package util

import (
	"path"
	"syscall"
	"unsafe"
)

// CSIDL_APPDATA https://technet.microsoft.com/ja-jp/library/cc749104%28v=ws.10%29.aspx
const CSIDL_APPDATA = 26

// GetApplicationConfigDirectory アプリケーションディレクトリの取得
func GetApplicationConfigDirectory(appName string) (string, error) {
	var (
		shell         = syscall.MustLoadDLL("Shell32.dll")
		getFolderPath = shell.MustFindProc("SHGetFolderPathW")
	)
	b := make([]uint16, syscall.MAX_PATH)
	r, _, err := getFolderPath.Call(0, CSIDL_APPDATA, 0, 0, uintptr(unsafe.Pointer(&b[0])))
	if uint32(r) != 0 {
		return "", Error(err)
	}
	return path.Join(syscall.UTF16ToString(b), appName), nil
}
