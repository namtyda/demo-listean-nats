package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/namtyda/demo-listean-nats/internal/app/cache"
	client "github.com/namtyda/demo-listean-nats/internal/app/nats"
	"github.com/namtyda/demo-listean-nats/internal/app/service"
	"github.com/namtyda/demo-listean-nats/internal/db"
	"github.com/namtyda/demo-listean-nats/internal/models"
	"github.com/namtyda/demo-listean-nats/internal/repository"
	"github.com/nats-io/stan.go"
)

func main() {
	ctx := context.Background()
	adpt, err := db.New(ctx)

	if err != nil {
		log.Fatalf("Err pgpool connect %s\n", err.Error())
	}

	repo := repository.New(adpt)
	cache := cache.New()
	service := service.New(repo)

	sl, _ := repo.ReadAll()
	for _, v := range sl {
		cache.Set(v.Order_uuid, v.Data)
	}

	c := client.NewConnect()

	sub, err := c.Sc.Subscribe("test", func(m *stan.Msg) {
		shape := new(models.Order)

		err := json.Unmarshal(m.Data, shape)

		if err != nil {
			log.Fatalf("Unmarshal err %s\n", err.Error())
		}

		// fmt.Printf("Received a message: %s\n", string(m.Data))
		fmt.Println(shape.CustomerID)
	})

	if err != nil {
		log.Fatal("err sub")
	}
	defer sub.Unsubscribe()
	// c.Sc.Publish("foo", []byte("Hello World"))
	time.Sleep(time.Minute * 2)
}
