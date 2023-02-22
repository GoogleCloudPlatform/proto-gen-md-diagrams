package markdown

import (
	"flag"
	"fmt"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
)

var Log *logging.Logger

const templateSuffix = ".tmpl"
const defaultTemplateDirectory = "templates"

func init() {
	var debug = true
	fmt.Println(flag.Lookup("debug"))
	if flag.Lookup("debug") != nil {
		b, err := strconv.ParseBool(flag.Lookup("debug").Value.String())
		if err == nil && b {
			debug = true
		}
	}
	Log = logging.NewLogger(debug, "markdown")
	Log.Debug("Enabled Debug for markdown")
}

func LoadTemplates(dir string) *template.Template {
	return template.Must(template.New("base").ParseGlob(filepath.Join(dir, "*.tmpl")))
}
