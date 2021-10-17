package router

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-chi/chi/v5"
)

var Es *elasticsearch.Client

func SetupRouter() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		entries := requestElasticSearch()
		json.NewEncoder(w).Encode(entries)
	})

	http.ListenAndServe(":8080", r)
}

func requestElasticSearch() []interface{} {
	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"ingredients": "pork chops",
			},
		},
		"size": 10,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(&buf),
		Es.Search.WithTrackTotalHits(true),
		Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	log.Println(strings.Repeat("=", 37))
	return r["hits"].(map[string]interface{})["hits"].([]interface{})
}
