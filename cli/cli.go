package cli

import (
	"fmt"
	"os"

	"github.com/aliyousefi84/gopher/service"
)



type Cli struct {
	svc *service.Svc
}

func Newcli (svc *service.Svc) Cli {
	return Cli{svc: svc}
}

func (c Cli) Cmd (arg []string) {
	switch arg[0] {
	case "download":
		c.downloadflags(arg[1:])
	case "list":
		c.svc.ShowData()
	case "del":
		c.delfromlist(arg[1:])
	case "hash":
		c.hashfile(arg[1:])
	default:
		fmt.Println("invalid option !")
		os.Exit(1)
	}
}