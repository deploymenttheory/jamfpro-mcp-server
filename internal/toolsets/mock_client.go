package toolsets

import (
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/stretchr/testify/mock"
)

// MockJamfProClient implements the jamfpro.Client interface for testing
type MockJamfProClient struct {
	mock.Mock
	*jamfpro.Client // Embed the real client to satisfy interface requirements
}
