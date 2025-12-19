package ui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aliyousefi84/gopher/models"
	"github.com/aquasecurity/table"
)

func Showuser (downloadlist *[]models.DownloadList) {
	t := table.New(os.Stdout)
	t.SetHeaders("#","Name","Cap","DownloadTime","FileType","Hash (sha256)")
	for index , value := range *downloadlist {
		t.AddRow(strconv.Itoa(index),value.Name , strconv.Itoa(int(value.Cap)) , value.DownloadTime , value.FileType , value.Hash)
	}
	t.Render()
}

func Showcommand () {
	fmt.Println(`
gopher download -u [URL]  	for download somthing
gopher list		download list
gopher hash -h [filepath]
gopher del -i [your index list]

	`)
}