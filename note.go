package note

type Note struct {
	Id          int    `json:"id" db:"id"`
	Text        string `json:"text" db:"text" binding:required`
	Description string `json:"description" db:"description"`
}

type UserNote struct {
	Id       int
	UserId   int
	NoticeId int
}
