package model

type AutocompReq struct {
	Prefix string
}

type PreviewReq struct {
	From     int
	FieldVal string
}

type RecipeReq struct {
	Id int64
}
