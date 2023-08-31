package main

import "fmt"

var CheckCacheMismatch = &Check{
	ID:  "cache-mismatch",
	Run: checkCacheMismatch,
}

func checkCacheMismatch(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	traceRecords := normalize(outputs.trace[len(outputs.trace)-1].Answers, config.RecordType)
	resolverRecords := normalize(outputs.resolverNoRecurse.Answers, config.RecordType)
	if outputs.resolverNoRecurse.Status != "NOERROR" {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Resolver doesn't have any records cached"),
		}, nil
	}

	onlyResolver, onlyAuthoritative := diff(resolverRecords, traceRecords)

	if len(onlyResolver) == 0 {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Cached records match authoritative records"),
		}, nil
	}

	return &CheckResult{
		Status: false,
		Message: fmt.Sprintf(`Cached records don't match authoritative records:
Only in resolver:
%s
Only in authoritative:
%s`, showRecords(onlyResolver), showRecords(onlyAuthoritative)),
	}, nil
}
