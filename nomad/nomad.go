package nomad

import (
	"context"

	nomad "github.com/hashicorp/nomad/api"
)

// DispatchJob is a job with parameters
type DispatchJob struct {
	JobID   string            `json:"JobID"`
	Meta    map[string]string `json:"Meta"`
	Payload []byte            `json:"Payload"`
}

func nomadClient() (*nomad.Jobs, error) {
	n, err := nomad.NewClient(nomad.DefaultConfig())
	if err != nil {
		return &nomad.Jobs{}, err
	}
	return n.Jobs(), nil
}

// GetJobs return a list of registered jobs
func GetJobs() ([]*nomad.JobListStub, error) {
	n, err := nomadClient()
	if err != nil {
		return []*nomad.JobListStub{}, err
	}

	jobs, _, err := n.List(nil)
	if err != nil {
		return []*nomad.JobListStub{}, err
	}

	return jobs, nil
}

// GetJob return a registered job
func GetJob(ctx context.Context, id string) (*nomad.Job, error) {
	n, err := nomadClient()
	if err != nil {
		return &nomad.Job{}, err
	}

	job, _, err := n.Info(id, nil)
	if err != nil {
		return &nomad.Job{}, err
	}

	return job, nil
}

// PostJobDispatch implements the Nomad Job Dispatch call
func PostJobDispatch(ctx context.Context, id string, j DispatchJob) (*nomad.JobDispatchResponse, error) {
	n, err := nomadClient()
	if err != nil {
		return &nomad.JobDispatchResponse{}, err
	}

	jobDispatchResponse, _, err := n.Dispatch(id, j.Meta, j.Payload, nil)
	if err != nil {
		return &nomad.JobDispatchResponse{}, err
	}

	return jobDispatchResponse, nil
}

// DeleteJob deletes a rgistered job
func DeleteJob(ctx context.Context, jobID string) error {
	n, err := nomadClient()
	if err != nil {
		return err
	}

	_, _, err = n.Deregister(jobID, false, nil)
	if err != nil {
		return err
	}

	return nil
}
