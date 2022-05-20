package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/namtyda/demo-listean-nats/internal/models"
	"github.com/nats-io/stan.go"
)

type service interface {
	AddOrder(orderUUID, rawJson string)
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{service: service}
}

func (n *handler) HandleSubcricbe(m *stan.Msg) {
	shape := new(models.Order)

	err := json.Unmarshal(m.Data, shape)

	if err != nil {
		log.Printf("Unmarshal err %s\n", err.Error())
		return
	}
	if shape.OrderUID == "" {
		log.Printf("Error model shape orders")
		return
	}

	n.service.AddOrder(shape.OrderUID, string(m.Data))

	fmt.Println(shape.OrderUID)
}
