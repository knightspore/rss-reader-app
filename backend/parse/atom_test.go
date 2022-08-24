package parse_test

import (
	_ "embed"
	"testing"

	"github.com/knightspore/rss-reader-app/backend/parse"
)

func TestAtomUnmarshal(t *testing.T) {

	atom, err := parse.NewFeed("https://www.groundup.org.za/sitenews/atom/")
	if err != nil {
		t.Log("Error creating new Atom Feed")
		t.Fatal(err)
	}

	_, err = parse.UnmarshalAtom(atom.Doc)
	if err != nil {
		t.Fatal(err)
	}
}