package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
	"suchmaschinen/model"
	"suchmaschinen/util"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/mitchellh/mapstructure"
)

var Es *elasticsearch.Client

var nameSearchQuery map[string]interface{} = map[string]interface{}{
	"from": 0,
	"size": 15,
	"query": map[string]interface{}{
		"fuzzy": map[string]interface{}{
			"name": map[string]interface{}{
				"value":     "",
				"fuzziness": 3,
			},
		},
	},
}

func SearchRecipesByName(fuzzyReq model.PreviewReq) []model.RecipePreview {
	var result map[string]interface{}
	var buf bytes.Buffer
	nameSearchQuery["query"].(map[string]interface{})["fuzzy"].(map[string]interface{})["name"].(map[string]interface{})["value"] = fuzzyReq.FieldVal
	nameSearchQuery["from"] = fuzzyReq.From
	util.JsonEncode(nameSearchQuery, &buf)
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
	util.JsonDecode(res.Body, &result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	var recipes []model.RecipePreview
	for _, hit := range hits {
		var recipe model.RecipePreview
		mapstructure.Decode(hit.(map[string]interface{})["_source"], &recipe)
		recipes = append(recipes, recipe)
	}
	return recipes
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

func NameAutocomplete(prefix string) []string {
	var result map[string]interface{}
	var buf bytes.Buffer
	nameAutocompQuery["suggest"].(map[string]interface{})["recipe-suggest"].(map[string]interface{})["prefix"] = prefix
	util.JsonEncode(nameAutocompQuery, &buf)
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
	util.JsonDecode(res.Body, &result)
	options := result["suggest"].(map[string]interface{})["recipe-suggest"].([]interface{})[0].(map[string]interface{})["options"].([]interface{})

	var suggestions []string
	for _, opt := range options {
		var convOpt model.Option
		mapstructure.Decode(opt, &convOpt)
		suggestions = append(suggestions, convOpt.Text)
	}
	return suggestions
}

var ingredientsAutocompQuery map[string]interface{} = map[string]interface{}{
	"_source": "ingredients_suggest",
	"suggest": map[string]interface{}{
		"ingredients-suggest": map[string]interface{}{
			"prefix": "",
			"completion": map[string]interface{}{
				"field":           "ingredients_suggest",
				"size":            7,
				"skip_duplicates": true,
			},
		},
	},
}

func IngredientAutocomplete(prefix string) []string {
	var result map[string]interface{}
	var buf bytes.Buffer
	ingredientsAutocompQuery["suggest"].(map[string]interface{})["ingredients-suggest"].(map[string]interface{})["prefix"] = prefix
	util.JsonEncode(ingredientsAutocompQuery, &buf)
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(&buf),
		Es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting ingredient suggestion response: %s", err)
	}
	defer res.Body.Close()
	checkIfResponseIsError(res)
	util.JsonDecode(res.Body, &result)
	options := result["suggest"].(map[string]interface{})["ingredients-suggest"].([]interface{})[0].(map[string]interface{})["options"].([]interface{})
	var suggestions []string
	for _, opt := range options {
		var convOpt model.Option
		mapstructure.Decode(opt, &convOpt)
		suggestions = append(suggestions, convOpt.Text)
	}
	return suggestions
}

var tagsAutocompQuery map[string]interface{} = map[string]interface{}{
	"_source": "tags_suggest",
	"suggest": map[string]interface{}{
		"tags-suggest": map[string]interface{}{
			"prefix": "",
			"completion": map[string]interface{}{
				"field":           "tags_suggest",
				"size":            7,
				"skip_duplicates": true,
			},
		},
	},
}

func TagsAutocomplete(prefix string) []string {
	var result map[string]interface{}
	var buf bytes.Buffer
	tagsAutocompQuery["suggest"].(map[string]interface{})["tags-suggest"].(map[string]interface{})["prefix"] = prefix
	util.JsonEncode(tagsAutocompQuery, &buf)
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(&buf),
		Es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting ingredient suggestion response: %s", err)
	}
	defer res.Body.Close()
	checkIfResponseIsError(res)
	util.JsonDecode(res.Body, &result)
	options := result["suggest"].(map[string]interface{})["tags-suggest"].([]interface{})[0].(map[string]interface{})["options"].([]interface{})
	var suggestions []string
	for _, opt := range options {
		var convOpt model.Option
		mapstructure.Decode(opt, &convOpt)
		suggestions = append(suggestions, convOpt.Text)
	}
	return suggestions
}

var recipeQuery map[string]interface{} = map[string]interface{}{
	"query": map[string]interface{}{
		"match": map[string]interface{}{
			"id": -1,
		},
	},
	"size": 1,
}

func SearchRecipeById(req model.RecipeReq) model.Recipe {
	var result map[string]interface{}
	var buf bytes.Buffer
	recipeQuery["query"].(map[string]interface{})["match"].(map[string]interface{})["id"] = req.Id
	util.JsonEncode(recipeQuery, &buf)
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(&buf),
		Es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting ingredient suggestion response: %s", err)
	}
	defer res.Body.Close()
	checkIfResponseIsError(res)
	util.JsonDecode(res.Body, &result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	log.Print(len(hits))
	var hit map[string]interface{}
	if len(hits) > 0 {
		hit = hits[0].(map[string]interface{})
	}
	log.Print(hit)
	var recipe model.Recipe
	mapstructure.Decode(hit, &recipe)
	log.Print(recipe)
	return recipe
}

func RequestElasticSearch() []interface{} {
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
	util.JsonEncode(query, &buf)
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
	util.JsonDecode(res.Body, &r)
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
