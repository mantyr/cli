package cron

import (
	"flag"
	"fmt"

	"github.com/mantyr/cli"
)

// Sync это команда для запуска синхронизации
type Sync struct {
	*flag.FlagSet

	// Адрес конфигурационного файла
	ConfigPath string
}

// NewSync возвращает новую команду
func NewSync() cli.Command {
	cmd := &Sync{
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
func (cmd *Sync) Run(args []string) error {
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

func (cmd *Sync) init() error {
	return nil
}

func (cmd *Sync) run() error {
	fmt.Println("\tCRON SYNC - OK")
	return nil
}
