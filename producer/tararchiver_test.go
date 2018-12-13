package producer_test

import (
	"testing"

	"github.com/charliemaiors/producer"
)

var tararchiver producer.ArchiveProducer

func init() {
	tararchiver, err = producer.NewArchiver("tar")
	if err != nil {
		panic(err)
	}
}

func TestTarArchive(t *testing.T) {
	files := []string{"test/ciao.txt", "test/lol.txt"}
	err = tararchiver.Archive(files, "test/archive.tar.gz")

	if err != nil {
		t.Fatalf("Error creating and defining the archive %v", err)
	}

	t.Log("Created success!!!")
}
