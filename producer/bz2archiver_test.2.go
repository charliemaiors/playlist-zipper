package producer_test

import (
	"testing"
)

var bz2archiver producer.ArchiveProducer

func init() {
	bz2archiver, err = producer.NewArchiver("bz2")
	if err != nil {
		panic(err)
	}
}

func TestArchive(t *testing.T) {
	files := []string{"test/ciao.txt", "test/lol.txt"}
	err = bz2archiver.Archive(files, "test/archive.bz2")

	if err != nil {
		t.Fatalf("Error creating and defining the archive %v", err)
	}

	t.Log("Created success!!!")
}
