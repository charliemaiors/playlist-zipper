package playlist

import (
	"errors"
	"fmt"
	"os"

	"github.com/grafov/m3u8"
)

type m3uReader struct {
}

func (reader *m3uReader) ReadPlaylist(playlistLocation string) ([]string, error) {
	playlist, err := os.Open(playlistLocation)
	if err != nil {
		return nil, err
	}

	playlistParsed, listType, err := m3u8.DecodeFrom(playlist, true)

	if listType == m3u8.MASTER {
		return nil, errors.New("Playlist not supported")
	}

	mediapl := playlistParsed.(*m3u8.MediaPlaylist)
	res := make([]string, 0)

	for _, segment := range mediapl.Segments {
		if segment != nil {
			fmt.Printf("Current segment is %+v\n", segment)
			res = append(res, segment.URI)
		}
	}
	return res, nil
}
