package scheduler

// Kind defines the type of the scheduler
type Kind string

const (

	// Nomad kind instructs the scheduler to talk to nomad
	Nomad Kind = "nomad"
)
