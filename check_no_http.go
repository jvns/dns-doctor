package main

import (
	"fmt"
	"net"
	"time"
)

var CheckNoHTTP = &Check{
	ID:  "no-http",
	Run: checkNoHTTP,
}

func checkNoHTTP(config *Config, outputs *DigOutputs) (*CheckResult, error) {
	if config.RecordType != "A" {
		return &CheckResult{
			Status:  true,
			Message: fmt.Sprintf("Skipping check for non-A record"),
		}, nil
	}
	answers := outputs.trace[len(outputs.trace)-1].Answers
	for _, answer := range answers {
		if answer.Type == "A" {
			result := connect(answer.Data, 80)
			if !result {
				return &CheckResult{
					Status:  false,
					Message: fmt.Sprintf("Failed to connect to %s:80", answer.Data),
				}, nil
			}
			result = connect(answer.Data, 443)
			if !result {
				return &CheckResult{
					Status:  false,
					Message: fmt.Sprintf("Failed to connect to %s:443", answer.Data),
				}, nil
			}
		}
	}
	return &CheckResult{
		Status:  true,
		Message: fmt.Sprintf("All A records have HTTP and HTTPS"),
	}, nil

}

func connect(ip string, port int) bool {
	dialer := net.Dialer{Timeout: 300 * time.Millisecond}
	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
