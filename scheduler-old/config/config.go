package config

// Client is used to create a configuration of a scheduler
type Client struct {
	// URL is the address to connect to the cluster manager
	URL string
}

// Docker is used to specify the configuration of the docker image
type Docker struct {
	// The docker image to use
	Image string
}

// Resources is the resource restrictions
type Resources struct {
	// Specified as a fraction of total number of CPUs x 100 (for Kubernetes) or in Hertz (for Nomad)
	CPU *int

	// Specified in MBs
	Memory *int
}

// Job is used to schedule a job on the scheduler
type Job struct {
	// The name of the job
	Name string

	// The configuration of the docker image
	Docker *Docker

	// The number of tasks we want
	Replicas int

	// The resource restrictions of each task
	Resources *Resources

	// The environment variables to apply to the container
	ENV map[string]string
}
