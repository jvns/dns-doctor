package main

import "fmt"

func checkNoRecord(config *Config, outputs *DigOutputs) {
	fmt.Println("Checking for no record...")
	last_response := outputs.trace[len(outputs.trace)-1]
	if len(last_response.Answers) == 0 {
		fmt.Println("  FAILED: No record found")
	} else {
		fmt.Println(fmt.Sprintf("  PASSED: Found record for %s", config.Domain))
	}
}
