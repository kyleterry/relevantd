package agent

import (
	"context"
	"errors"
	"time"

	"github.com/kyleterry/relevantd/cli"
	"github.com/kyleterry/relevantd/cue"
)

type Options struct{}

type Agent struct {
	cues    map[string]cue.Cue
	current string
}

func (a *Agent) Run(ctx context.Context) *cli.CommandInteraction {
	ctx, cancel := context.WithCancel(ctx)
	errch := make(chan error, 1)
	donech := make(chan int, 1)

	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				errch <- errors.New("this is a test error")
			case <-ctx.Done():
				donech <- 0
				close(donech)
				close(errch)

				return
			}
		}
	}()

	return &cli.CommandInteraction{
		Cancel: cancel,
		Errch:  errch,
		Donech: donech,
	}
}

func New(cfg Options) (*Agent, error) {
	return &Agent{}, nil
}
