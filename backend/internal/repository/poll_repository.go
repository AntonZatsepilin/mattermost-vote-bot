package repository

import (
	"github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/models"
	"github.com/tarantool/go-tarantool/v2"
)

type PollTarantool struct {
	db *tarantool.Connection
}

func NewPollRepository(db *tarantool.Connection) *PollTarantool {
	return &PollTarantool{db: db}
}

func (p *PollTarantool) CreatePoll(poll models.Poll) error {
	_, err := p.db.Insert("polls", []interface{}{
		poll.ID,
		poll.Creator,
		poll.Question,
		poll.Options,
		poll.Status,
		poll.CreatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}