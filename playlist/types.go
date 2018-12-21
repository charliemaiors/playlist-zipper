package playlist

import (
	"errors"
	"net/url"
	"strings"
)

var forbiddenProtocols = []string{"http", "https", "rtsp", "ftp"}

//Reader is the interface defined for categorize a general reader for different playlist types
type Reader interface {
	ReadPlaylist(playlistLocation string) ([]string, error)
}

//NewReader returns an instance of playlist.Reader interface (THIS IS NOT AN IMPLEMENTATION OF golang Reader)
func NewReader(playlistType string) (Reader, error) {
	switch playlistType {
	case "zpl":
		return &zplReader{}, nil
	case "m3u":
		return &m3uReader{}, nil
	case "xspf":
		return &xspfReader{}, nil
	default:
		return nil, errors.New("Type not supported")
	}
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	for _, frbd := range forbiddenProtocols {
		if strings.HasPrefix(toTest, frbd) {
			return true
		}
	}
	return false
}
