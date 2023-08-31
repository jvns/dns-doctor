package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCheckNoRecordNotRegistered(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "exampleffff.com", "A")
	autogold.Expect(&CheckResult{Message: "This domain is not registered"}).Equal(t, got)
}

func TestCheckNoRecordFail(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "nonexistent.example.com", "A")
	autogold.Expect(&CheckResult{Message: "No record found, using nameserver 'a.iana-servers.net'"}).Equal(t, got)
}

func TestCheckNoRecordSucceed(t *testing.T) {
	got := testRunCheck(t, CheckNoRecord, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Found record: '93.184.216.34'"}).Equal(t, got)
}
