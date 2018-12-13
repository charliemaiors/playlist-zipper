package producer

import (
	"errors"

	"github.com/mholt/archiver"
)

type ArchiveProducer = archiver.Archiver

//NewArchiver returns a new instance of Archiver by given target zip file
func NewArchiver(archiveType string) (ArchiveProducer, error) {
	switch archiveType {
	case "zip":
		archiveProducer := archiver.NewZip()
		archiveProducer.CompressionLevel = 9
		archiveProducer.ContinueOnError = true
		return archiveProducer, nil
	case "tar":
		archiveProducer := archiver.NewTarGz()
		archiveProducer.CompressionLevel = 9
		archiveProducer.ContinueOnError = true
		return archiveProducer, nil
	case "xz":
		archiveProducer := archiver.NewTarXz()
		archiveProducer.ContinueOnError = true
		return archiveProducer, nil
	case "bz2":
		archiveProducer := archiver.NewTarBz2()
		archiveProducer.CompressionLevel = 9
		return archiveProducer, nil
	default:
		return nil, errors.New("Archive format not yet supported")
	}
}
