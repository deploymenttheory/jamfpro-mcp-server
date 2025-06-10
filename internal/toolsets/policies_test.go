package toolsets

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// Helper function to create a mock client for testing
func createMockPoliciesClient() *mock.Mock {
	mockObj := &mock.Mock{}
	return mockObj
}

// PolicyClientAdapter implements the JamfProClient interface for policies tests
type PolicyClientAdapter struct {
	mock *mock.Mock
}

// Create a new adapter
func NewPolicyClientAdapter(mockObj *mock.Mock) *PolicyClientAdapter {
	return &PolicyClientAdapter{mock: mockObj}
}

// Policy-specific methods
func (m *PolicyClientAdapter) GetPolicies() (*jamfpro.ResponsePoliciesList, error) {
	args := m.mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponsePoliciesList), args.Error(1)
}

func (m *PolicyClientAdapter) GetPolicyByID(id string) (*jamfpro.ResourcePolicy, error) {
	args := m.mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourcePolicy), args.Error(1)
}

func (m *PolicyClientAdapter) GetPolicyByName(name string) (*jamfpro.ResourcePolicy, error) {
	args := m.mock.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResourcePolicy), args.Error(1)
}

func (m *PolicyClientAdapter) GetPolicyByCategory(category string) (*jamfpro.ResponsePoliciesList, error) {
	args := m.mock.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponsePoliciesList), args.Error(1)
}

func (m *PolicyClientAdapter) GetPoliciesByType(createdBy string) (*jamfpro.ResponsePoliciesList, error) {
	args := m.mock.Called(createdBy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponsePoliciesList), args.Error(1)
}

func (m *PolicyClientAdapter) CreatePolicy(policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error) {
	args := m.mock.Called(policy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponsePolicyCreateAndUpdate), args.Error(1)
}

func (m *PolicyClientAdapter) UpdatePolicyByID(id string, policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error) {
	args := m.mock.Called(id, policy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponsePolicyCreateAndUpdate), args.Error(1)
}

func (m *PolicyClientAdapter) UpdatePolicyByName(name string, policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error) {
	args := m.mock.Called(name, policy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponsePolicyCreateAndUpdate), args.Error(1)
}

func (m *PolicyClientAdapter) DeletePolicyByID(id string) error {
	args := m.mock.Called(id)
	return args.Error(0)
}

func (m *PolicyClientAdapter) DeletePolicyByName(name string) error {
	args := m.mock.Called(name)
	return args.Error(0)
}

func (m *PolicyClientAdapter) GetJamfProInformation() (*jamfpro.ResponseJamfProInformation, error) {
	args := m.mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseJamfProInformation), args.Error(1)
}

// Required methods to satisfy the JamfProClient interface
func (m *PolicyClientAdapter) GetComputers() (*jamfpro.ResponseComputersList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerByID(id string) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerByName(name string) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerGroups() (*jamfpro.ResponseComputerGroupsList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerGroupByID(id string) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) CreateComputer(computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) UpdateComputerByID(id string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) UpdateComputerByName(name string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) DeleteComputerByID(id string) error {
	return nil
}

func (m *PolicyClientAdapter) DeleteComputerByName(name string) error {
	return nil
}

// Implement the rest of the interface methods with nil responses for our test
func (m *PolicyClientAdapter) GetComputersInventory(params interface{}) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerInventoryByID(id string) (*jamfpro.ResourceComputerInventory, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerInventoryByName(name string) (*jamfpro.ResourceComputerInventory, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) UpdateComputerInventoryByID(id string, inventory *jamfpro.ResourceComputerInventory) (*jamfpro.ResourceComputerInventory, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) DeleteComputerInventoryByID(id string) error {
	return nil
}

func (m *PolicyClientAdapter) GetComputersFileVaultInventory(params interface{}) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerFileVaultInventoryByID(id string) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetComputerRecoveryLockPasswordByID(id string) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) RemoveComputerMDMProfile(id string) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) EraseComputerByID(id string, request jamfpro.RequestEraseDeviceComputer) error {
	return nil
}

func (m *PolicyClientAdapter) UploadAttachmentAndAssignToComputerByID(computerID string, filePaths []string) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) DeleteAttachmentByIDAndComputerID(computerID, attachmentID string) error {
	return nil
}

