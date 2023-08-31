package main

import "fmt"

var CheckNegativeCache = &Check{
	ID:  "negative-cache",
	Run: checkNegativeCache,
}

func checkNegativeCache(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	if outputs.resolverNoRecurse.Status != "NOERROR" && outputs.resolverNoRecurse.Status != "NXDOMAIN" {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Resolver doesn't have any records cached"),
		}, nil
	}

	authRecords := outputs.trace[len(outputs.trace)-1].Answers
	resolverRecords := outputs.resolverNoRecurse.Answers
	if len(authRecords) != 0 && len(resolverRecords) == 0 {
		return &CheckResult{
			Status:  false,
			Message: fmt.Sprintf("Resolver's cached result is empty, but authoritative answer is '%s'", authRecords[0].Data),
		}, nil
	}

	return &CheckResult{
		Status:  true,
		Message: fmt.Sprintf("Resolver's cached result is nonempty (%s %s)", resolverRecords[0].Type, resolverRecords[0].Data),
	}, nil
}
