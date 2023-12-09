package async_job

import (
	"context"
	"time"
)

// Job requirements:
// 1. Job can do sth
// 2. Job can retry
// 		2.1 Config retry times and duration
// 3. Should be stateful
// 4. We should have job manager to manage jobs

type JobState int

const (
	defaultMaxTimeout = time.Second * 10
)

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

var (
	defaultRetryTime = []time.Duration{
		time.Second,
		time.Second * 5,
		time.Second * 10,
	}
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

type JobHandler func(ctx context.Context) error

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		stopChan:   make(chan bool),
	}

	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	if err := j.handler(ctx); err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted

	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	// handle fail
	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return nil
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) SetRetryDurations(times []time.Duration) {
	j.config.Retries = times
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}
