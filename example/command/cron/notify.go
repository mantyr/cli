package cron

import (
	"flag"
	"fmt"

	"github.com/mantyr/cli"
)

// Notify это команда для запуска синхронизации
type Notify struct {
	*flag.FlagSet

	// Адрес конфигурационного файла
	ConfigPath string
}

// NewNotify возвращает новую команду
func NewNotify() cli.Command {
	cmd := &Notify{
		FlagSet: flag.NewFlagSet("notify", flag.ContinueOnError),
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
func (cmd *Notify) Run(args []string) error {
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

func (cmd *Notify) init() error {
	return nil
}

func (cmd *Notify) run() error {
	fmt.Println("\tCRON NOTIFY - OK")
	return nil
}
