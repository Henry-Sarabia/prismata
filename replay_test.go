package main

import (
	"os"
	"testing"
	"time"
)

var (
	testReplay1 = "replay1.json"
	testReplay2 = "replay2.json"
	testReplay3 = "replay3.json"
	testEmpty   = "empty.json"
)

func TestParse(t *testing.T) {
	var cases = []struct {
		name string
		file string
		exp  Replay
		err  bool
	}{
		{
			"Success: replay 1",
			testReplay1,
			Replay{Code: "ib0Qt-pp8PL", EndTime: 1521092570.323161},
			false,
		},
		{
			"Success: replay 2",
			testReplay2,
			Replay{Code: "VyrET-IGxyL", EndTime: 1532955756.487255},
			false,
		},
		{
			"Success: replay 3",
			testReplay3,
			Replay{Code: "yjUKQ-HzFRz", EndTime: 1533169252.757027},
			false,
		},
		{
			"Error: empty",
			testEmpty,
			Replay{Code: "fake-code", EndTime: 123.456},
			true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			r, err := parse(f)
			if tt.err && err == nil {
				t.Fatalf("got: <nil>, want: <error>")
			}

			if !tt.err && err != nil {
				t.Fatalf("got: <%v>, want: <nil>", err)
			}

			if tt.err {
				return
			}

			if r.Code != tt.exp.Code {
				t.Errorf("got: <%v>, want: <%v>", r.Code, tt.exp.Code)
			}

			if r.EndTime != tt.exp.EndTime {
				t.Errorf("got: <%v>, want <%v>", r.EndTime, tt.exp.EndTime)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	now := time.Now()
	var cases = []struct {
		name  string
		start time.Time
		end   time.Time
		exp   time.Duration
	}{
		{
			"Success: start > end",
			now,
			now.Add(time.Second * 450),
			time.Second * 450,
		},
		{
			"Success: end > start",
			now.Add(time.Minute * 10),
			now,
			time.Minute * 10,
		},
		{
			"Success: start = end",
			now,
			now,
			time.Second * 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			d := duration(tt.start, tt.end)
			if d != tt.exp {
				t.Errorf("got: <%v>,  want <%v>", d, tt.exp)
			}
		})
	}
}

// func TestReplayDuration(t *testing.T) {

// }
