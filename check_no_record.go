package main

import "fmt"

var CheckNoRecord = &Check{
	ID:  "no-record",
	Run: checkNoRecord,
}

func checkNoRecord(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	last_response := outputs.trace[len(outputs.trace)-1]
	if len(outputs.trace) == 3 {
		return &CheckResult{
			Status:  false,
			Message: fmt.Sprintf("This domain is not registered"),
		}, nil
	}
	if len(last_response.Answers) == 0 {
		nameserver := last_response.ServerName
		return &CheckResult{
			Status:  false,
			Message: fmt.Sprintf("No record found, using nameserver '%s'", nameserver),
		}, nil
	} else {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Found record: '%v'", last_response.Answers[0].Data),
		}, nil
	}
}
