# Job Runner

The job runner is a Go program that can execute predefined jobs step by step. It sends HTTP requests to trigger actions on remote services, waits for the actions to complete, and prints the response when the actions are finished.

## Prerequisites

- Go programming language (version 1.14+)
- Internet connection to send HTTP requests to remote services

## Installation
- Clone the repository or download the main.go file to your local machine.
- Navigate to the directory where the main.go file is located.

## Usage
1. Defined jobs in the main.go file by updating the Job struct with the necessary fields such as ID, Step1URL, Step2URL, etc. Each job represents a sequence of steps to be executed.
2. Modified jobRunner() function in the main.go file to include the logic for specific job processing. For example, sending HTTP requests, waiting for responses, and handling errors.
3. Run the main.go file using the go run command:
   `go run main.go`
4. The job runner will start processing jobs from the jobQueue channel and execute them step by step according to the defined logic.
5. You can monitor the progress of the jobs in the terminal output, which will display messages about job status and response when the actions are finished.

## Testing
1. You can run tests for the job runner using the go test command. 
2. The tests should be written following Go testing conventions, including using the testing package, writing test functions with names starting with Test, and using the t *testing.T parameter to perform assertions and report test results. 
3. Tests are provided in the main_test.go file. Run the file using the go test command: `go test `
  
