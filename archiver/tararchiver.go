package archiver

import (
	"archive/tar"
	"io"
	"os"
)

type tarArchiver struct{}

func (archiver *tarArchiver) Archive(filename string, files []string) error {
	tarFile, err := os.Create(filename + "tar.gz")
	if err != nil {
		return err
	}

	writer := tar.NewWriter(tarFile)
	defer writer.Close()

	for _, file := range files {
		currentfile, err := os.Open(file)
		if err != nil {
			return err
		}

		currentstats, err := currentfile.Stat()
		if err != nil {
			return err
		}
		currentHeader, err := tar.FileInfoHeader(currentstats, currentstats.Name())
		if err != nil {
			return err
		}
		if err := writer.WriteHeader(currentHeader); err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(writer, currentfile); err != nil {
			return err
		}

		currentfile.Close()
	}
	return nil
}
