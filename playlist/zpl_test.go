package playlist_test

import (
	"testing"

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
	sources, err := zplReader.ReadPlaylist("./test.zpl")
	if err != nil {
		test.Fatalf("Got error %v", err)
	}
	test.Logf("Sources are %v", sources)
}
