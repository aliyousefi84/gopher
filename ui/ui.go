package ui

import (
	"os"
	"strconv"

	"github.com/aliyousefi84/gopher/models"
	"github.com/aquasecurity/table"
)

func Showuser (downloadlist *[]models.DownloadList) {
	t := table.New(os.Stdout)
	t.SetHeaders("Name","Cap","DownloadTime","FileType","Hash (sha256)")
	for _ , value := range *downloadlist {
		t.AddRow(value.Name , strconv.Itoa(int(value.Cap)) , value.DownloadTime , value.FileType , value.Hash)
	}
	t.Render()
}