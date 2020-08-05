package appname

import (
	"os"
	"path/filepath"
)

func Folder() (f string, n string) {
	exe, _ := os.Executable()
	f, n = filepath.Split(exe)
	f = f[:len(f)-1] // get rid of trailing backslash of folder
	if len(n) >= 4 && n[len(n)-4:] == ".exe" {
		n = n[:len(n)-4] // get rid of .exe of name
	}
	return
}
