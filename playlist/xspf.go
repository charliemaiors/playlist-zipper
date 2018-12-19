package playlist

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type xspfReader struct{}

//Playlist generated using xmltogo
type Playlist struct {
	XMLName xml.Name `xml:"playlist"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Title   struct {
		Text string `xml:",chardata"`
	} `xml:"title"`
	Creator struct {
		Text string `xml:",chardata"`
	} `xml:"creator"`
	Annotation struct {
		Text string `xml:",chardata"`
	} `xml:"annotation"`
	Info struct {
		Text string `xml:",chardata"`
	} `xml:"info"`
	Location struct {
		Text string `xml:",chardata"`
	} `xml:"location"`
	Identifier struct {
		Text string `xml:",chardata"`
	} `xml:"identifier"`
	Image struct {
		Text string `xml:",chardata"`
	} `xml:"image"`
	Date struct {
		Text string `xml:",chardata"`
	} `xml:"date"`
	License struct {
		Text string `xml:",chardata"`
	} `xml:"license"`
	Attribution struct {
		Text       string `xml:",chardata"`
		Identifier struct {
			Text string `xml:",chardata"`
		} `xml:"identifier"`
		Location struct {
			Text string `xml:",chardata"`
		} `xml:"location"`
	} `xml:"attribution"`
	Link struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
	} `xml:"link"`
	Meta struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
	} `xml:"meta"`
	Extension struct {
		Text        string `xml:",chardata"`
		Application string `xml:"application,attr"`
		Clip        struct {
			Text  string `xml:",chardata"`
			Start string `xml:"start,attr"`
			End   string `xml:"end,attr"`
		} `xml:"clip"`
	} `xml:"extension"`
	TrackList struct {
		Text  string `xml:",chardata"`
		Track []struct {
			Text     string `xml:",chardata"`
			Location struct {
				Text string `xml:",chardata"`
			} `xml:"location"`
			Identifier struct {
				Text string `xml:",chardata"`
			} `xml:"identifier"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Creator struct {
				Text string `xml:",chardata"`
			} `xml:"creator"`
			Annotation struct {
				Text string `xml:",chardata"`
			} `xml:"annotation"`
			Info struct {
				Text string `xml:",chardata"`
			} `xml:"info"`
			Image struct {
				Text string `xml:",chardata"`
			} `xml:"image"`
			Album struct {
				Text string `xml:",chardata"`
			} `xml:"album"`
			TrackNum struct {
				Text string `xml:",chardata"`
			} `xml:"trackNum"`
			Duration struct {
				Text string `xml:",chardata"`
			} `xml:"duration"`
			Link struct {
				Text string `xml:",chardata"`
				Rel  string `xml:"rel,attr"`
			} `xml:"link"`
			Meta struct {
				Text string `xml:",chardata"`
				Rel  string `xml:"rel,attr"`
			} `xml:"meta"`
			Extension struct {
				Text        string `xml:",chardata"`
				Application string `xml:"application,attr"`
				Clip        struct {
					Text  string `xml:",chardata"`
					Start string `xml:"start,attr"`
					End   string `xml:"end,attr"`
				} `xml:"clip"`
			} `xml:"extension"`
		} `xml:"track"`
	} `xml:"trackList"`
}

func (reader *xspfReader) ReadPlaylist(playlistLocation string) ([]string, error) {

	xspfPlaylist, err := ioutil.ReadFile(playlistLocation)
	if err != nil {
		return nil, err
	}

	var playlist Playlist
	err = xml.Unmarshal(xspfPlaylist, &playlist)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	for _, track := range playlist.TrackList.Track {
		fmt.Printf("Current is %s\n", track.Location.Text)
		if !isValidURL(track.Location.Text) {
			fmt.Printf("Adding %s\n", track.Location.Text)
			res = append(res, track.Location.Text)
		}
	}
	return res, nil
}
