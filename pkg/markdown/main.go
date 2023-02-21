package markdown

import (
	"flag"
	"fmt"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func LoadTemplates(dir string) map[string]string {
	if dir == "" {
		dir = defaultTemplateDirectory
	}

	Log.Debugf("Reading %s Directory for templates.", dir)

	out := make(map[string]string)

	err := filepath.Walk(defaultTemplateDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			Log.Errorf("%v", err)
		}
		if strings.HasSuffix(path, templateSuffix) {
			b, e := os.ReadFile(path)
			if e != nil {
				return e
			}
			out[info.Name()] = string(b)
		}
		return nil
	})
	if err != nil {
		Log.Errorf("failed to read file: %v", err)
	}

	return out
}

func DebugTemplates(in map[string]string) {
	if Log.IsDebug() {
		for k, v := range in {
			Log.Debugf("[Key: %s] Value Lengths: %d\n", k, len(v))
		}
	}
}

/*
Render is responsible for reading a root directory of protocol buffers, converting them into a DOM and using it
to populate templates from the provided or default template directory.
*/
func Render(root string, templateDir string) map[string]string {
	out := LoadTemplates(templateDir)
	DebugTemplates(out)

	return out
}
