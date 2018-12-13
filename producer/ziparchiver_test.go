package producer_test

import (
	"testing"

	"github.com/charliemaiors/producer"
)

var zipper producer.ArchiveProducer
var err error

func init() {
	zipper, err = producer.NewArchiver("zip")
	if err != nil {
		panic(err)
	}
}

func TestArchive(t *testing.T) {
	files := []string{"test/ciao.txt", "test/lol.txt"}
	err = zipper.Archive(files, "test/archive.zip")

	if err != nil {
		t.Fatalf("Error creating and defining the archive %v", err)
	}

	t.Log("Created success!!!")
}
