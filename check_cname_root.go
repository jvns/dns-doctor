package main

import (
	"fmt"
	"strings"
)

var CheckCnameRoot = &Check{
	ID:  "cname-root",
	Run: checkCnameRoot,
}

func normalizeDomain(domain string) string {
	if domain[len(domain)-1] != '.' {
		return domain + "."
	}
	return strings.ToLower(domain)
}

func checkCnameRoot(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	domain := outputs.trace[2].Authorities[0].Name
	if normalizeDomain(domain) != normalizeDomain(config.Domain) {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Skipping: this is a subdomain"),
		}, nil
	}

	answers := outputs.trace[len(outputs.trace)-1].Answers
	for _, answer := range answers {
		if answer.Type == "CNAME" {
			return &CheckResult{
				Status:  false,
				Message: fmt.Sprintf("CNAME at root points to '%s'", answer.Data),
			}, nil
		}
	}
	return &CheckResult{
		Status:  true,
		Message: fmt.Sprintf("No CNAME at root"),
	}, nil
}
