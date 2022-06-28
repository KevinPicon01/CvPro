package models

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	UserId      string `json:"user_id"`
	CreateAt    string `json:"create_at"`
	Description string `json:"description"`
	link        string `json:"link"`
	Category    string `json:"category"`
	likes       int    `json:"likes"`
}
