# Scheduler

Provides an easy to use abstraction over the Kubernetes and Nomad drivers.

Example:

```golang
import (
  "log"

  "github.com/spaceuptech/scheduler"
  "github.com/spaceuptech/scheduler/config"
)

config := &config.Client{URL: "http://127.0.0.1:4648"}
client, err := scheduler.New(scheduler.Nomad, config)
if err != nil {
  log.Fatal(err)
}

// Use the client here...
```