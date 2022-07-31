package handler

type User string

type Feed struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Event struct {
	UserID string `json:"user_id"`
	Feeds  []Feed `json:"feeds,omitempty"`
	ReadPostUrl string `json:"read_post_url,omitempty"`
}

type PostContent struct {
	Title   string `json:"title"`
	Byline  string `json:"byline"`
	Site    string `json:"site"`
	Text    string `json:"text"`
	Favicon string `json:"favicon"`
	URL 		string `json:"url"`
}
