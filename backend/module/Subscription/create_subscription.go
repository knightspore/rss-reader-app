package subscription

import (
	"encoding/xml"
	"fmt"

	"github.com/knightspore/rss-reader-app/backend/module/Network"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

func Create(url string) (vo.Subscription, error) {

	data, err := Network.Fetch(url)
	if err != nil {
		fmt.Printf("Error Fetching URL: %s", url)
		return vo.Subscription{}, err	
	}

	var s vo.Subscription

	err = xml.Unmarshal(data, &s)
	if err != nil {
		fmt.Printf("Error Unmarshaling XML: %s", url)
		return vo.Subscription{}, err	
	}


	return s, nil
}
