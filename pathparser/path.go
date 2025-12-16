package pathparser

import (
	"os"
	"path"
	"strings"
)

func FileName (pathfile string) string {
	fileslice := strings.Split(pathfile , "/")
	nameslice := strings.Split(fileslice[len(fileslice)-1] , ".")
	return nameslice[0]
}

func FileType (pathfile string) string {
	return path.Ext(pathfile)
}

func FileCap (pathfile string) int64 {
	file , err := os.Open(pathfile)
	if err!=nil {
		panic(err)
	}
	cap , _ := file.Stat()
	return cap.Size()
}
 
