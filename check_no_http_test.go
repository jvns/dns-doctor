package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCheckNoHTTP(t *testing.T) {
	got := testRunCheck(t, CheckNoHTTP, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "All A records have HTTP and HTTPS"}).Equal(t, got)

	got = testRunCheck(t, CheckNoHTTP, "bad-ip.ring269.messwithdns.com", "A")
	autogold.Expect(&CheckResult{Status: false, Message: "Failed to connect to 203.0.113.0:80"}).Equal(t, got)
}
