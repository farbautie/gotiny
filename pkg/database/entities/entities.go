package entities

type Link struct {
	ID        string `json:"id"`
	Url       string `json:"url"`
	ShortUrl  string `json:"short_url"`
	Hits      int    `json:"hits"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
