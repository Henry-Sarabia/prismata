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
		fail bool
	}{
		{
			"Pass: replay 1",
			testFile1,
			Replay{Code: "ib0Qt-pp8PL", EndTimeUnix: 1521092570.323161},
			false,
		},
		{
			"Pass: replay 2",
			testFile2,
			Replay{Code: "VyrET-IGxyL", EndTimeUnix: 1532955756.487255},
			false,
		},
		{
			"Pass: replay 3",
			testFile3,
			Replay{Code: "yjUKQ-HzFRz", EndTimeUnix: 1533169252.757027},
			false,
		},
		{
			"Error: empty",
			testFileEmpty,
			Replay{Code: "fake-code", EndTimeUnix: 123.456},
			true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			r, err := Decode(f)
			assertError(t, err, tt.fail)

			if tt.fail {
				return
			}

			if r.Code != tt.exp.Code {
				t.Errorf("got: <%v>, want: <%v>", r.Code, tt.exp.Code)
			}

			if r.EndTimeUnix != tt.exp.EndTimeUnix {
				t.Errorf("got: <%v>, want: <%v>", r.EndTimeUnix, tt.exp.EndTimeUnix)
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
			"Pass: start > end",
			now,
			now.Add(time.Second * 450),
			time.Second * 450,
		},
		{
			"Pass: end > start",
			now.Add(time.Minute * 10),
			now,
			time.Minute * 10,
		},
		{
			"Pass: start = end",
			now,
			now,
			time.Second * 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			d := duration(tt.start, tt.end)
			if d != tt.exp {
				t.Errorf("got: <%v>,  want: <%v>", d, tt.exp)
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
			"Pass: replay 1",
			Replay{StartTimeUnix: 1521091944.949862, EndTimeUnix: 1521092570.323161},
			(time.Minute * 10) + (time.Second * 26),
			false,
		},
		{
			"Pass: replay 2",
			Replay{StartTimeUnix: 1532955317.245488, EndTimeUnix: 1532955756.487255},
			(time.Minute * 7) + (time.Second * 19),
			false,
		},
		{
			"Pass: replay 3",
			Replay{StartTimeUnix: 1533168612.190871, EndTimeUnix: 1533169252.757027},
			(time.Minute * 10) + (time.Second * 40),
			false,
		},
		{
			"Error: zero start time",
			Replay{StartTimeUnix: 0, EndTimeUnix: 1533169252},
			0,
			true,
		},
		{
			"Error: zero end time",
			Replay{StartTimeUnix: 1533169252, EndTimeUnix: 0},
			0,
			true,
		},
		{
			"Error: zero start and end time",
			Replay{StartTimeUnix: 0, EndTimeUnix: 0},
			0,
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
				t.Errorf("got: <%v>, want: <%v>", d, tt.exp)
			}
		})
	}
}

func TestStartTime(t *testing.T) {
	var cases = []struct {
		name string
		r    Replay
		exp  time.Time
	}{
		{
			"Pass: replay 1",
			Replay{StartTimeUnix: 1521091944.949862},
			time.Unix(int64(1521091944), 0),
		},
		{
			"Pass: replay 2",
			Replay{StartTimeUnix: 1532955317.245488},
			time.Unix(int64(1532955317), 0),
		},
		{
			"Pass: replay 3",
			Replay{StartTimeUnix: 1533168612.190871},
			time.Unix(int64(1533168612), 0),
		},
		{
			"Error: zero start time",
			Replay{StartTimeUnix: 0},
			time.Unix(0, 0),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			st := tt.r.StartTime()
			if st != tt.exp {
				t.Errorf("got: <%v>, want: <%v>", st, tt.exp)
			}
		})
	}
}

func TestEndTime(t *testing.T) {
	var cases = []struct {
		name string
		r    Replay
		exp  time.Time
	}{
		{
			"Pass: replay 1",
			Replay{EndTimeUnix: 1521092570.323161},
			time.Unix(int64(1521092570), 0),
		},
		{
			"Pass: replay 2",
			Replay{EndTimeUnix: 1532955756.487255},
			time.Unix(int64(1532955756), 0),
		},
		{
			"Pass: replay 3",
			Replay{EndTimeUnix: 1533169252.757027},
			time.Unix(int64(1533169252), 0),
		},
		{
			"Error: zero end time",
			Replay{EndTimeUnix: 0},
			time.Unix(0, 0),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			st := tt.r.EndTime()
			if st != tt.exp {
				t.Errorf("got: <%v>, want: <%v>", st, tt.exp)
			}
		})
	}
}
