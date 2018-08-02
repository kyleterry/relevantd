package cli

import "context"

type Command interface {
	Run(ctx context.Context) *CommandInteraction
}

type CommandInteraction struct {
	Cancel context.CancelFunc
	Errch  <-chan error
	Donech <-chan int
}
