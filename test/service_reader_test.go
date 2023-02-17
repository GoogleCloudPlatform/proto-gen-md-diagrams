package test

import (
	"testing"

	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
	"github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/reader"
	"github.com/stretchr/testify/assert"
)

func TestServiceReader(t *testing.T) {

	log := logging.NewLogger(true, "service read test")
	pkg, err := reader.ReadPackage("data/test/service/service.proto", false)

	if err != nil {
		log.Errorf("Failed to read service file: %v", err)
		assert.Fail(t, "TestServiceReader Failed")
	}

	assert.NotNil(t, pkg)
	assert.Equal(t, 2, len(pkg.Messages()))
	assert.Equal(t, 1, len(pkg.Services()))

	service := pkg.Services()[0]

	assert.Equal(t, "LocationService", service.Name())

	assert.Equal(t, 2, len(service.RemoteProcedureCalls()))

}
