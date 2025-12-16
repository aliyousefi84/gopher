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

func (d *Downloader) Download(url string,file *os.File) {
	req , err := http.NewRequest("GET" , url,nil)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := http.Client{}
	resp , err := client.Do(req)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(file , resp.Body)
	
}