package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCheckBadCNAMENone(t *testing.T) {
	got := testRunCheck(t, CheckBadCNAME, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "No CNAME found"}).Equal(t, got)
}

func TestCheckBadCNAMEOK(t *testing.T) {
	got := testRunCheckCNAME(t, CheckBadCNAME, "www.github.com", "A", "github.com.")
	autogold.Expect(&CheckResult{Status: true, Message: "Record found for CNAME domain 'github.com.': 140.82.114.4"}).Equal(t, got)
}

func TestCheckBadCNAMEFail(t *testing.T) {
	got := testRunCheckCNAME(t, CheckBadCNAME, "bad-cname.ring269.messwithdns.com", "A", "examplefffffff.com.")
	autogold.Expect(&CheckResult{Status: false, Message: "No record found for CNAME domain 'examplefffffff.com.'"}).Equal(t, got)
}
