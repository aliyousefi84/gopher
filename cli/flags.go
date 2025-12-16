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
	}

	
	if *url == "" {
		fmt.Println("url should not to be empty")
		os.Exit(1)
	}
	c.svc.DownloadFile(*url)
}

