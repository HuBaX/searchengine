package model

type Option struct {
	Text     string
	_index   string
	_type    string
	_id      string
	_score   string
	_ignored []string
	_source  map[string]interface{}
}
