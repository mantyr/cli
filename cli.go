package cli

import (
	"errors"
	"fmt"
	"strings"
)

// Command это интерфейс для встраивания приложения в cli утилиту
type Command interface {
	// Name возвращает название команды по которой определяется обработчик
	Name() string

	// Run запускает обработку запроса
	// args это параметры командной строки
	Run(args []string) error

	// PrintDefault печатает настройки по умолчанию
	PrintDefaults()
}

type CLI struct {
	name     string
	err      error
	commands map[string]Command
	list     []Command
}

func New(
	programName string,
) *CLI {
	return &CLI{
		name:     programName,
		commands: make(map[string]Command),
	}
}

func (c *CLI) Name() string {
	return c.name
}

func (c *CLI) Group(name string) *CLI {
	group := New(
		name,
	)
	c.Add(group)
	return group
}

func (c *CLI) Add(commands ...Command) {
	if c.err != nil {
		return
	}
	for n, command := range commands {
		if command == nil {
			c.err = fmt.Errorf("empty #%d command", n)
			return
		}
		name := strings.TrimSpace(command.Name())
		if name == "" {
			c.err = fmt.Errorf("empty #%d command name", n)
			return
		}
		if _, ok := c.commands[name]; ok {
			c.err = fmt.Errorf("command %s is already exists", name)
			return
		}
		c.commands[name] = command
		c.list = append(c.list, command)
	}
}

func (c *CLI) Run(args []string) error {
	if c.err != nil {
		c.PrintDefaults()
		return c.err
	}
	if len(args) == 0 {
		c.PrintDefaults()
		return errors.New("empty args")
	}
	name := strings.TrimSpace(args[0])
	switch name {
	case "help", "-help", "--help":
		c.PrintDefaults()
		return nil
	}
	command, ok := c.commands[args[0]]
	if ok {
		return command.Run(args[1:])
	}
	c.PrintDefaults()
	return fmt.Errorf("unexpected command %s", name)
}

func (c *CLI) PrintDefault(command Command) {
	fmt.Printf(
		"\nUSAGE\n  %s %s\n\nOPTIONS\n",
		c.name,
		strings.TrimSpace(command.Name()),
	)
	command.PrintDefaults()
	fmt.Printf("\n")
}

/*
SYNOPSIS
  merchants sync
  merchants exists
  merchants add
  merchants edit

USAGE
  merchants sync

OPTIONS
  -config string
    Адрес конфигурационного файла для merchants sync (default "./merchants.conf")

*/

func (c *CLI) PrintDefaults() {
	fmt.Printf("SYNOPSIS\n")
	for _, command := range c.list {
		fmt.Printf(
			"  %s %s\n",
			c.name,
			strings.TrimSpace(command.Name()),
		)
	}
	for _, command := range c.list {
		fmt.Printf(
			"\nUSAGE\n  %s %s\n\nOPTIONS\n",
			c.name,
			strings.TrimSpace(command.Name()),
		)
		command.PrintDefaults()
	}
	fmt.Printf("\n")
}
