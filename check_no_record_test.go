package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func testRunCheck(t *testing.T, traceFilename string, resolverFilename string, recordType string, domain string) *CheckResult {
	trace := parseDigTraceOutput(readFile(t, "testdata/"+traceFilename))
	resolver := parseDigOutput(readFile(t, "testdata/"+resolverFilename))
	config := &Config{
		RecordType: recordType,
		Domain:     domain,
	}
	outputs := &DigOutputs{
		trace:             trace,
		resolverNoRecurse: resolver,
	}
	got, err := checkNoRecord(config, outputs)
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func TestCheckNoRecordFail(t *testing.T) {
	got := testRunCheck(t, "exampleffff.com_A_trace.dig", "exampleffff.com_A_norecurse.dig", "A", "exampleffff.com")
	autogold.Expect(&CheckResult{Status: false, Message: "No record found"}).Equal(t, got)
}

func TestCheckNoRecordSucceed(t *testing.T) {
	got := testRunCheck(t, "example.com_A_trace.dig", "example.com_A_norecurse.dig", "A", "example.com")
	autogold.Expect(&CheckResult{Status: true, Message: "Found record: '93.184.216.34'"}).Equal(t, got)
}
