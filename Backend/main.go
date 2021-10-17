package main

import (
	"log"
	"suchmaschinen/router"

	"github.com/elastic/go-elasticsearch/v7"
)

var Es *elasticsearch.Client

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://node-1.hska.io:9200/",
			"http://node-2.hska.io:9200/",
		},
	}

	Es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := Es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
	router.Es = Es
	router.SetupRouter()
}
