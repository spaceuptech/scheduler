package scheduler

import (
	"context"
	"errors"

	"github.com/spaceuptech/scheduler/config"
	"github.com/spaceuptech/scheduler/nomad"
)

// Client is the client to talk to the scheduler
type Client interface {
	StartJob(ctx context.Context, spec *config.Job) error
}

// New creates a new instance of the scheduler client
func New(kind Kind, clientConfig *config.Client) (Client, error) {
	switch kind {
	case Nomad:
		return nomad.New(clientConfig)
	default:
		return nil, errors.New("invalid scheduler kind provided")
	}
}
