package nomad

import (
	"context"

	"github.com/hashicorp/nomad/api"
	uuid "github.com/satori/go.uuid"

	"github.com/spaceuptech/scheduler/config"
)

// StartJob starts a job task on the nomad cluster
func (d *Driver) StartJob(ctx context.Context, spec *config.Job) error {
	// Create a unique id
	id := uuid.NewV1().String()

	// Create a task
	task := api.NewTask(spec.Name, "docker")
	task.SetConfig("image", spec.Docker.Image)
	task.Env = spec.ENV

	// Set the resources if provided
	if spec.Resources != nil {
		resources := new(api.Resources)
		if spec.Resources.CPU != nil {
			resources.CPU = spec.Resources.CPU
		}

		if spec.Resources.Memory != nil {
			resources.MemoryMB = spec.Resources.Memory
		}

		task.Resources = resources
	}

	// Create a task group
	taskGroup := api.NewTaskGroup(spec.Name, spec.Replicas)
	taskGroup.AddTask(task)

	// Create a job
	job := api.NewBatchJob(id, spec.Name, "global", 50)
	job.AddTaskGroup(taskGroup)

	// Register the task
	_, _, err := d.client.Jobs().Register(job, nil)
	return err
}
