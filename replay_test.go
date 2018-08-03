package main

import (
	"os"
	"testing"
	"time"
)

var (
	testFile1     = "testdata/replay1.json"
	testFile2     = "testdata/replay2.json"
	testFile3     = "testdata/replay3.json"
	testFileEmpty = "testdata/empty.json"
)

// assertError fails the test if the given error does not match the expected
// outcome provided by the exp bool.
func assertError(t *testing.T, err error, exp bool) {
	if exp && err == nil {
		t.Fatalf("got: <%v>, want: <error>", err)
	}

	if !exp && err != nil {
		t.Fatalf("got: <%v>, want: <nil>", err)
	}
}

func TestParse(t *testing.T) {
	var cases = []struct {
		name string
		file string
		exp  Replay
		err  bool
	}{
		{
			"Success: replay 1",
			testFile1,
			Replay{Code: "ib0Qt-pp8PL", EndTime: 1521092570.323161},
			false,
		},
		{
			"Success: replay 2",
			testFile2,
			Replay{Code: "VyrET-IGxyL", EndTime: 1532955756.487255},
			false,
		},
		{
			"Success: replay 3",
			testFile3,
			Replay{Code: "yjUKQ-HzFRz", EndTime: 1533169252.757027},
			false,
		},
		{
			"Error: empty",
			testFileEmpty,
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

func TestReplayDuration(t *testing.T) {
	var cases = []struct {
		name string
		r    Replay
		exp  time.Duration
		fail bool
	}{
		{
			"Success: replay 1",
			Replay{StartTime: 1521091944.949862, EndTime: 1521092570.323161},
			(time.Minute * 10) + (time.Second * 26),
			false,
		},
		{
			"Success: replay 2",
			Replay{StartTime: 1532955317.245488, EndTime: 1532955756.487255},
			(time.Minute * 7) + (time.Second * 19),
			false,
		},
		{
			"Success: replay 3",
			Replay{StartTime: 1533168612.190871, EndTime: 1533169252.757027},
			(time.Minute * 10) + (time.Second * 40),
			false,
		},
		{
			"Failure: empty",
			Replay{StartTime: 0, EndTime: 0},
			(time.Second * 0),
			true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			d, err := tt.r.Duration()

			assertError(t, err, tt.fail)
			if tt.fail {
				return
			}

			if d != tt.exp {
				t.Errorf("got: <%v>, want <%v>", d, tt.exp)
			}
		})
	}
}
