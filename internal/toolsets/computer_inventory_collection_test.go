package toolsets

import (
	"context"
	"net/url"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// Helper function to create a mock client for testing
func createMockComputerInventoryClient() *mock.Mock {
	mockObj := &mock.Mock{}
	return mockObj
}

// ComputerInventoryClientAdapter implements the JamfProClient interface for computer inventory tests
type ComputerInventoryClientAdapter struct {
	mock *mock.Mock
}

// Create a new adapter
func NewComputerInventoryClientAdapter(mockObj *mock.Mock) *ComputerInventoryClientAdapter {
	return &ComputerInventoryClientAdapter{mock: mockObj}
}

// Computer Inventory specific methods
func (m *ComputerInventoryClientAdapter) GetComputersInventory(params url.Values) (*jamfpro.ResponseComputerInventoryList, error) {
	args := m.mock.Called(params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputerInventoryList), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) GetComputerInventoryByID(id string) (*jamfpro.ResourceComputerInventory, error) {
	args := m.mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceComputerInventory), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) GetComputerInventoryByName(name string) (*jamfpro.ResourceComputerInventory, error) {
	args := m.mock.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceComputerInventory), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) UpdateComputerInventoryByID(id string, inventory *jamfpro.ResourceComputerInventory) (*jamfpro.ResourceComputerInventory, error) {
	args := m.mock.Called(id, inventory)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourceComputerInventory), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) DeleteComputerInventoryByID(id string) error {
	args := m.mock.Called(id)
	return args.Error(0)
}

func (m *ComputerInventoryClientAdapter) GetComputersFileVaultInventory(params url.Values) (*jamfpro.FileVaultInventoryList, error) {
	args := m.mock.Called(params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.FileVaultInventoryList), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) GetComputerFileVaultInventoryByID(id string) (*jamfpro.FileVaultInventory, error) {
	args := m.mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.FileVaultInventory), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) GetComputerRecoveryLockPasswordByID(id string) (*jamfpro.ResponseRecoveryLockPassword, error) {
	args := m.mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseRecoveryLockPassword), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) RemoveComputerMDMProfile(id string) (*jamfpro.ResponseRemoveMDMProfile, error) {
	args := m.mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseRemoveMDMProfile), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) EraseComputerByID(id string, request jamfpro.RequestEraseDeviceComputer) error {
	args := m.mock.Called(id, request)
	return args.Error(0)
}

func (m *ComputerInventoryClientAdapter) UploadAttachmentAndAssignToComputerByID(computerID string, filePaths []string) (*jamfpro.ResponseUploadAttachment, error) {
	args := m.mock.Called(computerID, filePaths)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseUploadAttachment), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) DeleteAttachmentByIDAndComputerID(computerID, attachmentID string) error {
	args := m.mock.Called(computerID, attachmentID)
	return args.Error(0)
}

// Required methods to satisfy the JamfProClient interface
func (m *ComputerInventoryClientAdapter) GetJamfProInformation() (*jamfpro.ResponseJamfProInformation, error) {
	args := m.mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseJamfProInformation), args.Error(1)
}

