package main

import (
	"fmt"
	"sort"
)

type Check struct {
	ID          string
	Description string
	Run         func(config *Config, outputs *DigOutputs) (*CheckResult, error)
}

type CheckResult struct {
	Status  bool
	Message string
}

func runCheck(check *Check, config *Config, outputs *DigOutputs) {
	// color the output based on the result of the check
	result, err := check.Run(config, outputs)
	if err != nil {
		fmt.Printf("Error running check '%s': %s\n", check.ID, err)
		return
	}
	if result.Status {
		fmt.Println("\033[32mSUCCESS:\033[0m", check.ID)
	} else {
		fmt.Println("\033[31mFAILURE:\033[0m", check.ID)
		fmt.Printf("  \033[31mFAILURE\033[0m: %s\n", result.Message)
	}
}

func filterRecords(records []Record, typ string) []Record {
	filtered := []Record{}
	for _, record := range records {
		if record.Type == typ || record.Type == "CNAME" {
			filtered = append(filtered, record)
		}
	}
	return filtered
}

func sortRecords(records []Record) []Record {
	// copy records
	sorted := []Record{}
	for _, record := range records {
		sorted = append(sorted, record)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Data < sorted[j].Data
	})
	return sorted
}

func normalize(records []Record, typ string) []Record {
	records = filterRecords(records, typ)
	return sortRecords(records)
}

func containsRecord(records []Record, record Record) bool {
	for _, r := range records {
		if r.Data == record.Data && r.Type == record.Type && r.Name == record.Name && r.Class == record.Class {
			return true
		}
	}
	return false
}

func isSubset(resolverRecords []Record, authoritativeRecords []Record) bool {
	for _, record := range resolverRecords {
		if !containsRecord(authoritativeRecords, record) {
			return false
		}
	}
	return true
}

func diff(resolverRecords []Record, authoritativeRecords []Record) ([]Record, []Record) {
	plus := []Record{}
	minus := []Record{}
	for _, record := range resolverRecords {
		if !containsRecord(authoritativeRecords, record) {
			plus = append(plus, record)
		}
	}
	for _, record := range authoritativeRecords {
		if !containsRecord(resolverRecords, record) {
			minus = append(minus, record)
		}
	}
	/* only resolver, only authortative */
	return plus, minus
}

func showRecords(records []Record) string {
	str := ""
	for _, record := range records {
		str += record.String() + "\n"
	}
	return str
}
