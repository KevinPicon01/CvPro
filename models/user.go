package models

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Views       int    `json:"views"`
	Linkedin    string `json:"linkedin"`
	Github      string `json:"github"`
	Twitter     string `json:"twitter"`
}
type User_education struct {
	Id         string `json:"id"`
	User_id    string `json:"user_id"`
	School     string `json:"school"`
	Tittle     string `json:"tittle"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}
type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
