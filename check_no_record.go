package main

import "fmt"

var CheckNoRecord = &Check{
	ID:  "check-no-record",
	Run: checkNoRecord,
}

func checkNoRecord(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	last_response := outputs.trace[len(outputs.trace)-1]
	if len(last_response.Answers) == 0 {
		return &CheckResult{
			Status:  false,
			Message: "No record found",
		}, nil
	} else {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Found record: '%v'", last_response.Answers[0].Data),
		}, nil
	}
}
