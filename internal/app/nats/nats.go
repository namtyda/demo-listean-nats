package nats

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

type NatsClient struct {
	Sc stan.Conn
}

func NewConnect(clusterID, clientID string) *NatsClient {
	sc, err := stan.Connect(clusterID, clientID)

	if err != nil {
		log.Fatalf("Error with connect to nats server %s", err.Error())
	}
	return &NatsClient{Sc: sc}
}
