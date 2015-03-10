// +build darwin freebsd linux netbsd openbsd

package util

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

// GetApplicationConfigDirectory アプリケーションのコンフィグディレクトリを取得する
func GetApplicationConfigDirectory(appName string) (string, error) {
	var p string

	if runtime.GOOS == "darwin" {
		p = "~/Library/Application Support/"
	} else {
		xdgDataHome := os.Getenv("XDG_DATA_HOME")
		if len(xdgDataHome) > 0 {
			p = xdgDataHome
		} else {
			p = "eval echo ~/.local/share"
		}
	}

	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", p)
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	appDirectory := strings.TrimSpace(stdout.String())
	if len(appDirectory) == 0 {
		return "", errors.New("アプリケーションコンフィグディレクトリの取得に失敗しました")
	}
	return path.Join(appDirectory, appName), nil
}
