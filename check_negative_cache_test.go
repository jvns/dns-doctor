package main

import (
	"testing"

	"github.com/hexops/autogold/v2"
)

func TestCheckNegativeCacheFail(t *testing.T) {
	got := testRunCheck(t, CheckNegativeCache, "neg-cache.ring269.messwithdns.com", "A")
	autogold.Expect(&CheckResult{Message: "Resolver's cached result is empty, but authoritative answer is '1.2.3.4'"}).Equal(t, got)
}

func TestCheckNegativeCacheSuccess(t *testing.T) {
	got := testRunCheck(t, CheckNegativeCache, "example.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Resolver's cached result is nonempty (A 93.184.216.34)"}).Equal(t, got)
}

func TestCheckNegativeCacheSuccess2(t *testing.T) {
	got := testRunCheck(t, CheckNegativeCache, "examplecat.com", "A")
	autogold.Expect(&CheckResult{Status: true, Message: "Resolver doesn't have any records cached"}).Equal(t, got)
}
