package service

import (
	"sync"

	"github.com/namtyda/demo-listean-nats/internal/repository"
)

type service struct {
	m    sync.RWMutex
	repo repository.Repository
}

func New(repo repository.Repository) *service {
	return &service{repo: repo, m: sync.RWMutex{}}
}

func (s *service) AddOrder() {

}
