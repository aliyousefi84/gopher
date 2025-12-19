package cli

import (
	"flag"
	"fmt"
	"os"
)

func (c Cli) downloadflags (arg []string) {
	gopher := flag.NewFlagSet("gopher",flag.ExitOnError)

	url := gopher.String("u" , "" , "")
	
	err := gopher.Parse(arg)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	
	if *url == "" {
		fmt.Println("url should not to be empty")
		os.Exit(1)
	}
	c.svc.DownloadFile(*url)
}

func (c Cli) hashfile (arg []string) {
	gopher := flag.NewFlagSet("gopher" , flag.ExitOnError)
	hash := gopher.String("h" , "" , "hash file and show on terminal")

	err := gopher.Parse(arg)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *hash == "" {
		fmt.Println("miss filepath !")
	}
	c.svc.ShowHash(*hash)
}

func (c Cli) delfromlist (arg []string) {
	gopher := flag.NewFlagSet("gopher" , flag.ExitOnError)
	del := gopher.Int("i" , -1 , "for deleting")

	err := gopher.Parse(arg)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *del == -1 {
		fmt.Println("you should set index for deleting from list")
	}
	c.svc.DeleteFromDownloadList(*del)
}

