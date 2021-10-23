package router

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"suchmaschinen/model"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/go-chi/chi/v5"
	"github.com/mitchellh/mapstructure"
)

var Es *elasticsearch.Client

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}

func SetupRouter() {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		entries := requestElasticSearch()
		json.NewEncoder(w).Encode(entries)
	})

	r.Post("/name_autocomplete", func(w http.ResponseWriter, r *http.Request) {
		var autoCompReq model.AutocompReq
		if err := json.NewDecoder(r.Body).Decode(&autoCompReq); err != nil {
			log.Printf("Error parsing the request body: %s", err)
		}
		suggestions := nameAutocomplete(autoCompReq.Prefix)
		json.NewEncoder(w).Encode(suggestions)
	})

	http.ListenAndServe(":8080", r)
}

var nameAutocompQuery map[string]interface{} = map[string]interface{}{
	"_source": "name_suggest",
	"suggest": map[string]interface{}{
		"recipe-suggest": map[string]interface{}{
			"prefix": "",
			"completion": map[string]interface{}{
				"field":           "name_suggest",
				"size":            7,
				"skip_duplicates": true,
			},
		},
	},
}

func nameAutocomplete(prefix string) []string {
	var result map[string]interface{}
	var buf bytes.Buffer
	nameAutocompQuery["suggest"].(map[string]interface{})["recipe-suggest"].(map[string]interface{})["prefix"] = prefix
	if err := json.NewEncoder(&buf).Encode(nameAutocompQuery); err != nil {
		log.Printf("Error encoding recipe suggestion query: %s", err)
	}
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(&buf),
		Es.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting recipe suggestion response: %s", err)
	}
	defer res.Body.Close()
	checkIfResponseIsError(res)
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	}
	options := result["suggest"].(map[string]interface{})["recipe-suggest"].([]interface{})[0].(map[string]interface{})["options"].([]interface{})

	var suggestions []string
	for _, opt := range options {
		var convOpt model.Option
		mapstructure.Decode(opt, &convOpt)
		suggestions = append(suggestions, convOpt.Text)
	}
	return suggestions
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
		log.Printf("Error encoding query: %s", err)
	}
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(&buf),
		Es.Search.WithTrackTotalHits(true),
		Es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	checkIfResponseIsError(res)
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
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

func checkIfResponseIsError(res *esapi.Response) {
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
}
