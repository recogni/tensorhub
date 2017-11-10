package main

////////////////////////////////////////////////////////////////////////////////

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

////////////////////////////////////////////////////////////////////////////////

func newUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

////////////////////////////////////////////////////////////////////////////////

// Job describes either a pending or current tensorboard session launched by
// the server.
type Job struct {
	uuid string // unique id of the job
	path string // tensorboard log path
	pid  int    // pid of job (-1 => not running)
}

// newJob returns a new instance of a job with a UUID and path to a tensorboard
// log dir.
func newJob(path string) (*Job, error) {
	id, err := newUUID()
	if err != nil {
		return nil, err
	}
	return &Job{uuid: id, path: path, pid: -1}, nil
}

// startTensorboard kicks off a tensorboard instance at the next port in the
// port range.
func (j *Job) startTensorboard() error {
	return errors.New("startTensorboard - not implemented yet")
}

////////////////////////////////////////////////////////////////////////////////

// Jobs aggregates all pending and current jobs.
type Jobs struct {
	sync.RWMutex
	jobs []*Job // queue of current followed by pending jobs
}

// loadJobs loads the jobs saved to the local "db".
// TODO: Check PID of jobs
func (js *Jobs) loadJobs() error { return nil }

// saveJobs saves the current and pending jobs to the local "db".
func (js *Jobs) saveJobs() error { return nil }

// addJob() adds a new job with a specified path to the job list.
func (js *Jobs) addJob(path string) error {
	js.Lock()
	defer js.Unlock()

	j, err := newJob(path)
	if err != nil {
		return err
	}

	js.jobs = append(js.jobs, j)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("GET /\n")
	}
}

func newTensorboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("GET /tensorboard\n")
	}
}

func main() {
	http.HandleFunc("/", rootHandler())
	http.HandleFunc("/tensorboard", newTensorboardHandler())

	http.ListenAndServe(":8018", nil)
}

////////////////////////////////////////////////////////////////////////////////

func init() {
}

////////////////////////////////////////////////////////////////////////////////
