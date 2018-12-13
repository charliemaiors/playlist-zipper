package archiver_test

import (
	"testing"

	"github.com/charliemaiors/playlist-zipper/archiver"
)

var tararchiver archiver.Archiver

func init() {
	tararchiver, err = archiver.NewArchiver("tar")
	if err != nil {
		panic(err)
	}
}

func TestTarArchive(t *testing.T) {
	files := []string{"test/ciao.txt", "test/lol.txt"}
	err = tararchiver.Archive("test/archive", files)

	if err != nil {
		t.Fatalf("Error creating and defining the archive %v", err)
	}

	t.Log("Created success!!!")
}
