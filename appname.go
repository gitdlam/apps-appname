package appname

import (
	"os"
	"path/filepath"
)

var (
	Folder         string
	FolderAndSlash string
	Name           string
	NameWithExe    string
	FilePath       string
)

func init() {
	exe, _ := os.Executable()
	FolderAndSlash, NameWithExe := filepath.Split(exe)
	FilePath = FolderAndSlash + NameWithExe
	Folder = FolderAndSlash[:len(FolderAndSlash)-1] // get rid of trailing backslash of folder
	Name = NameWithExe
	if len(Name) >= 4 && Name[len(Name)-4:] == ".exe" {
		Name = Name[:len(Name)-4] // get rid of .exe of name
	}
}
