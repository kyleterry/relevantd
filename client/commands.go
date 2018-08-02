package client

import (
	"context"
	"fmt"

	"github.com/kyleterry/relevantd/cli"
)

type currentCmd struct{}

func (c currentCmd) Run(ctx context.Context) *cli.CommandInteraction {
	ctx, cancel := context.WithCancel(ctx)
	errch := make(chan error, 1)
	donech := make(chan int, 1)

	go func() {
		fmt.Println("I'm the current command")
		donech <- 0
		close(donech)
	}()

	return &cli.CommandInteraction{
		Cancel: cancel,
		Errch:  errch,
		Donech: donech,
	}
}

type showCmd struct{}

func (c showCmd) Run(ctx context.Context) *cli.CommandInteraction {
	ctx, cancel := context.WithCancel(ctx)
	errch := make(chan error, 1)
	donech := make(chan int, 1)

	go func() {
		fmt.Println("I'm the show command")
		donech <- 0
		close(donech)
	}()

	return &cli.CommandInteraction{
		Cancel: cancel,
		Errch:  errch,
		Donech: donech,
	}
}

type nextCmd struct{}

func (c nextCmd) Run(ctx context.Context) *cli.CommandInteraction {
	ctx, cancel := context.WithCancel(ctx)
	errch := make(chan error, 1)
	donech := make(chan int, 1)

	go func() {
		fmt.Println("I'm the next command")
		donech <- 0
		close(donech)
	}()

	return &cli.CommandInteraction{
		Cancel: cancel,
		Errch:  errch,
		Donech: donech,
	}
}

type prevCmd struct{}

func (c prevCmd) Run(ctx context.Context) *cli.CommandInteraction {
	ctx, cancel := context.WithCancel(ctx)
	errch := make(chan error, 1)
	donech := make(chan int, 1)

	go func() {
		fmt.Println("I'm the prev command")
		donech <- 0
		close(donech)
	}()

	return &cli.CommandInteraction{
		Cancel: cancel,
		Errch:  errch,
		Donech: donech,
	}
}
