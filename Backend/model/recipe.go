package model

type RecipePreview struct {
	Id          int64  `json:"id"`
	Minutes     int16  `json:"minutes"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LoadedRecipes struct {
	Recipes       []RecipePreview `json:"recipes"`
	AlreadyLoaded int             `json:"alreadyLoaded"`
}

type Recipe struct {
	N_Steps       int16     `json:"n_steps"`
	N_Ingredients int16     `json:"n_ingedients"`
	Minutes       int16     `json:"minutes"`
	Description   string    `json:"description"`
	Steps         []string  `json:"steps"`
	Tags          []string  `json:"tags"`
	Nutrition     []float32 `json:"nutrition"`
	Name          string    `json:"name"`
	Ingredients   []string  `json:"ingredients"`
}
