package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/models"
	"github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/repository"
)

type PollServiceImpl struct {
	repo repository.PollRepository
}

func (s *PollServiceImpl) CreatePoll(creator, question string, options []string) (*models.Poll, error) {
    pollID := uuid.New().String()
    optionsMap := make(map[string]string)
    for i, opt := range options {
        optionsMap[string(rune(i+1))] = opt
    }
    
    poll := models.Poll{
        ID:        pollID,
        Creator:   creator,
        Question: question,
        Options:  optionsMap,
        Status:   "active",
        CreatedAt: time.Now(),
    }
    
    if err := s.repo.CreatePoll(poll); err != nil {
        return nil, err
    }
    return &poll, nil
}