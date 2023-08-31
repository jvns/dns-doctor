package main

import "fmt"

type Check struct {
	Name        string
	Description string
	Run         func(config *Config, outputs *DigOutputs) *CheckResult
}

type CheckResult struct {
	Status  bool
	Message string
}

func checkNoRecord(config *Config, outputs *DigOutputs) *CheckResult {
	fmt.Println("Checking for no record...")
	last_response := outputs.trace[len(outputs.trace)-1]
	if len(last_response.Answers) == 0 {
		return &CheckResult{
			Status:  false,
			Message: "No record found",
		}
	} else {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Found record: %v", last_response.Answers[0]),
		}
	}
}
