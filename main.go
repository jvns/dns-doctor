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
	trace []DNSResponse
	//resolver          DNSResponse
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
		trace: runDigTrace(config),
		//resolver:          runDig(config),
		resolverNoRecurse: runDigNorecurse(config),
	}

	runCheck(CheckNoRecord, config, outputs)
	runCheck(CheckCacheMismatch, config, outputs)
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
	writeFile(config.Domain+"_"+config.RecordType+"_trace.dig", string(stdout))
	if len(trace) == 0 {
		fmt.Println("No trace output found")
		os.Exit(1)
	}
	return trace
}

func runDigNorecurse(config *Config) DNSResponse {
	// run dig +all +trace {record_type} {domain}
	cmd := exec.Command("dig", "+all", "+norecurse", config.RecordType, config.Domain)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error running `%v`: %v", cmd, err)
		os.Exit(1)
	}
	writeFile(config.Domain+"_"+config.RecordType+"_norecurse.dig", string(stdout))
	return parseDigOutput(string(stdout))
}

func runDig(config *Config) DNSResponse {
	// run dig +all +trace {record_type} {domain}
	cmd := exec.Command("dig", "+all", config.RecordType, config.Domain)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error running `%v`: %v", cmd, err)
		os.Exit(1)
	}
	writeFile(config.Domain+"_"+config.RecordType+".dig", string(stdout))
	return parseDigOutput(string(stdout))
}

func writeFile(filename string, contents string) {
	// return
	f, err := os.Create("testdata/" + filename)
	if err != nil {
		fmt.Printf("error creating file %v: %v", filename, err)
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.WriteString(contents)
	if err != nil {
		fmt.Printf("error writing to file %v: %v", filename, err)
		os.Exit(1)
	}
}
