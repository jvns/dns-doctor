package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCacheMismatchFail(t *testing.T) {
	domain := "ring269.messwithdns.com"
	got := testRunCheck(t, CheckCacheMismatch, domain, "A")
	autogold.Expect(&CheckResult{Status: false, Message: `Cached records don't match authoritative records:
Only in resolver:
ring269.messwithdns.com. 267 A 1.2.3.4

Only in authoritative:
ring269.messwithdns.com. 300 A 5.6.7.8
`}).Equal(t, got)
}

func TestCacheMismatchSuccess(t *testing.T) {
	got := testRunCheck(t, CheckCacheMismatch, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Cached records match authoritative records"}).Equal(t, got)
}
