package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aliyousefi84/gopher/downloader"
	"github.com/aliyousefi84/gopher/models"
	"github.com/aliyousefi84/gopher/pathparser"
	"github.com/aliyousefi84/gopher/store"
	"github.com/aliyousefi84/gopher/ui"
)

var (
	dbpath = "gopher"
	downloadpath = fmt.Sprintf("/home/%s/Downloads/" , os.Getenv("USER"))
)

type Svc struct {
	hashlist []string
	downloadlist []models.DownloadList
	download downloader.Downloader
	db *store.DB
}

func Newsvc (db *store.DB) *Svc {
	return &Svc{db: db}
}

func (s *Svc) hash (downloadstr string) string {
	file ,  err := os.Open(downloadstr)
	if err!=nil {
		fmt.Println(err)
	}
	defer file.Close()
	data , err := io.ReadAll(file)
	if err!=nil {
		fmt.Println("error is here")
	}
	hashdata := sha256.Sum256(data)
	return hex.EncodeToString(hashdata[0:])
}

func (s *Svc) DownloadFile (url string) {
	err := s.db.Load(dbpath , &s.downloadlist)
	if err!=nil {
		fmt.Println("please wait , database  initializing...")
	}
	filename := pathparser.Urlstr(url)
	file , err := os.OpenFile(downloadpath + filename , os.O_CREATE | os.O_WRONLY , 0644)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	s.download.Download(url , file)
	datamodel := models.Newmodel(
		pathparser.FileName(downloadpath + filename),
		time.Now().Format(time.RFC3339),
		pathparser.FileType(downloadpath + filename),
		s.hash(downloadpath + filename),
		pathparser.FileCap(downloadpath + filename),
	)
	s.downloadlist = append(s.downloadlist, *datamodel)
	s.db.Save(dbpath , &s.downloadlist)
}

func (s *Svc) ShowHash (filepath string) {
	file , err := os.Open(filepath)
	if err!=nil {
		fmt.Println(err)
	}
	defer file.Close()
	data , _ := io.ReadAll(file)
	hash := sha256.Sum256(data)
	fmt.Println(hex.EncodeToString(hash[0:]))
}

func (s *Svc) DeleteFromDownloadList (index int) {
	err := s.db.Load(dbpath,&s.downloadlist)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.downloadlist = append(s.downloadlist[:index],s.downloadlist[index+1:]...)
	err = s.db.Save(dbpath,&s.downloadlist)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (s *Svc) ShowData () {
	err := s.db.Load(dbpath , &s.downloadlist)
	if err!=nil {
		fmt.Println("database is not initiailizing")
		os.Exit(1)
	}
	ui.Showuser(&s.downloadlist)
}