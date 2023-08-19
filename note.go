package note

type Note struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:user_id db:"user_id binding:required`
	Text        string `json:"text" db:"text" binding:required`
	Description string `json:"description" db:"description"`
}
