package test

import (
	"flag"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/markdown"
	"testing"
)

var path = flag.String("path", "test/data/test", "path to proto buf directory")
var debug = flag.Bool("debug", false, "add debug logging.")

func TestMarkdownRender(t *testing.T) {
	log := logging.NewLogger(*debug, "test")
	log.Debug("Running Test")
	markdown.Render(*path, "data/templates/")
}
