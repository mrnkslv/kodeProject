package note

type TextsInput struct {
	Texts []string `json:"texts" binding:"required"`
	Lang  string   `form:"lang,omitempty"`
}

type SpellResult struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Row         int      `json:"row"`
	Col         int      `json:"col"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

type InputResult []SpellResult
