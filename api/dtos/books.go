package dtos

type BookData struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookDataV2 struct {
	ID     uint   `json:"book_id"`
	Title  string `json:"book_title"`
	Author string `json:"book_author"`
}
