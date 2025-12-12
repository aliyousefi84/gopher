package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


type Downloader struct {
	path string
}

func Newdownloader (path string) Downloader {
	return Downloader{path: path}
}
// concurrency pattern
func (d *Downloader) Download(name,url string) {
	file , err := os.OpenFile(d.path + name , os.O_CREATE | os.O_WRONLY , 0644)
	if err!=nil {
		panic(err)
	}
	defer file.Close()
	resp , err := http.Get(url)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(file , resp.Body)

}