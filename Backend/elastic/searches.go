package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"

	"suchmaschinen/model"
	"suchmaschinen/util"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/mitchellh/mapstructure"
)

var Es *elasticsearch.Client

func SearchRecipesByName(fuzzyReq model.PreviewReq) []model.RecipePreview {
	var nameSearchQuery map[string]interface{} = map[string]interface{}{
		"from": fuzzyReq.From,
		"size": 15,
		"query": map[string]interface{}{
			"fuzzy": map[string]interface{}{
				"name": map[string]interface{}{
					"value":     fuzzyReq.Name,
					"fuzziness": 3,
				},
			},
		},
	}
	var result map[string]interface{}
	var buf bytes.Buffer
	util.JsonEncode(nameSearchQuery, &buf)
	res, err := executeSearch(&buf)
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

func NameAutocomplete(prefix string) []string {
	var nameAutocompQuery map[string]interface{} = map[string]interface{}{
		"_source": "name_suggest",
		"suggest": map[string]interface{}{
			"recipe-suggest": map[string]interface{}{
				"prefix": prefix,
				"completion": map[string]interface{}{
					"field":           "name_suggest",
					"size":            7,
					"skip_duplicates": true,
				},
			},
		},
	}
	var result map[string]interface{}
	var buf bytes.Buffer
	util.JsonEncode(nameAutocompQuery, &buf)
	res, err := executeSearch(&buf)

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

func IngredientAutocomplete(prefix string) []string {
	var ingredientsAutocompQuery map[string]interface{} = map[string]interface{}{
		"_source": "ingredients_suggest",
		"suggest": map[string]interface{}{
			"ingredients-suggest": map[string]interface{}{
				"prefix": prefix,
				"completion": map[string]interface{}{
					"field":           "ingredients_suggest",
					"size":            7,
					"skip_duplicates": true,
				},
			},
		},
	}
	var result map[string]interface{}
	var buf bytes.Buffer
	util.JsonEncode(ingredientsAutocompQuery, &buf)
	res, err := executeSearch(&buf)
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

func TagsAutocomplete(prefix string) []string {
	var tagsAutocompQuery map[string]interface{} = map[string]interface{}{
		"_source": "tags_suggest",
		"suggest": map[string]interface{}{
			"tags-suggest": map[string]interface{}{
				"prefix": prefix,
				"completion": map[string]interface{}{
					"field":           "tags_suggest",
					"size":            7,
					"skip_duplicates": true,
				},
			},
		},
	}
	var result map[string]interface{}
	var buf bytes.Buffer
	util.JsonEncode(tagsAutocompQuery, &buf)
	res, err := executeSearch(&buf)
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

func SearchRecipeById(req model.RecipeReq) model.Recipe {
	var recipeQuery map[string]interface{} = map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.Id,
			},
		},
		"size": 1,
	}
	var result map[string]interface{}
	var buf bytes.Buffer
	util.JsonEncode(recipeQuery, &buf)
	res, err := executeSearch(&buf)
	if err != nil {
		log.Printf("Error getting ingredient suggestion response: %s", err)
	}
	defer res.Body.Close()
	checkIfResponseIsError(res)
	util.JsonDecode(res.Body, &result)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	var hit map[string]interface{}
	if len(hits) > 0 {
		hit = hits[0].(map[string]interface{})
	}
	var recipe model.Recipe
	mapstructure.Decode(hit["_source"], &recipe)
	return recipe
}

func SearchRecipesByFilter(filterReq model.FilterPreviewReq) []model.RecipePreview {

	var recipeFilterQuery map[string]interface{} = map[string]interface{}{
		"from": 0,
		"size": 15,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{},
			},
		},
	}
	var result map[string]interface{}
	var buf bytes.Buffer
	recipeFilterQuery["from"] = filterReq.From
	must := recipeFilterQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{})
	if len(filterReq.Name) > 0 {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"name": map[string]interface{}{
					"query":     filterReq.Name,
					"fuzziness": 3,
				},
			},
		})
	}
	if len(filterReq.Tags) > 0 {
		must = append(must, map[string]interface{}{
			"terms": map[string]interface{}{
				"tags": filterReq.Tags,
			},
		})
	}
	if len(filterReq.Ingredients) > 0 {
		must = append(must, map[string]interface{}{
			"terms": map[string]interface{}{
				"ingredients": filterReq.Ingredients,
			},
		})
	}
	util.JsonEncode(recipeFilterQuery, &buf)
	res, err := executeSearch(&buf)
	if err != nil {
		log.Printf("Error getting ingredient suggestion response: %s", err)
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

func executeSearch(query io.Reader) (*esapi.Response, error) {
	return Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex("gami1018_rufl1020_recipes"),
		Es.Search.WithBody(query),
		Es.Search.WithPretty(),
	)
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