func (m *PolicyClientAdapter) GetMobileDevices() (*jamfpro.ResponseMobileDeviceList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetMobileDeviceByID(id string) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetMobileDeviceByName(name string) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetMobileDeviceGroups() (*jamfpro.ResponseMobileDeviceGroupsList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetMobileDeviceGroupByID(id string) (interface{}, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetMobileDeviceApplications() (*jamfpro.ResponseMobileDeviceApplicationsList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetMobileDeviceConfigurationProfiles() (*jamfpro.ResponseMobileDeviceConfigurationProfilesList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) CreateMobileDevice(device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) UpdateMobileDeviceByID(id string, device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) DeleteMobileDeviceByID(id string) error {
	return nil
}

func (m *PolicyClientAdapter) GetScripts(params ...interface{}) (*jamfpro.ResponseScriptsList, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetScriptByID(id string) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) GetScriptByName(name string) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) CreateScript(script *jamfpro.ResourceScript) (*jamfpro.ResponseScriptCreate, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) UpdateScriptByID(id string, script *jamfpro.ResourceScript) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) UpdateScriptByName(name string, script *jamfpro.ResourceScript) (*jamfpro.ResourceScript, error) {
	return nil, nil
}

func (m *PolicyClientAdapter) DeleteScriptByID(id string) error {
	return nil
}

func (m *PolicyClientAdapter) DeleteScriptByName(name string) error {
	return nil
}

// TestNewPoliciesToolset tests the NewPoliciesToolset function
func TestNewPoliciesToolset(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Verify the toolset was created successfully
	assert.NotNil(t, toolset)
	assert.Equal(t, "policies", toolset.GetName())
	assert.NotEmpty(t, toolset.GetDescription())
	assert.NotEmpty(t, toolset.GetTools())

	// Verify specific tools exist
	tools := toolset.GetTools()
	toolNames := make([]string, len(tools))
	for i, tool := range tools {
		toolNames[i] = tool.Name
	}

	assert.Contains(t, toolNames, "get_policies")
	assert.Contains(t, toolNames, "get_policy_by_id")
	assert.Contains(t, toolNames, "get_policy_by_name")
	assert.Contains(t, toolNames, "create_policy")
	assert.Contains(t, toolNames, "delete_policy_by_id")
}

// TestGetPolicies tests the get_policies tool
func TestGetPolicies(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetPolicies").Return(&jamfpro.ResponsePoliciesList{
		Size: 2,
		Policy: []jamfpro.ResponsePolicyListItem{
			{
				ID:   1,
				Name: "Policy 1",
			},
			{
				ID:   2,
				Name: "Policy 2",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_policies", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Found 2 policies")
	assert.Contains(t, result, "Policy 1")
	assert.Contains(t, result, "Policy 2")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetPolicyByID tests the get_policy_by_id tool
func TestGetPolicyByID(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetPolicyByID", "1").Return(&jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			ID:      1,
			Name:    "Policy 1",
			Enabled: true,
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_policy_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Policy details for ID 1")
	assert.Contains(t, result, "Policy 1")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetPolicyByName tests the get_policy_by_name tool
func TestGetPolicyByName(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetPolicyByName", "Policy 1").Return(&jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			ID:      1,
			Name:    "Policy 1",
			Enabled: true,
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_policy_by_name", map[string]interface{}{
		"name": "Policy 1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Policy details for name 'Policy 1'")
	assert.Contains(t, result, "Policy 1")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetPoliciesByCategory tests the get_policies_by_category tool
func TestGetPoliciesByCategory(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetPolicyByCategory", "Security").Return(&jamfpro.ResponsePoliciesList{
		Size: 1,
		Policy: []jamfpro.ResponsePolicyListItem{
			{
				ID:   1,
				Name: "Security Policy",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_policies_by_category", map[string]interface{}{
		"category": "Security",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Found 1 policies in category 'Security'")
	assert.Contains(t, result, "Security Policy")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestGetPoliciesByType tests the get_policies_by_type tool
func TestGetPoliciesByType(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("GetPoliciesByType", "jss").Return(&jamfpro.ResponsePoliciesList{
		Size: 1,
		Policy: []jamfpro.ResponsePolicyListItem{
			{
				ID:   1,
				Name: "JSS Policy",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_policies_by_type", map[string]interface{}{
		"created_by": "jss",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Found 1 policies created by Jamf Pro GUI/API")
	assert.Contains(t, result, "JSS Policy")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestCreatePolicy tests the create_policy tool
func TestCreatePolicy(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client - use mock.AnythingOfType to match the struct
	mockObj.On("CreatePolicy", mock.AnythingOfType("*jamfpro.ResourcePolicy")).Return(&jamfpro.ResponsePolicyCreateAndUpdate{
		ID: 1,
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "create_policy", map[string]interface{}{
		"name":    "Test Policy",
		"enabled": true,
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully created policy 'Test Policy'")
	assert.Contains(t, result, "ID 1")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestDeletePolicyByID tests the delete_policy_by_id tool
func TestDeletePolicyByID(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("DeletePolicyByID", "1").Return(nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "delete_policy_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully deleted policy with ID 1")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestDeletePolicyByName tests the delete_policy_by_name tool
func TestDeletePolicyByName(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Set up the mock client
	mockObj.On("DeletePolicyByName", "Test Policy").Return(nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "delete_policy_by_name", map[string]interface{}{
		"name": "Test Policy",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully deleted policy with name 'Test Policy'")

	// Verify the mock was called
	mockObj.AssertExpectations(t)
}

// TestExecuteToolInvalidTool tests calling an invalid tool
func TestPoliciesExecuteToolInvalidTool(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Call an invalid tool
	_, err := toolset.ExecuteTool(context.Background(), "invalid_tool", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown tool")
}

// TestMissingRequiredArgument tests calling a tool without required arguments
func TestPoliciesMissingRequiredArgument(t *testing.T) {
	// Create a mock and adapter
	mockObj := createMockPoliciesClient()
	mockClient := NewPolicyClientAdapter(mockObj)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a policies toolset
	toolset := NewPoliciesToolset(mockClient, logger)

	// Call get_policy_by_id without the required id argument
	_, err := toolset.ExecuteTool(context.Background(), "get_policy_by_id", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")
}
