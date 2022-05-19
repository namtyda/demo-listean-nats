package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/nats-io/stan.go"
)

func main() {
	var clusterID, serverID string
	flag.StringVar(&clusterID, "cid", "test-cluster", "Cluster ID")
	flag.StringVar(&serverID, "sid", "", "Server ID")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("You need pass subject and msg")
	}
	sc, err := stan.Connect(clusterID, serverID)

	if err != nil {
		log.Fatalf("Can't connect publish %s", err.Error())
	}

	subject, file := args[0], args[1]
	bstr, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error open file: %v\n", err)
	}

	sc.Publish(subject, bstr)
}
