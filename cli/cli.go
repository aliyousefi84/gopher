package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyousefi84/gopher/downloader"
)



type Cli struct {
	download downloader.Downloader
}

func Newcli (download downloader.Downloader) Cli {
	return Cli{download: download}
}

func (c Cli) Cmd (arg []string) {
	switch arg[0] {
	case "download":
		c.downloadflags(arg[1:])
	default:
		fmt.Println("invalid option !")
		os.Exit(1)
	}
}

func (c Cli) downloadflags (arg []string) {
	gopher := flag.NewFlagSet("gopher",flag.ExitOnError)
	name := gopher.String("n" , "" , "set name for downloads")
	url := gopher.String("u" , "" , "set url")

	err := gopher.Parse(arg)
	if err!=nil {
		fmt.Println(err)
	}

	if *name == "" {
		fmt.Println("name should not to be empty")
		os.Exit(1)
	}
	if *url == "" {
		fmt.Println("url should not to be empty")
		os.Exit(1)
	}
	c.download.Download(*name , *url)

}