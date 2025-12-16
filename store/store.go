package store

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aliyousefi84/gopher/models"
)

type DB struct {}

func NewDB () *DB {
	return &DB{}
}

func (d *DB) Save (path string,downloadlist *[]models.DownloadList) error {
	data , err := json.MarshalIndent(downloadlist , "" , "  ")
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return os.WriteFile(path , data , 0644)
}


func (d *DB) Load (path string,downloadlist *[]models.DownloadList) error {
	data ,err := os.ReadFile(path)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return json.Unmarshal(data , &downloadlist)	
}