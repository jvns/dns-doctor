package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	Name  string
	TTL   int
	Class string
	Type  string
	Data  string
}

func (r *Record) String() string {
	return fmt.Sprintf("%s %d %s %s", r.Name, r.TTL, r.Type, r.Data)
}

type Question struct {
	Name  string
	Type  string
	Class string
}

type DNSResponse struct {
	Status      string
	ServerIP    string
	ServerName  string
	Question    Question
	Answers     []Record
	Authorities []Record
	Additionals []Record
}

type TraceOutput struct {
	Server  string
	IP      string
	Records []Record
}

func parseQuestion(line string) Question {
	fields := strings.Fields(line[1:])
	if len(fields) < 3 {
		panic(fmt.Sprintf("Invalid record: %s", line))
	}
	return Question{
		Name:  fields[0],
		Class: fields[1],
		Type:  fields[2],
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
		Name:  fields[0],
		TTL:   ttl,
		Class: fields[2],
		Type:  fields[3],
		Data:  strings.Join(fields[4:], " "),
	}
}

func parseDigOutput(output string) DNSResponse {
	lines := strings.Split(output, "\n")
	part := "start"
	resp := DNSResponse{
		Status:      "",
		Question:    Question{},
		Answers:     make([]Record, 0),
		Authorities: make([]Record, 0),
		Additionals: make([]Record, 0),
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
					resp.Status = fields[i+1][:len(fields[i+1])-1]
					break
				}
			}
		} else if strings.Contains(line, "SERVER:") {

			// SERVER: 192.48.79.30#53(j.gtld-servers.net) (UDP)
			regex := regexp.MustCompile(`SERVER: (.+)#(\d+)\(([\w:\.\-]+)\)`)
			matches := regex.FindStringSubmatch(line)
			if len(matches) != 4 {
				panic(fmt.Sprintf("Invalid server line: %s", line))
			}
			resp.ServerIP = fmt.Sprintf("%s:%s", matches[1], matches[2])
			resp.ServerName = matches[3]
		} else if part == "question" {
			resp.Question = parseQuestion(line)
		} else if part == "answer" {
			resp.Answers = append(resp.Answers, parseRecord(line))
		} else if part == "authority" {
			resp.Authorities = append(resp.Authorities, parseRecord(line))
		} else if part == "additional" {
			resp.Additionals = append(resp.Additionals, parseRecord(line))
		}

	}
	return resp
}

func parseDigTraceOutput(output string) []DNSResponse {
	parts := strings.Split(output, "Got answer")[1:]
	responses := make([]DNSResponse, 0)
	for _, part := range parts {
		responses = append(responses, parseDigOutput(part))
	}
	return responses
}
