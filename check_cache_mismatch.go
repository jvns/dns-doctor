/*
### **`cache-mismatch`**

#### How `cache-mismatch` is implemented:

1. Look up the record with the authoritative nameserver (the equivalent of `dig +trace some.domain.com`)
2. Look up the record with the local resolver (the equivalent of `dig some.domain.com`)
3. If the record the local resolver returns is outdated (if it's not), fail this check

It also runs the same check with a few popular resolvers (`8.8.8.8`, `1.1.1.1`)


*/

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
