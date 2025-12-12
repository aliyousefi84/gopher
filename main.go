package main

import (
	"os"

	"github.com/aliyousefi84/gopher/cli"
	"github.com/aliyousefi84/gopher/downloader"
)

func main () {
	d := downloader.Newdownloader("/home/ali/Downloads/")
	cli := cli.Newcli(d)
	
	if len(os.Args) > 1 {
		cli.Cmd(os.Args[1:])
	}
}