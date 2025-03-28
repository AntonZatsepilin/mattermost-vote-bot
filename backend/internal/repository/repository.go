package repository

import "github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/models"

type PollRepository interface {
	CreatePoll(poll models.Poll) error
}

type Repository struct {
	PollRepository
}

func NewRepository(pollRepo PollRepository) *Repository {
	return &Repository{
		PollRepository: pollRepo,
	}
}