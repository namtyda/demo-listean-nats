package main

import (
	"context"
	"flag"
	"log"

	"github.com/namtyda/demo-listean-nats/internal/app/cache"
	"github.com/namtyda/demo-listean-nats/internal/app/handler"
	client "github.com/namtyda/demo-listean-nats/internal/app/nats"
	"github.com/namtyda/demo-listean-nats/internal/app/service"
	"github.com/namtyda/demo-listean-nats/internal/app/web"
	"github.com/namtyda/demo-listean-nats/internal/db"
	"github.com/namtyda/demo-listean-nats/internal/repository"
)

func main() {
	var clusterID, serverID string
	flag.StringVar(&clusterID, "cid", "test-cluster", "Cluster ID")
	flag.StringVar(&serverID, "sid", "", "Client ID")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("You need pass subject ")
	}

	subject := args[0]

	ctx := context.Background()
	adpt, err := db.New(ctx)

	if err != nil {
		log.Fatalf("Err pgpool connect %s\n", err.Error())
	}

	repo := repository.New(adpt, ctx)
	cache := cache.New()
	service := service.New(repo, cache)
	handler := handler.New(service)
	service.FillCache()
	webServer := web.New(service)

	c := client.NewConnect(clusterID, serverID)
	sub, err := c.Sc.Subscribe(subject, handler.HandleSubcricbe)

	if err != nil {
		log.Fatalf("Error subscribe %s\n", err.Error())
	}

	defer (func() {
		sub.Unsubscribe()
		c.Sc.Close()
		adpt.Close()
	})()

	webServer.Run()
}
