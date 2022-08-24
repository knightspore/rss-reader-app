package parse

import (
	"encoding/json"
	"encoding/xml"
	"errors"

	"github.com/knightspore/rss-reader-app/backend/util"
)

type Feed struct {
	URL string
	Doc []byte
	JSON []byte
	Type string
}

func NewFeed(url string) (Feed, error ) {

	var f Feed

	f.URL = url

	doc, err := util.Fetch(f.URL)
	if err != nil {
		return f, err
	}
	f.Doc = doc

	t, err := detectFeedType(string(f.Doc))
	if err != nil {
		return f, err
	}
	f.Type = t

	if len(f.Doc) == 0 || f.Type == "" {
		return f, errors.New("could not parse Feed from URL")
	}

	j, err := f.json()
	if err != nil {
		return f, errors.New("could not create Subscription JSON from Feed")
	}
	f.JSON = j

	return f, nil

}

func detectFeedType(doc string) (string, error) {

	type detector struct {
		RSS string `xml:"channel>title"`
		Atom string `xml:"title"`
	}

	var d detector

	err := xml.Unmarshal([]byte(doc), &d)
	if err != nil {
		return "", err
	}

	if d.RSS != "" {
		return "rss", nil
	} 
	
	if d.Atom != "" {
		return "atom", nil
	}

	return "", errors.New("no feed type detected")

}

func (f *Feed) json() ([]byte, error) {

	var j []byte

	switch {
	case f.Type == "rss":
		rss, err := UnmarshalRSS(f.Doc)
		if err != nil {
			return j, err
		}
		j, err = json.Marshal(rss)
		return j, err
	case f.Type == "atom":
		atom, err := UnmarshalAtom(f.Doc)
		if err != nil {
			return j, err
		}
		j, err = json.Marshal(atom)
		return j, err
	}

	return j, nil

}