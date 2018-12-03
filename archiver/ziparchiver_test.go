package archiver_test

import (
	"testing"

	"github.com/charliemaiors/playlist-zipper/archiver"
)

var zipper archiver.Archiver
var err error

func init() {
	zipper, err = archiver.NewArchiver("zip")
	if err != nil {
		panic(err)
	}
}

func TestArchive(t *testing.T) {
	files := []string{"test/ciao.txt", "test/lol.txt"}
	err = zipper.Archive("test/archive.zip", files)

	if err != nil {
		t.Fatalf("Error creating and defining the archive %v", err)
	}

	t.Log("Created success!!!")
}
