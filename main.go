package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Record struct {
	name  string
	ttl   int
	class string
	typ   string
	data  string
}

type Question struct {
	name  string
	typ   string
	class string
}

type DNSResponse struct {
	status     string
	question   Question
	answer     []Record
	authority  []Record
	additional []Record
}

type TraceOutput struct {
	server  string
	ip      string
	records []Record
}

func main() {
}

func parseQuestion(line string) Question {
	fields := strings.Fields(line[1:])
	if len(fields) < 3 {
		panic(fmt.Sprintf("Invalid record: %s", line))
	}
	return Question{
		name:  fields[0],
		class: fields[1],
		typ:   fields[2],
	}
}
func parseRecord(line string) Record {
	fields := strings.Fields(line)
	if len(fields) < 5 {
		panic(fmt.Sprintf("Invalid record: %s", line))
	}
	ttl, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(fmt.Sprintf("Invalid ttl: %s", fields[1]))
	}
	return Record{
		name:  fields[0],
		ttl:   ttl,
		class: fields[2],
		typ:   fields[3],
		data:  strings.Join(fields[4:], " "),
	}
}

func parseDigOutput(output string) DNSResponse {
	lines := strings.Split(output, "\n")
	part := "start"
	resp := DNSResponse{
		status:     "",
		question:   Question{},
		answer:     make([]Record, 0),
		authority:  make([]Record, 0),
		additional: make([]Record, 0),
	}

	for _, line := range lines {
		if strings.Contains(line, "ANSWER SECTION") {
			part = "answer"
		} else if strings.Contains(line, "AUTHORITY SECTION") {
			part = "authority"
		} else if len(line) == 0 {
			part = "start"
		} else if strings.Contains(line, "ADDITIONAL SECTION") {
			part = "additional"
		} else if strings.Contains(line, "QUESTION SECTION") {
			part = "question"
		} else if strings.Contains(line, "->>HEADER<<-") {
			fields := strings.Fields(line)
			for i, field := range fields {
				if field == "status:" {
					resp.status = fields[i+1][:len(fields[i+1])-1]
					break
				}
			}
		} else if part == "question" {
			resp.question = parseQuestion(line)
		} else if part == "answer" {
			resp.answer = append(resp.answer, parseRecord(line))
		} else if part == "authority" {
			resp.authority = append(resp.authority, parseRecord(line))
		} else if part == "additional" {
			resp.additional = append(resp.additional, parseRecord(line))
		}

	}
	return resp
}
