package playlist_test

import (
	"testing"

	"github.com/charliemaiors/playlist-zipper/playlist"
)

var reader playlist.Reader
var err error

func init() {
	reader, err = playlist.NewReader("zpl")
	if err != nil {
		panic(err)
	}
}

func TestPlaylist(test *testing.T) {
	sources, err := reader.ReadPlaylist("./test.zpl")
	if err != nil {
		test.Fatalf("Got error %v", err)
	}
	test.Logf("Sources are %v", sources)
}
