/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package reader

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
)

func ReadPackage(fileName string, debug bool) (api.Package, error) {

	Log := logging.NewLogger(debug, "read_package")

	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := NewProtobufFileScanner(readFile)

	var pkg api.Package

	var comment = ""

	for scanner.Scan() {
		line := scanner.ReadLine()

		Log.Debugf("Current Line: `%s`\n", line)

		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(line) {
				qualifier := fileName
				if pkg != nil && len(pkg.Name()) > 0 {
					qualifier = pkg.Name()
				}
				rt := visitor.Visit(scanner, line, qualifier)
				switch t := rt.(type) {
				case api.Option:
					t.SetComment(SplitComment(Join(Space, comment, line.Comment)))
					pkg.AddOption(t.Name(), t.Value(), t.Comment())
					comment = ""
				case api.Import:
					t.SetComment(SplitComment(Join(Space, comment, line.Comment)))
					pkg.AddImport(t.Path(), t.Comment())
					comment = ""
				case api.Message:
					t.SetComment(SplitComment(Join(Space, comment, line.Comment)))
					pkg.AddMessage(t)
					comment = ""
				case api.Enum:
					t.SetComment(SplitComment(Join(Space, comment, line.Comment)))
					pkg.AddEnum(t)
					comment = ""
				case api.Service:
					t.SetComment(SplitComment(Join(Space, comment, line.Comment)))
					pkg.AddService(t)
					comment = ""
				case api.Package:
					t.SetComment(SplitComment(Join(Space, comment, line.Comment)))
					pkg = t
				case string:
					comment = Join(Space, comment, t)
				default:
					Log.Debugf("Unhandled Return type for package: %T visitor\n", t)
				}
			}
		}
	}
	return pkg, err
}

func ReadAllPackages(dir string, debug bool) (out map[string]api.Package, err error) {
	log := logging.NewLogger(debug, "read_all_packages")
	log.Debugf("Reading protobuf directory %s", dir)

	out = make(map[string]api.Package)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		log.Debugf("-- visiting %s", path)
		if strings.HasSuffix(path, ".proto") {
			log.Debugf("Reading protobuf: %s", path)
			pkg, e := ReadPackage(path, debug)
			if e != nil {
				log.Errorf("error reading package: %v", err)
			} else {
				out[pkg.Qualifier()] = pkg
			}
		}
		return nil
	})
	return out, err
}
