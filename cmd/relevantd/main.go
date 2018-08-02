package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/kr/pretty"
	"github.com/kyleterry/relevantd/agent"
	"github.com/kyleterry/relevantd/cli"
	"github.com/kyleterry/relevantd/client"
	"github.com/kyleterry/relevantd/config"
	"github.com/pkg/errors"
)

type commandType string

const (
	agentCmd   commandType = "agent"
	currentCmd commandType = "current"
	showCmd    commandType = "show"
	nextCmd    commandType = "next"
	prevCmd    commandType = "prev"
)

func SelectedCmd(cfg *config.Config) (cli.Command, error) {
	args := flag.Args()

	if len(args) < 1 {
		c := client.New(client.Options{})

		return c.Show, nil
	}

	cmdStr := commandType(strings.ToLower(args[0]))

	if cmdStr == agentCmd {
		a, err := agent.New(agent.Options{})
		if err != nil {
			return nil, errors.Wrap(err, "failed to create agent command")
		}

		return a, nil
	} else {
		c := client.New(client.Options{})

		switch cmdStr {
		case currentCmd:
			return c.Current, nil
		case showCmd:
			return c.Show, nil
		case nextCmd:
			return c.Next, nil
		case prevCmd:
			return c.Prev, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("command %s not found", cmdStr))
}

func main() {
	defaultConfigPath := filepath.Join(os.Getenv("HOME"), ".config", "relevantd", "relevantd.yaml")
	configPath := flag.String("config", defaultConfigPath, "Path to configuration file")

	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		log.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	pretty.Println(cfg)

	cmd, err := SelectedCmd(cfg)
	if err != nil {
		log.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	interaction := cmd.Run(context.Background())

	sigch := make(chan os.Signal, 1)
	exit := make(chan int, 1)

	signal.Notify(
		sigch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	go func() {
		select {
		case <-sigch:
			interaction.Cancel()
			exit <- <-interaction.Donech
		case <-interaction.Donech:
			exit <- 0
		case err := <-interaction.Errch:
			log.Println(err)
			exit <- 1
		}

		close(exit)
	}()

	code := <-exit
	os.Exit(code)
}
