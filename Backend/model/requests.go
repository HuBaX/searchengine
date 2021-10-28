package model

type AutocompReq struct {
	Prefix string
}

type PreviewReq struct {
	From int
	Name string
}

type RecipeReq struct {
	Id int64
}

type FilterPreviewReq struct {
	From        int
	Name        string
	Tags        []string
	Ingredients []string
}
