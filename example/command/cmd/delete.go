package cmd

import (
	"flag"
	"fmt"

	"github.com/mantyr/cli"
)

// Delete это команда для удаления данных
type Delete struct {
	*flag.FlagSet

	// Адрес конфигурационного файла
	ConfigPath string
}

// NewDelete возвращает новую команду
func NewDelete() cli.Command {
	cmd := &Delete{
		FlagSet: flag.NewFlagSet("sync", flag.ContinueOnError),
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
func (cmd *Delete) Run(args []string) error {
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

func (cmd *Delete) init() error {
	return nil
}

func (cmd *Delete) run() error {
	fmt.Println("\tCRON SYNC - OK")
	return nil
}
