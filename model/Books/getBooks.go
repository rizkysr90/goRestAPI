package book

type Info struct {
}
type GetBook struct {
	VolumeInfo struct {
		Title   string   `json:"title"`
		Authors []string `json:"authors"`
		Cover   struct {
			Medium string `json:"medium"`
		} `json:"imageLinks"`
	} `json:"volumeInfo"`
}
