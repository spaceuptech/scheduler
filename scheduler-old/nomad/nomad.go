package nomad

import (
	"github.com/hashicorp/nomad/api"

	"github.com/spaceuptech/scheduler/config"
)

// Driver is the abstraction for using the nomad driver
type Driver struct {
	client *api.Client
}

// New creates a new instance of the nomad driver
func New(c *config.Client) (*Driver, error) {

	// Create a new nomad client config
	config := api.DefaultConfig()
	config.Address = c.URL

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Driver{client}, nil
}
