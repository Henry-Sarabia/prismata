package main

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// Replay contains information on the replay of a Prismata match.
type Replay struct {
	Code         string   `json:"code"`
	DeckInfo     Deck     `json:"deckInfo"`
	EndCondition int      `json:"endCondition"`
	Format       int      `json:"format"`
	RawHash      int      `json:"rawHash"`
	PlayerInfo   []Player `json:"playerInfo"`
	// CommandInfo  CmdInfo    `json:"commandInfo"`
	TimeInfo    TimeInfo   `json:"timeInfo"`
	Result      int        `json:"result"`
	StartTime   float64    `json:"startTime"`
	VersionInfo Version    `json:"versionInfo"`
	Seed        int        `json:"seed"`
	RatingInfo  RatingInfo `json:"ratingInfo"`
	EndTime     float64    `json:"endTime"`
}

// Deck represents a collection of units.
type Deck struct {
	MergedDeck []Unit          `json:"mergedDeck"`
	Base       [][]interface{} `json:"base"`
	DeckName   string          `json:"deckName"`
	Randomizer [][]string      `json:"randomizer"`
}

// Unit represents a single deployable unit of play.
type Unit struct {
	BaseSet int    `json:"baseSet,omitempty"`
	Name    string `json:"name"`
	UIName  string `json:"UIName,omitempty"`
}

// Player contains information on a participating agent in a Prismata replay.
type Player struct {
	DisplayName      string   `json:"displayName"`
	Name             string   `json:"name"`
	LoadingCompleted bool     `json:"loadingCompleted"`
	Bot              string   `json:"bot"`
	Trophies         []string `json:"trophies"`
	ID               int      `json:"id"`
	PercentLoaded    float64  `json:"percentLoaded"`
	AvatarFrame      string   `json:"avatarFrame"`
	Portrait         string   `json:"portrait"`
}

// CmdInfo contains information on the sequence of commands executed by players
// in a Prismata replay. Multiple commands are executed per turn.
type CmdInfo struct {
	TimesRemaining     []int     `json:"timesRemaining"`
	TimeBanksRemaining []float64 `json:"timeBanksRemaining"`
	MoveDurations      []float64 `json:"moveDurations"`
	CommandList        []Cmd     `json:"commandList"`
	CommandTimes       []float64 `json:"commandTimes"`
	CommandForced      []bool    `json:"commandForced"`
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
	ScoreChanges   []int       `json:"scoreChanges"`
	FinalRatings   []Rating    `json:"finalRatings"`
	RatingChanges  [][]float64 `json:"ratingChanges"`
	InitialRatings []Rating    `json:"initialRatings"`
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

func parse(r io.Reader) (*Replay, error) {
	rep := &Replay{}
	if err := json.NewDecoder(r).Decode(rep); err != nil {
		return nil, err
	}

	return rep, nil
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

// Duration returns the duration of the replay in seconds.
func (r *Replay) Duration() (time.Duration, error) {
	if r.StartTime == 0 {
		return 0, errors.New("missing start time")
	}
	if r.EndTime == 0 {
		return 0, errors.New("missing end time")
	}
	start := time.Unix(int64(r.StartTime), 0)
	end := time.Unix(int64(r.EndTime), 0)

	return duration(start, end), nil
}
