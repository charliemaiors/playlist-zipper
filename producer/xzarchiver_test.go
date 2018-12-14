package producer_test

import (
	"testing"

	"github.com/charliemaiors/playlist-zipper/producer"
)

var xzarchiver producer.ArchiveProducer

func init() {
	xzarchiver, err = producer.NewArchiver("xz")
	if err != nil {
		panic(err)
	}
}

func TestXzArchive(t *testing.T) {
	files := []string{"test/ciao.txt", "test/lol.txt"}
	err = xzarchiver.Archive(files, "test/archive.tar.xz")

	if err != nil {
		t.Fatalf("Error creating and defining the archive %v", err)
	}

	t.Log("Created success!!!")
}
