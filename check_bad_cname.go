package main

import "fmt"

var CheckBadCNAME = &Check{
	ID:  "bad-cname",
	Run: checkBadCNAME,
}

func checkBadCNAME(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	if outputs.cnameTrace == nil {
		return &CheckResult{
			Status:  true,
			Message: "No CNAME found",
		}, nil
	}
	records := outputs.cnameTrace[len(outputs.cnameTrace)-1].Answers
	if len(records) == 0 {
		return &CheckResult{
			Status:  false,
			Message: fmt.Sprintf("No record found for CNAME domain '%s'", outputs.cname),
		}, nil
	}
	return &CheckResult{
		Status:  true,
		Message: fmt.Sprintf("Record found for CNAME domain '%s': %s", outputs.cname, records[0].Data),
	}, nil
}
