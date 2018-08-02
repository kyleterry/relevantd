package client

type Options struct{}

type Client struct {
	Current currentCmd
	Show    showCmd
	Next    nextCmd
	Prev    prevCmd
}

func New(opts Options) *Client {
	return &Client{
		Current: currentCmd{},
		Show:    showCmd{},
		Next:    nextCmd{},
		Prev:    prevCmd{},
	}
}
