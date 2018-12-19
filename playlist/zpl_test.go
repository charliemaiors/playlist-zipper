package playlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charliemaiors/playlist-zipper/playlist"
)

var zplReader playlist.Reader
var err error

func init() {
	zplReader, err = playlist.NewReader("zpl")
	if err != nil {
		panic(err)
	}
}

func TestZplPlaylist(test *testing.T) {
	sources, err := zplReader.ReadPlaylist("test/test.zpl")
	if err != nil {
		test.Fatalf("Got error %v", err)
	}
	assert.Equal(test, 127, len(sources), "Wrong number of sources read")
}
