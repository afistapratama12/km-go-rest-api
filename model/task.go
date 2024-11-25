package model

type Task struct {
	ID    int    `json:"id" faker:"int"` // 1, unik
	Title string `json:"title" faker:"title_male"`
	Body  string `json:"body" faker:"sentence"`
}

type TaskRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
