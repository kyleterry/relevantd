package cue

type Cue interface {
	Run() (string, error)
}
