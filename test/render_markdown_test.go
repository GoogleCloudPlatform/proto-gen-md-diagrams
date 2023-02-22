package test

import (
	"flag"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/markdown"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/reader"
	"github.com/stretchr/testify/assert"
)

var (
	path      = flag.String("p", "data/api", "path to proto buf directory")
	templates = flag.String("t", "data/templates", "the path to your template directory")
	debug     = flag.Bool("debug", false, "add debug logging.")
)

func TestMarkdownRender(t *testing.T) {
	log := logging.NewLogger(*debug, "test")
	log.Debug("Running Test")

	template := markdown.LoadTemplates(*templates)

	for _, t := range template.Templates() {
		log.Debugf("Template FQN: %v", t.Name())
	}

	assert.NotNil(t, template)

	packages, err := reader.ReadAllPackages(*path, *debug)

	if err != nil {
		log.Error(err.Error())
		assert.Fail(t, "Failed to load packages")
	}

	assert.NotNil(t, packages)
	assert.Equal(t, 2, len(packages))

	for k, v := range packages {
		log.Infof("Template for %s", k)
		err := template.ExecuteTemplate(os.Stdout, "base.tmpl", v)
		if err != nil {
			log.Errorf("error parsing template: %v", err)
		}
	}
}
