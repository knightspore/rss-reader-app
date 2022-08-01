package vo

type Subscription struct {
	ID string `xml:"channel>id" json:"id"`
	Title string `xml:"channel>title" json:"title"`
	Description string `xml:"channel>description" json:"description"`
	URL        string   `xml:"channel>link" json:"url"`
	LastUpdated string `xml:"channel>updated" json:"lastUpdated"`
	Muted bool `json:"channel>muted"`
	Icon string `json:"channel>icon"`
	Articles []Article `xml:"channel>item" json:"articles"`
}

func NewSubscription(url, title) *Subscription {
	data, err := Network.Fetch(url)
	if err != nil {
		fmt.Printf("Error Fetching Feed URL: %s", url)
		return vo.Subscription{}, err	
	}

	var s vo.Subscription
	err = xml.Unmarshal(data, &s)

	if err != nil {
		fmt.Printf("Error Unmarshaling XML: %s", url)
	}

	if title != "" {
		s.Title = title
	}

	return &s

}

