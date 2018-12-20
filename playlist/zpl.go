package playlist

import (
	"encoding/xml"
	"io/ioutil"
)

type Smil struct {
	Head zplHead `xml:"head"`
	Body zplBody `xml:"body"`
}

type zplHead struct {
	Title string `xml:"title"`
}

type zplBody struct {
	Sequence []zplMedia `xml:"seq>media"`
}

type zplMedia struct {
	Src         string `xml:"src,attr"`
	AlbumTitle  string `xml:"albumTitle,attr"`
	AlbumArtist string `xml:"albumArtist,attr"`
	TrackTitle  string `xml:"trackTitle,attr"`
	TrackArtist string `xml:"trackArtist,attr"`
	Duration    string `xml:"duration,attr"`
}

type zplReader struct {
}

func (seq zplBody) getAllSources() []string {
	res := make([]string, 0)
	for _, element := range seq.Sequence {
		if !isValidURL(element.Src) {
			res = append(res, element.Src)
		}
	}
	return res
}

func (reader *zplReader) ReadPlaylist(playlistLocation string) ([]string, error) {
	zplList, err := ioutil.ReadFile(playlistLocation)
	if err != nil {
		return nil, err
	}

	var playlist Smil
	err = xml.Unmarshal(zplList, &playlist)
	if err != nil {
		return nil, err
	}

	return playlist.Body.getAllSources(), nil
}
