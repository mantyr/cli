package cmd

import (
	"flag"
	"fmt"

	"github.com/mantyr/cli"
)

// Add это команда для добавления данных
type Add struct {
	*flag.FlagSet

	// Адрес конфигурационного файла
	ConfigPath string
}

// NewAdd возвращает новую команду
func NewAdd() cli.Command {
	cmd := &Add{
		FlagSet: flag.NewFlagSet("add", flag.ContinueOnError),
	}
	cmd.FlagSet.StringVar(
		&cmd.ConfigPath,
		"config",
		"./config.conf",
		"Адрес конфигурационного файла",
	)
	return cmd
}

// Run запускает обработку команды
func (cmd *Add) Run(args []string) error {
	err := cmd.FlagSet.Parse(args)
	if err != nil {
		return err
	}
	err = cmd.init()
	if err != nil {
		return err
	}
	return cmd.run()
}

func (cmd *Add) init() error {
	return nil
}

func (cmd *Add) run() error {
	fmt.Println("\tCRON NOTIFY - OK")
	return nil
}
