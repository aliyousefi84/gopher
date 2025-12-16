package models


type DownloadList struct {
	Name string
	Cap int64
	DownloadTime string
	FileType string
	Hash string
}

func Newmodel (name , time , filetype , hash string,cap int64) *DownloadList {
	return &DownloadList{
		Name: name,
		DownloadTime: time,
		FileType: filetype,
		Hash: hash,
		Cap: cap,
	}
}