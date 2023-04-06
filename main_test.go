package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestJobRunner(t *testing.T) {
	// Create a mock server for step 1
	step1Server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send a success response with a message indicating the action has started
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "started"}`))
	}))
	defer step1Server.Close()

	// Create a mock server for step 2
	step2Server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send a success response with a message indicating the action is finished
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "finished", "response": "Job completed"}`))
	}))
	defer step2Server.Close()

	// Create a job with the mock servers as Step1URL and Step2URL
	job := Job{
		ID:       1,
		Step1URL: step1Server.URL,
		Step2URL: step2Server.URL,
	}

	// Create a wait group to wait for all jobs to complete
	var wg sync.WaitGroup

	// Start the job runner
	go func() {
		jobRunner()
	}()

	// Add the job to the job queue
	jobQueue <- &job

	// Wait for the job to complete
	wg.Add(1)
	wg.Wait()

	// Assert that the job is completed
	if !job.Completed {
		t.Errorf("Job not completed, expected: true, got: false")
	}
}
