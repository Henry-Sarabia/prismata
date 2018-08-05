package main

// Player contains information on a participating agent in a Prismata replay.
type Player struct {
	Name             string   `json:"name"`
	DisplayName      string   `json:"displayName"`
	ID               int      `json:"id"`
	Bot              string   `json:"bot"`
	AvatarFrame      string   `json:"avatarFrame"`
	Portrait         string   `json:"portrait"`
	Trophies         []string `json:"trophies"`
	LoadingCompleted bool     `json:"loadingCompleted"`
	PercentLoaded    float64  `json:"percentLoaded"`
}
