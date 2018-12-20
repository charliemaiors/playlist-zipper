package playlist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charliemaiors/playlist-zipper/playlist"
)

var xspfReader playlist.Reader

func init() {
	xspfReader, err = playlist.NewReader("xspf")
	if err != nil {
		panic(err)
	}
}

func TestXspfPlaylist(test *testing.T) {
	sources, err := zplReader.ReadPlaylist("test/test.xspf")
	if err != nil {
		test.Fatalf("Got error %v", err)
	}
	fmt.Printf("Sources are %+v", sources)
	assert.Equal(test, 1, len(sources), "Wrong number of sources read")
}
