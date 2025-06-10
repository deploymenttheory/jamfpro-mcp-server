package toolsets

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// Add mobile device methods to MockJamfProClient

// GetMobileDevices mocks the GetMobileDevices method
func (m *MockJamfProClient) GetMobileDevices() (*jamfpro.ResponseMobileDeviceList, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseMobileDeviceList), args.Error(1)
}

// GetMobileDeviceByID mocks the GetMobileDeviceByID method
func (m *MockJamfProClient) GetMobileDeviceByID(id string) (*jamfpro.ResourceMobileDevice, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceMobileDevice), args.Error(1)
}

// GetMobileDeviceByName mocks the GetMobileDeviceByName method
func (m *MockJamfProClient) GetMobileDeviceByName(name string) (*jamfpro.ResourceMobileDevice, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceMobileDevice), args.Error(1)
}

// GetMobileDeviceGroups mocks the GetMobileDeviceGroups method
func (m *MockJamfProClient) GetMobileDeviceGroups() (*jamfpro.ResponseMobileDeviceGroupsList, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseMobileDeviceGroupsList), args.Error(1)
}

// GetMobileDeviceGroupByID mocks the GetMobileDeviceGroupByID method
func (m *MockJamfProClient) GetMobileDeviceGroupByID(id string) (*jamfpro.ResourceMobileDeviceGroup, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceMobileDeviceGroup), args.Error(1)
}

// GetMobileDeviceApplications mocks the GetMobileDeviceApplications method
func (m *MockJamfProClient) GetMobileDeviceApplications() (*jamfpro.ResponseMobileDeviceApplicationsList, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseMobileDeviceApplicationsList), args.Error(1)
}

// GetMobileDeviceConfigurationProfiles mocks the GetMobileDeviceConfigurationProfiles method
func (m *MockJamfProClient) GetMobileDeviceConfigurationProfiles() (*jamfpro.ResponseMobileDeviceConfigurationProfilesList, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseMobileDeviceConfigurationProfilesList), args.Error(1)
}

// CreateMobileDevice mocks the CreateMobileDevice method
func (m *MockJamfProClient) CreateMobileDevice(device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error) {
	args := m.Called(device)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceMobileDevice), args.Error(1)
}

// UpdateMobileDeviceByID mocks the UpdateMobileDeviceByID method
func (m *MockJamfProClient) UpdateMobileDeviceByID(id string, device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error) {
	args := m.Called(id, device)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceMobileDevice), args.Error(1)
}

// DeleteMobileDeviceByID mocks the DeleteMobileDeviceByID method
func (m *MockJamfProClient) DeleteMobileDeviceByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestNewMobileDevicesToolset tests the NewMobileDevicesToolset function
func TestNewMobileDevicesToolset(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Verify the toolset was created successfully
	assert.NotNil(t, toolset)
	assert.Equal(t, "mobile-devices", toolset.GetName())
	assert.NotEmpty(t, toolset.GetDescription())
	assert.NotEmpty(t, toolset.GetTools())

	// Verify specific tools exist
	tools := toolset.GetTools()
	toolNames := make([]string, len(tools))
	for i, tool := range tools {
		toolNames[i] = tool.Name
	}

	expectedTools := []string{
		"get_mobile_devices",
		"get_mobile_device_by_id",
		"get_mobile_device_by_name",
		"get_mobile_device_groups",
		"get_mobile_device_group_by_id",
		"get_mobile_device_applications",
		"get_mobile_device_configuration_profiles",
		"delete_mobile_device",
		"create_mobile_device",
		"update_mobile_device_by_id",
		"get_mobile_device_template",
	}

	for _, expectedTool := range expectedTools {
		assert.Contains(t, toolNames, expectedTool, "Expected tool %s not found", expectedTool)
	}
}

