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
	// strange problem: if FolderAndSlash is used instead of f, the change in FolderAndSlash is not visible outside of this function.
	f, n := filepath.Split(exe)
	FolderAndSlash = f
	NameWithExe = n
	FilePath = FolderAndSlash + NameWithExe
	Folder = FolderAndSlash[:len(FolderAndSlash)-1] // get rid of trailing backslash of folder
	Name = NameWithExe
	if len(Name) >= 4 && Name[len(Name)-4:] == ".exe" {
		Name = Name[:len(Name)-4] // get rid of .exe of name
	}

}
