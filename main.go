package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type Config struct {
	RecordType string
	Domain     string
	Resolver   string
}

type DigOutputs struct {
	trace             []DNSResponse
	resolver          DNSResponse
	resolverNoRecurse DNSResponse
}

func main() {
	// dns doctor usage: dnsdoctor <record-type> <domain-name> [flags]
	// example: dnsdoctor A google.com
	// example: dnsdoctor google.com

	cmd := &cobra.Command{
		Use:   "dnsdoctor [record-type] domain-name",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			/* require at least 1 arg */
			if len(args) < 1 {
				cmd.Help()
				return
			}
			config := &Config{}
			if len(args) == 1 {
				config.RecordType = "A"
				config.Domain = args[0]
			} else {
				config.RecordType = args[0]
				config.Domain = args[1]
			}
			doctor(config)
		},
	}
	// add 2 arguments, one for record type and one for domain name
	// first one is optional, second one is required
	cmd.Execute()
}

func doctor(config *Config) {
	outputs := &DigOutputs{
		trace:             runDigTrace(config),
		resolver:          runDig(config),
		resolverNoRecurse: runDig(config),
	}

	check_no_record(config, outputs)
}

func check_no_record(config *Config, outputs *DigOutputs) {
	fmt.Println("Checking for no record...")
	last_response := outputs.trace[len(outputs.trace)-1]
	if len(last_response.Answers) == 0 {
		fmt.Println("  FAILED: No record found")
	} else {
		fmt.Println(fmt.Sprintf("  PASSED: Found record for %s", config.Domain))
	}
}

func runDigTrace(config *Config) []DNSResponse {
	// run dig +all +trace {record_type} {domain}
	cmd := exec.Command("dig", "+trace", "+all", config.RecordType, config.Domain)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error running `%v`: %v", cmd, err)
		os.Exit(1)
	}
	trace := parseDigTraceOutput(string(stdout))
	if len(trace) == 0 {
		fmt.Println("No trace output found")
		os.Exit(1)
	}
	return trace
}

func runDig(config *Config) DNSResponse {
	// run dig +all +trace {record_type} {domain}
	cmd := exec.Command("dig", "+all", config.RecordType, config.Domain)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error running `%v`: %v", cmd, err)
		os.Exit(1)
	}
	return parseDigOutput(string(stdout))
}
