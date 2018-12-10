package cmd_test

import (
	"testing"

	"github.com/charliemaiors/playlist-zipper/cmd"
)

func TestProduceArchives(test *testing.T) {
	cmd.Execute()
	test.Logf("It works")
}
