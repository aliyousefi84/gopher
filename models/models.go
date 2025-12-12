package models

import "time"

type DownloadList struct {
	Name string
	Cap string
	DownloadTime time.Time
	FileType string
	Hash string
}
