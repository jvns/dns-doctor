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
	cname             string
	cnameTrace        []DNSResponse
	cnameNoRecurse    *DNSResponse
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
	trace := runDigTrace(config)
	cname := getCNAME(trace)
	outputs := &DigOutputs{
		trace:             trace,
		cname:             cname,
		resolverNoRecurse: runDigNorecurse(config),
		cnameTrace:        runDigCNAMETrace(config, cname),
		cnameNoRecurse:    runDigCNAMENorecurse(config, cname),
	}

	runCheck(CheckNoRecord, config, outputs)
	runCheck(CheckCacheMismatch, config, outputs)
	runCheck(CheckBadCNAME, config, outputs)
}

func run(cmd *exec.Cmd) string {
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error running `%v`: %v", cmd, err)
		os.Exit(1)
	}
	return string(stdout)
}

func runDigTrace(config *Config) []DNSResponse {
	cmd := exec.Command("dig", "+trace", "+all", config.RecordType, config.Domain)
	stdout := run(cmd)
	writeFile(config.Domain+"_"+config.RecordType+"_trace.dig", stdout)
	trace := parseDigTraceOutput(stdout)
	if len(trace) == 0 {
		fmt.Println("No trace output found")
		os.Exit(1)
	}
	return trace
}

func runDigNorecurse(config *Config) DNSResponse {
	cmd := exec.Command("dig", "+all", "+norecurse", config.RecordType, config.Domain)
	stdout := run(cmd)
	writeFile(config.Domain+"_"+config.RecordType+"_norecurse.dig", string(stdout))
	return parseDigOutput(stdout)
}

func runDig(config *Config) DNSResponse {
	cmd := exec.Command("dig", "+all", config.RecordType, config.Domain)
	stdout := run(cmd)
	writeFile(config.Domain+"_"+config.RecordType+".dig", string(stdout))
	return parseDigOutput(stdout)
}

func getCNAME(trace []DNSResponse) string {
	records := trace[len(trace)-1].Answers
	for _, record := range records {
		if record.Type == "CNAME" {
			return record.Data
		}
	}
	return ""
}

func runDigCNAMETrace(rootConfig *Config, cname string) []DNSResponse {
	if cname == "" {
		return nil
	}
	return runDigTrace(&Config{
		RecordType: rootConfig.RecordType,
		Domain:     cname,
	})
}

func runDigCNAMENorecurse(rootConfig *Config, cname string) *DNSResponse {
	if cname == "" {
		return nil
	}
	resp := runDigNorecurse(&Config{
		RecordType: rootConfig.RecordType,
		Domain:     cname,
	})
	return &resp
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
