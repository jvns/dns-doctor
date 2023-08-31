package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCheckNoRecordFail(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "exampleffff.com", "A")
	autogold.Expect(&CheckResult{Status: false, Message: "No record found"}).Equal(t, got)
}

func TestCheckNoRecordSucceed(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Found record: '93.184.216.34'"}).Equal(t, got)
}
