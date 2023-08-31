package main

import (
	"os"
	"testing"

	"github.com/hexops/autogold/v2"
)

// put autogold.Expect(nil).Equal(t, got) to make a test

func TestParseBasic(t *testing.T) {
	output, err := os.ReadFile("testdata/dig_example_com.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %s", err)
	}
	got := parseDigOutput(string(output))

	autogold.Expect(DNSResponse{
		status: "NOERROR", question: Question{
			name:  "example.com.",
			typ:   "A",
			class: "IN",
		},
		answer: []Record{{
			name:  "example.com.",
			ttl:   75789,
			class: "IN",
			typ:   "A",
			data:  "93.184.216.34",
		}},
		authority:  []Record{},
		additional: []Record{},
	}).Equal(t, got)
}

func TestParseRootNameserver(t *testing.T) {
	output, err := os.ReadFile("testdata/dig_example_com_authority.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %s", err)
	}
	got := parseDigOutput(string(output))

	autogold.Expect(DNSResponse{
		status: "NOERROR", question: Question{
			name:  "example.com.",
			typ:   "A",
			class: "IN",
		},
		answer: []Record{},
		authority: []Record{
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "a.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "b.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "c.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "d.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "e.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "f.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "g.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "h.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "i.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "j.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "k.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "l.gtld-servers.net.",
			},
			{
				name:  "com.",
				ttl:   172800,
				class: "IN",
				typ:   "NS",
				data:  "m.gtld-servers.net.",
			},
		},
		additional: []Record{
			{
				name:  "a.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.5.6.30",
			},
			{
				name:  "b.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.33.14.30",
			},
			{
				name:  "c.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.26.92.30",
			},
			{
				name:  "d.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.31.80.30",
			},
			{
				name:  "e.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.12.94.30",
			},
			{
				name:  "f.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.35.51.30",
			},
			{
				name:  "g.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.42.93.30",
			},
			{
				name:  "h.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.54.112.30",
			},
			{
				name:  "i.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.43.172.30",
			},
			{
				name:  "j.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.48.79.30",
			},
			{
				name:  "k.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.52.178.30",
			},
			{
				name:  "l.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.41.162.30",
			},
			{
				name:  "m.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "A",
				data:  "192.55.83.30",
			},
			{
				name:  "a.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:a83e::2:30",
			},
			{
				name:  "b.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:231d::2:30",
			},
			{
				name:  "c.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:83eb::30",
			},
			{
				name:  "d.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:500:856e::30",
			},
			{
				name:  "e.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:502:1ca1::30",
			},
			{
				name:  "f.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:d414::30",
			},
			{
				name:  "g.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:eea3::30",
			},
			{
				name:  "h.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:502:8cc::30",
			},
			{
				name:  "i.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:39c1::30",
			},
			{
				name:  "j.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:502:7094::30",
			},
			{
				name:  "k.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:503:d2d::30",
			},
			{
				name:  "l.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:500:d937::30",
			},
			{
				name:  "m.gtld-servers.net.",
				ttl:   172800,
				class: "IN",
				typ:   "AAAA",
				data:  "2001:501:b1f9::30",
			},
		},
	}).Equal(t, got)
}
