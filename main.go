package main

import (
	"os"

	"github.com/aliyousefi84/gopher/cli"
	"github.com/aliyousefi84/gopher/service"
	"github.com/aliyousefi84/gopher/store"
)

func main () {
	db := store.NewDB()
	svc := service.Newsvc(db)
	roorcli := cli.Newcli(svc)
	if len(os.Args) > 1 {
		roorcli.Cmd(os.Args[1:])
	}
}