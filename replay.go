package main

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// Replay contains information on the replay of a Prismata match.
type Replay struct {
	Code          string   `json:"code"`
	StartTimeUnix float64  `json:"startTime"`
	EndTimeUnix   float64  `json:"endTime"`
	Deck          Deck     `json:"deckInfo"`
	Players       []Player `json:"playerInfo"`
	// CommandInfo  CmdInfo    `json:"commandInfo"`
	TimeInfo     TimeInfo   `json:"timeInfo"`
	RatingInfo   RatingInfo `json:"ratingInfo"`
	Result       int        `json:"result"`
	VersionInfo  Version    `json:"versionInfo"`
	Seed         int        `json:"seed"`
	EndCondition int        `json:"endCondition"`
	Format       int        `json:"format"`
	RawHash      int        `json:"rawHash"`
}

// Unit represents a single deployable unit of play.
type Unit struct {
	Name    string `json:"name"`
	UIName  string `json:"UIName,omitempty"`
	BaseSet int    `json:"baseSet,omitempty"`
}

// CmdInfo contains information on the sequence of commands executed by players
// in a Prismata replay. Multiple commands are executed per turn.
type CmdInfo struct {
	CommandList        []Cmd     `json:"commandList"`
	CommandTimes       []float64 `json:"commandTimes"`
	CommandForced      []bool    `json:"commandForced"`
	TimesRemaining     []int     `json:"timesRemaining"`
	TimeBanksRemaining []float64 `json:"timeBanksRemaining"`
	MoveDurations      []float64 `json:"moveDurations"`
	ClicksPerTurn      []int     `json:"clicksPerTurn"`
}

// Cmd represents a single command executed by a player.
type Cmd struct {
	Type   string `json:"_type"`
	ID     int    `json:"_id"`
	Params Emote  `json:"_params,omitempty"`
}

// Emote contains information about a particular emote executed as a command
// by a player.
type Emote struct {
	MBackground    string `json:"mBackground"`
	MTextAnimation string `json:"mTextAnimation"`
	MColour        string `json:"mColour"`
	MTint          string `json:"mTint"`
	MFrame         string `json:"mFrame"`
}

// TimeInfo contains information about the time controls for a Prismata replay.
type TimeInfo struct {
	Correspondence         bool         `json:"correspondence"`
	PlayerCurrentTimeBanks []float64    `json:"playerCurrentTimeBanks"`
	PlayerTime             []PlayerTime `json:"playerTime"`
	GracePeriod            int          `json:"gracePeriod"`
	PlayerCurrentTimes     []int        `json:"playerCurrentTimes"`
	TurnNumber             int          `json:"turnNumber"`
	GraceCurrentTime       int          `json:"graceCurrentTime"`
	UseClocks              bool         `json:"useClocks"`
}

// PlayerTime contains information on a player's time bank.
type PlayerTime struct {
	BankDilution float64 `json:"bankDilution"`
	Initial      int     `json:"initial"`
	Bank         int     `json:"bank"`
	Increment    int     `json:"increment"`
}

// Version contains information on the version of Prismata being used
// for a particular replay.
type Version struct {
	ServerVersion  int      `json:"serverVersion"`
	PlayerVersions []string `json:"playerVersions"`
}

// RatingInfo contains information about the ratings of players in a Prismata replay.
type RatingInfo struct {
	InitialRatings []Rating    `json:"initialRatings"`
	FinalRatings   []Rating    `json:"finalRatings"`
	RatingChanges  [][]float64 `json:"ratingChanges"`
	ScoreChanges   []int       `json:"scoreChanges"`
}

// Rating contains information on a player's rating at the time of the Prismata replay.
type Rating struct {
	DisplayRating       float64 `json:"displayRating"`
	WinLastLast         bool    `json:"winLastLast"`
	DominionELO         int     `json:"dominionELO"`
	WinLast             bool    `json:"winLast"`
	PeakAdjustedShalevU float64 `json:"peakAdjustedShalevU"`
	ShalevV             float64 `json:"shalevV"`
	ShalevU             float64 `json:"shalevU"`
	Tier                int     `json:"tier"`
	CustomGamesPlayed   int     `json:"customGamesPlayed"`
	TierPercent         float64 `json:"tierPercent"`
	CasualGamesWon      int     `json:"casualGamesWon"`
	HStars              int     `json:"hStars"`
	Version             int     `json:"version"`
	Exp                 int     `json:"exp"`
	RatedGamesPlayed    int     `json:"ratedGamesPlayed"`
	BotGamesPlayed      int     `json:"botGamesPlayed"`
}

// Decode reads from the provided reader and decodes the JSON into a Replay.
func Decode(r io.Reader) (*Replay, error) {
	rep := &Replay{}
	if err := json.NewDecoder(r).Decode(rep); err != nil {
		return nil, err
	}

	return rep, nil
}

// Duration returns the duration of the replay.
func (r *Replay) Duration() (time.Duration, error) {
	if r.StartTimeUnix == 0 {
		return 0, errors.New("missing start time")
	}
	if r.EndTimeUnix == 0 {
		return 0, errors.New("missing end time")
	}
	start := r.StartTime()
	end := r.EndTime()

	return duration(start, end), nil
}

// duration returns the duration between two given times.
func duration(a, b time.Time) time.Duration {
	if b.Before(a) {
		temp := a
		a = b
		b = temp
	}

	return b.Sub(a)
}

// StartTime returns the time at which the match began.
func (r *Replay) StartTime() time.Time {
	return time.Unix(int64(r.StartTimeUnix), 0)
}

// EndTime returns the time at which the match ended.
func (r *Replay) EndTime() time.Time {
	return time.Unix(int64(r.EndTimeUnix), 0)
}
