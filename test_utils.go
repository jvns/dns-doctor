package main

import (
	"fmt"
	"os"
	"testing"
)

func readFile(t *testing.T, path string) string {
	t.Helper()
	output, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(output)
}

func testRunCheck(t *testing.T, check *Check, domain string, recordType string) *CheckResult {
	trace := parseDigTraceOutput(readFile(t, fmt.Sprintf("testdata/%s_%s_trace.dig", domain, recordType)))
	resolver := parseDigOutput(readFile(t, fmt.Sprintf("testdata/%s_%s_norecurse.dig", domain, recordType)))
	config := &Config{
		RecordType: recordType,
		Domain:     domain,
	}
	outputs := &DigOutputs{
		trace:             trace,
		resolverNoRecurse: resolver,
	}
	got, err := check.Run(config, outputs)
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func testRunCheckCNAME(t *testing.T, check *Check, domain string, recordType string, cname string) *CheckResult {
	trace := parseDigTraceOutput(readFile(t, fmt.Sprintf("testdata/%s_%s_trace.dig", domain, recordType)))
	resolver := parseDigOutput(readFile(t, fmt.Sprintf("testdata/%s_%s_norecurse.dig", domain, recordType)))
	traceCNAME := parseDigTraceOutput(readFile(t, fmt.Sprintf("testdata/%s_%s_trace.dig", cname, recordType)))
	resolverCNAME := parseDigOutput(readFile(t, fmt.Sprintf("testdata/%s_%s_norecurse.dig", cname, recordType)))
	config := &Config{
		RecordType: recordType,
		Domain:     domain,
	}
	outputs := &DigOutputs{
		trace:             trace,
		resolverNoRecurse: resolver,
		cname:             cname,
		cnameTrace:        traceCNAME,
		cnameNoRecurse:    &resolverCNAME,
	}
	got, err := check.Run(config, outputs)
	if err != nil {
		t.Fatal(err)
	}
	return got
}
