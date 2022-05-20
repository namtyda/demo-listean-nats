package service

import (
	"log"
	"sync"

	"github.com/namtyda/demo-listean-nats/internal/models"
)

type repository interface {
	AddOrder(OrderUID, rawJson string) (err error)
	ReadAll() (rowSlice []models.Cache, err error)
}

type cache interface {
	Set(key, value string)
	Get(key string) (value string, err error)
}
type service struct {
	m     sync.RWMutex
	repo  repository
	cache cache
}

func New(repo repository, cache cache) *service {
	return &service{repo: repo, m: sync.RWMutex{}, cache: cache}
}

func (s *service) FillCache() {
	sl, _ := s.repo.ReadAll()
	for _, v := range sl {
		s.cache.Set(v.Order_uuid, v.Data)
	}
}

func (s *service) AddOrder(orderUUID, rawJson string) {
	if _, err := s.cache.Get(orderUUID); err != nil {
		s.cache.Set(orderUUID, rawJson)

		if err := s.repo.AddOrder(orderUUID, rawJson); err != nil {
			log.Printf("Error AddOrder %s\n", err.Error())
		}
	}
}

func (s *service) GetOrderFromCache(orderUUID string) (rawJson string, err error) {
	data, err := s.cache.Get(orderUUID)
	if err != nil {
		return rawJson, err
	}
	return data, err
}