func (m *ComputerInventoryClientAdapter) GetComputers() (*jamfpro.ResponseComputersList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetComputerByID(id string) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetComputerByName(name string) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetComputerGroups() (*jamfpro.ResponseComputerGroupsList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetComputerGroupByID(id string) (*jamfpro.ResourceComputerGroup, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) CreateComputer(computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdateComputerByID(id string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdateComputerByName(name string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) DeleteComputerByID(id string) error {
	return nil
}

func (m *ComputerInventoryClientAdapter) DeleteComputerByName(name string) error {
	return nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDevices() (*jamfpro.ResponseMobileDeviceList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDeviceByID(id string) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDeviceByName(name string) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDeviceGroups() (*jamfpro.ResponseMobileDeviceGroupsList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDeviceGroupByID(id string) (*jamfpro.ResourceMobileDeviceGroup, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDeviceApplications() (*jamfpro.ResponseMobileDeviceApplicationsList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetMobileDeviceConfigurationProfiles() (*jamfpro.ResponseMobileDeviceConfigurationProfilesList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) CreateMobileDevice(device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdateMobileDeviceByID(id string, device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) DeleteMobileDeviceByID(id string) error {
	return nil
}

func (m *ComputerInventoryClientAdapter) GetPolicies() (*jamfpro.ResponsePoliciesList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetPolicyByID(id string) (*jamfpro.ResourcePolicy, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetPolicyByName(name string) (*jamfpro.ResourcePolicy, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetPolicyByCategory(category string) (*jamfpro.ResponsePoliciesList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetPoliciesByType(createdBy string) (*jamfpro.ResponsePoliciesList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) CreatePolicy(policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdatePolicyByID(id string, policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdatePolicyByName(name string, policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) DeletePolicyByID(id string) error {
	return nil
}

func (m *ComputerInventoryClientAdapter) DeletePolicyByName(name string) error {
	return nil
}

func (m *ComputerInventoryClientAdapter) GetScripts(params url.Values) (*jamfpro.ResponseScriptsList, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetScriptByID(id string) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) GetScriptByName(name string) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) CreateScript(script *jamfpro.ResourceScript) (*jamfpro.ResponseScriptCreate, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdateScriptByID(id string, script *jamfpro.ResourceScript) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) UpdateScriptByName(name string, script *jamfpro.ResourceScript) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *ComputerInventoryClientAdapter) DeleteScriptByID(id string) error {
	return nil
}

func (m *ComputerInventoryClientAdapter) DeleteScriptByName(name string) error {
	return nil
}

// TestNewComputerInventoryToolset tests the NewComputerInventoryToolset function
func TestNewComputerInventoryToolset(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Verify the toolset was created successfully
	assert.NotNil(t, toolset)
	assert.Equal(t, "computer-inventory", toolset.GetName())
	assert.NotEmpty(t, toolset.GetDescription())
	assert.NotEmpty(t, toolset.GetTools())

	// Verify specific tools exist
	tools := toolset.GetTools()
	toolNames := make([]string, len(tools))
	for i, tool := range tools {
		toolNames[i] = tool.Name
	}

	// Verify basic inventory tools
	assert.Contains(t, toolNames, "get_computers_inventory")
	assert.Contains(t, toolNames, "get_computer_inventory_by_id")
	assert.Contains(t, toolNames, "get_computer_inventory_by_name")
	assert.Contains(t, toolNames, "update_computer_inventory")
	assert.Contains(t, toolNames, "delete_computer_inventory")

	// Verify FileVault tools
	assert.Contains(t, toolNames, "get_computers_filevault_inventory")
	assert.Contains(t, toolNames, "get_computer_filevault_inventory_by_id")
	assert.Contains(t, toolNames, "get_computer_recovery_lock_password")

	// Verify device management tools
	assert.Contains(t, toolNames, "remove_computer_mdm_profile")
	assert.Contains(t, toolNames, "erase_computer")

	// Verify attachment tools
	assert.Contains(t, toolNames, "upload_computer_attachment")
	assert.Contains(t, toolNames, "delete_computer_attachment")
}

// TestGetComputersInventory tests the get_computers_inventory tool
func TestGetComputersInventory(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetComputersInventory", mock.Anything).Return(&jamfpro.ResponseComputerInventoryList{
		TotalCount: 2,
		Results: []jamfpro.ResourceComputerInventory{
			{
				ID: "1",
				General: jamfpro.ComputerInventorySubsetGeneral{
					Name: "MacBook Pro 1",
				},
			},
			{
				ID: "2",
				General: jamfpro.ComputerInventorySubsetGeneral{
					Name: "MacBook Pro 2",
				},
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computers_inventory", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Found 2 computers")
	assert.Contains(t, result, "MacBook Pro 1")
	assert.Contains(t, result, "MacBook Pro 2")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetComputerInventoryByID tests the get_computer_inventory_by_id tool
func TestGetComputerInventoryByID(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetComputerInventoryByID", "1").Return(&jamfpro.ResourceComputerInventory{
		ID: "1",
		General: jamfpro.ComputerInventorySubsetGeneral{
			Name: "MacBook Pro 1",
		},
		Hardware: jamfpro.ComputerInventorySubsetHardware{
			Model: "MacBook Pro",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computer_inventory_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Computer inventory for ID 1")
	assert.Contains(t, result, "MacBook Pro 1")
	assert.Contains(t, result, "MacBook Pro")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetComputerInventoryByName tests the get_computer_inventory_by_name tool
func TestGetComputerInventoryByName(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetComputerInventoryByName", "MacBook Pro 1").Return(&jamfpro.ResourceComputerInventory{
		ID: "1",
		General: jamfpro.ComputerInventorySubsetGeneral{
			Name: "MacBook Pro 1",
		},
		Hardware: jamfpro.ComputerInventorySubsetHardware{
			Model: "MacBook Pro",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computer_inventory_by_name", map[string]interface{}{
		"name": "MacBook Pro 1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Computer inventory for name 'MacBook Pro 1'")
	assert.Contains(t, result, "MacBook Pro 1")
	assert.Contains(t, result, "MacBook Pro")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetComputersFileVaultInventory tests the get_computers_filevault_inventory tool
func TestGetComputersFileVaultInventory(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetComputersFileVaultInventory", mock.Anything).Return(&jamfpro.FileVaultInventoryList{
		TotalCount: 2,
		Results: []jamfpro.FileVaultInventory{
			{
				ComputerId:                          "1",
				Name:                                "MacBook Pro 1",
				IndividualRecoveryKeyValidityStatus: "VALID",
				InstitutionalRecoveryKeyPresent:     true,
			},
			{
				ComputerId:                          "2",
				Name:                                "MacBook Pro 2",
				IndividualRecoveryKeyValidityStatus: "VALID",
				InstitutionalRecoveryKeyPresent:     false,
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computers_filevault_inventory", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Found 2 computers with FileVault information")
	assert.Contains(t, result, "MacBook Pro 1")
	assert.Contains(t, result, "MacBook Pro 2")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetComputerFileVaultInventoryByID tests the get_computer_filevault_inventory_by_id tool
func TestGetComputerFileVaultInventoryByID(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetComputerFileVaultInventoryByID", "1").Return(&jamfpro.FileVaultInventory{
		ComputerId:                          "1",
		Name:                                "MacBook Pro 1",
		IndividualRecoveryKeyValidityStatus: "VALID",
		InstitutionalRecoveryKeyPresent:     true,
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computer_filevault_inventory_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "FileVault inventory for computer ID 1")
	assert.Contains(t, result, "MacBook Pro 1")
	assert.Contains(t, result, "VALID")
	assert.Contains(t, result, "true")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestExecuteToolInvalidTool tests calling an invalid tool
func TestComputerInventoryExecuteToolInvalidTool(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Call an invalid tool
	_, err := toolset.ExecuteTool(context.Background(), "invalid_tool", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown tool")
}

// TestMissingRequiredArgument tests calling a tool without required arguments
func TestComputerInventoryMissingRequiredArgument(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockComputerInventoryClient()
	mockClient := NewComputerInventoryClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computer inventory toolset
	toolset := NewComputerInventoryToolset(mockClient, logger)

	// Call get_computer_inventory_by_id without the required id argument
	_, err := toolset.ExecuteTool(context.Background(), "get_computer_inventory_by_id", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")
}
