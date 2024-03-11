package proto

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"strings"
)

var directoryFlag *string
var recursiveFlag *bool
var debugFlag *bool
var writeOutputFlag *bool
var visualizeFlag *bool
var outputFlag *string

const (
	ProtobufSuffix = ".proto"
)

func init() {
	directoryFlag = flag.String("d", ".", "The directoryFlag to read.")
	recursiveFlag = flag.Bool("r", true, "Read recursively.")
	debugFlag = flag.Bool("debugFlag", false, "Enable debugging")
	writeOutputFlag = flag.Bool("w", true, "Enable writing output")
	visualizeFlag = flag.Bool("v", true, "Enable Visualization")
	outputFlag = flag.String("o", ".", "Specifies the outputFlag directoryFlag, if not specified, the processor will write markdown in the proto directories.")
}

func debugPackages(packages []*Package, logger *Logger) {
	if *debugFlag {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		for _, pkg := range packages {
			err := enc.Encode(pkg)
			if err != nil {
				logger.Errorf("Error encoding package %v", err)
			}
		}
	}
}

func Execute() {
	flag.Parse()

	SetDebug(*debugFlag)
	logger := Log
	logger.Infof("Reading Directory : %s\n", *directoryFlag)
	logger.Infof("Recursively: %v\n", *recursiveFlag)

	packages := make([]*Package, 0)

	err := filepath.Walk(*directoryFlag, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ProtobufSuffix) {

			pkg := NewPackage(path)
			err := pkg.Read(*debugFlag)
			if err != nil {
				logger.Errorf("error while reading package %s with value: %v", path, err)
			}
			packages = append(packages, pkg)
		}
		return nil
	})

	// Send outputFlag to debugFlag if enabled.
	debugPackages(packages, logger)

	if err != nil {
		logger.Errorf("failed to process directoryFlag: %s with error: %v", *directoryFlag, err)
	}

	for _, pkg := range packages {
		bName := filepath.Base(pkg.Path)
                out := *outputFlag + string(filepath.Separator) + bName + ".md"
		markdown := PackageToMarkDown(pkg, *visualizeFlag)
		if *writeOutputFlag {
			err = os.WriteFile(out, []byte(markdown), 0644)
		}
		if err != nil {
			logger.Errorf("failed to write file %v\n", err)
		}
	}
}
