package main

import (
	"fmt"
	"testing"

	"github.com/hexops/autogold/v2"
)

func testRunCheck(t *testing.T, check *Check, domain string, recordType string) *CheckResult {
	// trace := fmt.Sprintf("%s_A_trace.dig", domain)
	//resolver := fmt.Sprintf("%s_A_norecurse.dig", domain)
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

func TestCheckNoRecordFail(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "exampleffff.com", "A")
	autogold.Expect(&CheckResult{Status: false, Message: "No record found"}).Equal(t, got)
}

func TestCheckNoRecordSucceed(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Found record: '93.184.216.34'"}).Equal(t, got)
}
