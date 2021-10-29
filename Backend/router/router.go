package router

import (
	"net/http"

	"suchmaschinen/elastic"
	"suchmaschinen/model"
	"suchmaschinen/util"

	"github.com/go-chi/chi/v5"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetupRouter() {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	/*r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))*/
	r.Post("/name_autocomplete", nameAutocomplete)

	r.Post("/name_search", nameSearch)

	r.Post("/ingredients_autocomplete", ingredientsAutocomplete)

	r.Post("/tags_autocomplete", tagsAutocomplete)

	r.Post("/recipe_search", recipeSearch)

	r.Post("/filter_search", filterSearch)

	http.ListenAndServe(":8080", r)
}

func nameAutocomplete(w http.ResponseWriter, r *http.Request) {
	var autoCompReq model.AutocompReq
	util.JsonDecode(r.Body, &autoCompReq)
	suggestions := elastic.NameAutocomplete(autoCompReq.Prefix)
	util.JsonEncode(suggestions, w)
}

func ingredientsAutocomplete(w http.ResponseWriter, r *http.Request) {
	var autoCompReq model.AutocompReq
	util.JsonDecode(r.Body, &autoCompReq)
	suggestions := elastic.IngredientAutocomplete(autoCompReq.Prefix)
	util.JsonEncode(suggestions, w)
}

func tagsAutocomplete(w http.ResponseWriter, r *http.Request) {
	var autoCompReq model.AutocompReq
	util.JsonDecode(r.Body, &autoCompReq)
	suggestions := elastic.TagsAutocomplete(autoCompReq.Prefix)
	util.JsonEncode(suggestions, w)
}

func nameSearch(w http.ResponseWriter, r *http.Request) {
	var fuzzyReq model.PreviewReq
	util.JsonDecode(r.Body, &fuzzyReq)
	recipes := elastic.SearchRecipesByName(fuzzyReq)
	response := model.LoadedRecipes{Recipes: recipes, AlreadyLoaded: fuzzyReq.From + len(recipes)}
	util.JsonEncode(response, w)
}

func recipeSearch(w http.ResponseWriter, r *http.Request) {
	var recipeReq model.RecipeReq
	util.JsonDecode(r.Body, &recipeReq)
	recipe := elastic.SearchRecipeById(recipeReq)
	util.JsonEncode(recipe, w)
}

func filterSearch(w http.ResponseWriter, r *http.Request) {
	var filterReq model.FilterPreviewReq
	util.JsonDecode(r.Body, &filterReq)
	recipes := elastic.SearchRecipesByFilter(filterReq)
	response := model.LoadedRecipes{Recipes: recipes, AlreadyLoaded: filterReq.From + len(recipes)}
	util.JsonEncode(response, w)
}
