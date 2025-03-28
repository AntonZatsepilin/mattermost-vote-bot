package models

import (
	"time"
)

type Poll struct {
    ID        string
    Creator   string
    Question string
    Options  map[string]string
    Status   string
    CreatedAt time.Time
}

type Vote struct {
    PollID   string
    UserID   string
    Option  string
    Timestamp time.Time
}

type PollResults struct {
    Question string
    Options  map[string]OptionResult
    Total    int
}

type OptionResult struct {
    Text    string
    Count   int
    Percent float64
}