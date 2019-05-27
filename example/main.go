package main

import (
	"fmt"
	"os"

	"github.com/mantyr/cli"
	"github.com/mantyr/cli/example/command/cron"
	"github.com/mantyr/cli/example/command/cmd"
)

func main() {
	cli := cli.New("utils")
	group := cli.Group("cron")
	group.Add(
		cron.NewSync(),
		cron.NewNotify(),
	)
	
	group = cli.Group("cli")
	group.Add(
		cmd.NewAdd(),
		cmd.NewDelete(),
	)
	err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
