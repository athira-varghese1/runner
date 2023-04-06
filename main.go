package main

import (
	"fmt"
	"net/http"
	"time"
)

type Job struct {
	ID       int
	Step1URL string
	Step2URL string
}

func main() {
	job1 := Job{
		ID:       1,
		Step1URL: "https://remote-service.com/trigger-action",
		Step2URL: "https://remote-service.com/check-action-status",
	}

	job2 := Job{
		ID:       2,
		Step1URL: "https://remote-service.com/trigger-action",
		Step2URL: "https://remote-service.com/check-action-status",
	}

	job3 := Job{
		ID:       3,
		Step1URL: "https://remote-service.com/trigger-action",
		Step2URL: "https://remote-service.com/check-action-status",
	}

	jobList := []Job{job1, job2, job3}

	for _, job := range jobList {
		go executeJob(job)
	}

	// Wait for all jobs to finish
	time.Sleep(time.Second * 10)
}

func executeJob(job Job) {
	fmt.Printf("Starting Job ID: %d\n", job.ID)

	// Step 1: Send HTTP request to trigger action
	resp, err := http.Get(job.Step1URL)
	if err != nil {
		fmt.Printf("Job ID: %d - Step 1 failed: %s\n", job.ID, err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Job ID: %d - Step 1 completed\n", job.ID)

	// Step 2: Send HTTP request to check action status
	resp, err = http.Get(job.Step2URL)
	if err != nil {
		fmt.Printf("Job ID: %d - Step 2 failed: %s\n", job.ID, err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Job ID: %d - Step 2 completed\n", job.ID)

	// Step 3: Get response and print if action is finished
	// Assuming the response is a plain text string indicating action status
	// You may need to modify this part based on the actual response format
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Job ID: %d - Action finished: %s\n", job.ID, resp.Status)
	} else {
		fmt.Printf("Job ID: %d - Action not finished: %s\n", job.ID, resp.Status)
	}

	fmt.Printf("Job ID: %d - Job completed\n", job.ID)
}
