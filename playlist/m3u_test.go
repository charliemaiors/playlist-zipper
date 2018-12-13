package playlist_test

import (
	"testing"

	"github.com/charliemaiors/playlist-zipper/playlist"
)

var m3uReader playlist.Reader

func init() {
	m3uReader, err = playlist.NewReader("m3u")
	if err != nil {
		panic(err)
	}
}

func TestM3UPlaylist(test *testing.T) {
	sources, err := m3uReader.ReadPlaylist("./test.m3u")
	if err != nil {
		test.Fatalf("Got error %v", err)
	}
	test.Logf("Sources are %v", sources)
}
