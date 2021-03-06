package appname

import (
	"log"
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

	ff, err := os.OpenFile(FolderAndSlash+Name+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	log.SetOutput(ff)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func Panicln(v ...interface{}) {
	log.Panicln(v...)
}

func Print(v ...interface{}) {
	log.Print(v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func Printf(s string, v ...interface{}) {
	log.Printf(s, v...)
}