// TestGetMobileDevices tests the get_mobile_devices tool
func TestGetMobileDevices(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDevices").Return(&jamfpro.ResponseMobileDeviceList{
		MobileDevices: []jamfpro.MobileDeviceListItem{
			{
				ID:   1,
				Name: "iPad-001",
			},
			{
				ID:   2,
				Name: "iPhone-002",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_devices", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "mobile devices")
	assert.Contains(t, result, "iPad-001")
	assert.Contains(t, result, "iPhone-002")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceByID tests the get_mobile_device_by_id tool
func TestGetMobileDeviceByID(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDeviceByID", "1").Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "iPad-001",
			DeviceName:  "iPad-001",
			Name:        "iPad-001",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Mobile device details for ID 1")
	assert.Contains(t, result, "iPad-001")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceByName tests the get_mobile_device_by_name tool
func TestGetMobileDeviceByName(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDeviceByName", "iPad-001").Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "iPad-001",
			DeviceName:  "iPad-001",
			Name:        "iPad-001",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_by_name", map[string]interface{}{
		"name": "iPad-001",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Mobile device details for name 'iPad-001'")
	assert.Contains(t, result, "iPad-001")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceGroups tests the get_mobile_device_groups tool
func TestGetMobileDeviceGroups(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDeviceGroups").Return(&jamfpro.ResponseMobileDeviceGroupsList{
		MobileDeviceGroups: []jamfpro.MobileDeviceGroupsListItem{
			{
				ID:   1,
				Name: "iOS Devices",
			},
			{
				ID:   2,
				Name: "iPad Group",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_groups", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "mobile device groups")
	assert.Contains(t, result, "iOS Devices")
	assert.Contains(t, result, "iPad Group")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceGroupByID tests the get_mobile_device_group_by_id tool
func TestGetMobileDeviceGroupByID(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDeviceGroupByID", "1").Return(&jamfpro.ResourceMobileDeviceGroup{
		ID:   1,
		Name: "iOS Devices",
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_group_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Mobile device group details for ID 1")
	assert.Contains(t, result, "iOS Devices")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceApplications tests the get_mobile_device_applications tool
func TestGetMobileDeviceApplications(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDeviceApplications").Return(&jamfpro.ResponseMobileDeviceApplicationsList{
		MobileDeviceApplications: []jamfpro.MobileDeviceApplicationsListItem{
			{
				ID:   1,
				Name: "TestApp",
			},
			{
				ID:   2,
				Name: "ProductivityApp",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_applications", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "mobile device applications")
	assert.Contains(t, result, "TestApp")
	assert.Contains(t, result, "ProductivityApp")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceConfigurationProfiles tests the get_mobile_device_configuration_profiles tool
func TestGetMobileDeviceConfigurationProfiles(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetMobileDeviceConfigurationProfiles").Return(&jamfpro.ResponseMobileDeviceConfigurationProfilesList{
		ConfigurationProfiles: []jamfpro.MobileDeviceConfigurationProfilesListItem{
			{
				ID:   1,
				Name: "WiFi Profile",
			},
			{
				ID:   2,
				Name: "Security Profile",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_configuration_profiles", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "mobile device configuration profiles")
	assert.Contains(t, result, "WiFi Profile")
	assert.Contains(t, result, "Security Profile")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestCreateMobileDevice tests the create_mobile_device tool
func TestCreateMobileDevice(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("CreateMobileDevice", mock.AnythingOfType("*jamfpro.ResourceMobileDevice")).Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Test iPad",
			DeviceName:  "Test iPad",
			Name:        "Test iPad",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "create_mobile_device", map[string]interface{}{
		"name":          "Test iPad",
		"serial_number": "DMQVGC0DHLA0",
		"udid":          "270aae10800b6e61a2ee2bbc285eb967050b6112",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully created mobile device 'Test iPad'")
	assert.Contains(t, result, "ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestCreateMobileDeviceWithOptionalFields tests create_mobile_device with optional fields
func TestCreateMobileDeviceWithOptionalFields(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("CreateMobileDevice", mock.AnythingOfType("*jamfpro.ResourceMobileDevice")).Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Test iPad",
			DeviceName:  "Test iPad",
			Name:        "Test iPad",
			AssetTag:    "ASSET001",
		},
		Location: jamfpro.MobileDeviceSubsetLocation{
			Username:     "testuser",
			EmailAddress: "test@example.com",
		},
	}, nil)

	// Call the tool with optional fields
	result, err := toolset.ExecuteTool(context.Background(), "create_mobile_device", map[string]interface{}{
		"name":          "Test iPad",
		"serial_number": "DMQVGC0DHLA0",
		"udid":          "270aae10800b6e61a2ee2bbc285eb967050b6112",
		"asset_tag":     "ASSET001",
		"username":      "testuser",
		"email_address": "test@example.com",
		"department":    "Engineering",
		"building":      "Main Campus",
		"is_purchased":  true,
		"vendor":        "Apple",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully created mobile device 'Test iPad'")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestUpdateMobileDeviceByID tests the update_mobile_device_by_id tool
func TestUpdateMobileDeviceByID(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Mock the initial GetMobileDeviceByID call (toolset gets current data first)
	mockClient.On("GetMobileDeviceByID", "1").Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Original iPad",
			DeviceName:  "Original iPad",
			Name:        "Original iPad",
		},
	}, nil)

	// Mock the UpdateMobileDeviceByID call
	mockClient.On("UpdateMobileDeviceByID", "1", mock.AnythingOfType("*jamfpro.ResourceMobileDevice")).Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Updated iPad",
			DeviceName:  "Updated iPad",
			Name:        "Updated iPad",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "update_mobile_device_by_id", map[string]interface{}{
		"id":   "1",
		"name": "Updated iPad",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully updated mobile device with ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestDeleteMobileDevice tests the delete_mobile_device tool
func TestDeleteMobileDevice(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("DeleteMobileDeviceByID", "1").Return(nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "delete_mobile_device", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Mobile device with ID 1 has been successfully deleted")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetMobileDeviceTemplate tests the get_mobile_device_template tool
func TestGetMobileDeviceTemplate(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Call the tool via ExecuteTool (no mock needed for template)
	result, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_template", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Mobile device template:")
	assert.Contains(t, result, "Example iPad")

	// No mock expectations needed since this doesn't call external APIs
}

// TestMobileDevicesExecuteToolInvalidTool tests calling an invalid tool
func TestMobileDevicesExecuteToolInvalidTool(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Call an invalid tool
	_, err := toolset.ExecuteTool(context.Background(), "invalid_tool", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown tool")
}

// TestMobileDevicesMissingRequiredArgument tests calling a tool without required arguments
func TestMobileDevicesMissingRequiredArgument(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Call get_mobile_device_by_id without the required id argument
	_, err := toolset.ExecuteTool(context.Background(), "get_mobile_device_by_id", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")
}

// TestCreateMobileDeviceMissingRequiredFields tests create_mobile_device without required fields
func TestCreateMobileDeviceMissingRequiredFields(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Test missing name
	_, err := toolset.ExecuteTool(context.Background(), "create_mobile_device", map[string]interface{}{
		"serial_number": "DMQVGC0DHLA0",
		"udid":          "270aae10800b6e61a2ee2bbc285eb967050b6112",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")

	// Test missing serial_number
	_, err = toolset.ExecuteTool(context.Background(), "create_mobile_device", map[string]interface{}{
		"name": "Test iPad",
		"udid": "270aae10800b6e61a2ee2bbc285eb967050b6112",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")

	// Test missing udid
	_, err = toolset.ExecuteTool(context.Background(), "create_mobile_device", map[string]interface{}{
		"name":          "Test iPad",
		"serial_number": "DMQVGC0DHLA0",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")
}

// TestUpdateMobileDeviceWithLocationFields tests updating location fields
func TestUpdateMobileDeviceWithLocationFields(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Mock the initial GetMobileDeviceByID call
	mockClient.On("GetMobileDeviceByID", "1").Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Test iPad",
			DeviceName:  "Test iPad",
			Name:        "Test iPad",
		},
		Location: jamfpro.MobileDeviceSubsetLocation{
			Username: "olduser",
		},
	}, nil)

	// Mock the UpdateMobileDeviceByID call
	mockClient.On("UpdateMobileDeviceByID", "1", mock.AnythingOfType("*jamfpro.ResourceMobileDevice")).Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Test iPad",
			DeviceName:  "Test iPad",
			Name:        "Test iPad",
		},
		Location: jamfpro.MobileDeviceSubsetLocation{
			Username:     "newuser",
			EmailAddress: "newuser@example.com",
		},
	}, nil)

	// Call the tool with location updates
	result, err := toolset.ExecuteTool(context.Background(), "update_mobile_device_by_id", map[string]interface{}{
		"id":            "1",
		"username":      "newuser",
		"email_address": "newuser@example.com",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully updated mobile device with ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestUpdateMobileDeviceWithPurchasingFields tests updating purchasing fields
func TestUpdateMobileDeviceWithPurchasingFields(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a mobile devices toolset
	toolset := NewMobileDevicesToolset(mockClient, logger)

	// Mock the initial GetMobileDeviceByID call
	mockClient.On("GetMobileDeviceByID", "1").Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Test iPad",
			DeviceName:  "Test iPad",
			Name:        "Test iPad",
		},
	}, nil)

	// Mock the UpdateMobileDeviceByID call
	mockClient.On("UpdateMobileDeviceByID", "1", mock.AnythingOfType("*jamfpro.ResourceMobileDevice")).Return(&jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "Test iPad",
			DeviceName:  "Test iPad",
			Name:        "Test iPad",
		},
		Purchasing: jamfpro.MobileDeviceSubsetPurchasing{
			IsPurchased: true,
			Vendor:      "Apple",
			PONumber:    "PO-12345",
		},
	}, nil)

	// Call the tool with purchasing updates
	result, err := toolset.ExecuteTool(context.Background(), "update_mobile_device_by_id", map[string]interface{}{
		"id":           "1",
		"is_purchased": true,
		"vendor":       "Apple",
		"po_number":    "PO-12345",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully updated mobile device with ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}
