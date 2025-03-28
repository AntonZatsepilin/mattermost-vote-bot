package service

import "github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/models"

type PollService interface {
	CreatePoll(poll models.Poll) error
}

type Service struct {
	PollService
}