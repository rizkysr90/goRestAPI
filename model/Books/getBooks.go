package book

type Info struct {
}
type GetBook struct {
	VolumeInfo struct {
		Title         string   `json:"title"`
		Authors       []string `json:"authors"`
		Categories    []string `json:"categories"`
		PublishedDate string   `json:"publishedDate"`
		Cover         struct {
			Medium string `json:"medium"`
		} `json:"imageLinks"`
	} `json:"volumeInfo"`
}
