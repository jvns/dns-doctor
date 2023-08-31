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
	got := testRunCheck(t, "no_record_fail.txt", "resolver_no_record_fail.txt", "A", "exampleffff.com")
	autogold.Expect(&CheckResult{Status: false, Message: "No record found"}).Equal(t, got)
}

func TestCheckNoRecordSucceed(t *testing.T) {
	got := testRunCheck(t, "dig_trace_example_com.txt", "example_com_norecurse.txt", "A", "exampleffff.com")
	autogold.Expect(&CheckResult{Status: true, Message: "Found record: '93.184.216.34'"}).Equal(t, got)
}
