package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCheckCnameRoot(t *testing.T) {
	got := testRunCheck(t, CheckCnameRoot, "cnameroot.com", "A")
	autogold.Expect(&CheckResult{Message: "CNAME at root points to 'examplecat.com.'"}).Equal(t, got)

	got = testRunCheck(t, CheckCnameRoot, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "No CNAME at root"}).Equal(t, got)

	got = testRunCheck(t, CheckCnameRoot, "www.github.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Skipping: this is a subdomain"}).Equal(t, got)
}
