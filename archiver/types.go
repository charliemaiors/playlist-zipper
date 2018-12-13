package archiver

import "errors"

type Archiver interface {
	Archive(fileName string, files []string) error
}

//NewArchiver returns a new instance of Archiver by given target zip file
func NewArchiver(archiveType string) (Archiver, error) {
	switch archiveType {
	case "zip":
		return &zipArchiver{}, nil
	case "tar":
		return &tarArchiver{}, nil
	default:
		return nil, errors.New("Archive format not yet supported")
	}
}
