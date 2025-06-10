package toolsets

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// GetComputers mocks the GetComputers method - note: no parameters in actual SDK
func (m *MockJamfProClient) GetComputers() (*jamfpro.ResponseComputersList, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputersList), args.Error(1)
}

// GetComputerByID mocks the GetComputerByID method
func (m *MockJamfProClient) GetComputerByID(id string) (*jamfpro.ResponseComputer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputer), args.Error(1)
}

// GetComputerByName mocks the GetComputerByName method
func (m *MockJamfProClient) GetComputerByName(name string) (*jamfpro.ResponseComputer, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputer), args.Error(1)
}

// CreateComputer mocks the CreateComputer method - note: takes ResponseComputer, not pointer
func (m *MockJamfProClient) CreateComputer(computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	args := m.Called(computer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputer), args.Error(1)
}

// UpdateComputerByID mocks the UpdateComputerByID method
func (m *MockJamfProClient) UpdateComputerByID(id string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	args := m.Called(id, computer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputer), args.Error(1)
}

// UpdateComputerByName mocks the UpdateComputerByName method
func (m *MockJamfProClient) UpdateComputerByName(name string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error) {
	args := m.Called(name, computer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*jamfpro.ResponseComputer), args.Error(1)
}

// DeleteComputerByID mocks the DeleteComputerByID method
func (m *MockJamfProClient) DeleteComputerByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// DeleteComputerByName mocks the DeleteComputerByName method
func (m *MockJamfProClient) DeleteComputerByName(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

// Add other required methods to satisfy the jamfpro.Client interface
// These can be empty implementations for testing purposes
func (m *MockJamfProClient) GetJamfProInformation() (*jamfpro.ResponseJamfProInformation, error) {
	args := m.Called()
	return args.Get(0).(*jamfpro.ResponseJamfProInformation), args.Error(1)
}

// TestNewComputersToolset tests the NewComputersToolset function
func TestNewComputersToolset(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Verify the toolset was created successfully
	assert.NotNil(t, toolset)
	assert.Equal(t, "computers", toolset.GetName())
	assert.NotEmpty(t, toolset.GetDescription())
	assert.NotEmpty(t, toolset.GetTools())

	// Verify specific tools exist
	tools := toolset.GetTools()
	toolNames := make([]string, len(tools))
	for i, tool := range tools {
		toolNames[i] = tool.Name
	}

	assert.Contains(t, toolNames, "get_computers")
	assert.Contains(t, toolNames, "get_computer_by_id")
	assert.Contains(t, toolNames, "get_computer_by_name")
	assert.Contains(t, toolNames, "create_computer")
	assert.Contains(t, toolNames, "delete_computer_by_id")
}

// TestGetComputers tests the get_computers tool
func TestGetComputers(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client with correct struct types
	mockClient.On("GetComputers").Return(&jamfpro.ResponseComputersList{
		TotalCount: 2,
		Results: []jamfpro.ComputersListItem{
			{
				ID:   1,
				Name: "Computer 1",
			},
			{
				ID:   2,
				Name: "Computer 2",
			},
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computers", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Found 2 computers")
	assert.Contains(t, result, "Computer 1")
	assert.Contains(t, result, "Computer 2")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetComputerByID tests the get_computer_by_id tool
func TestGetComputerByID(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetComputerByID", "1").Return(&jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			ID:   1,
			Name: "Computer 1",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computer_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Computer details for ID 1")
	assert.Contains(t, result, "Computer 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetComputerByName tests the get_computer_by_name tool
func TestGetComputerByName(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("GetComputerByName", "Computer 1").Return(&jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			ID:   1,
			Name: "Computer 1",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "get_computer_by_name", map[string]interface{}{
		"name": "Computer 1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Computer details for name 'Computer 1'")
	assert.Contains(t, result, "Computer 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestCreateComputer tests the create_computer tool
func TestCreateComputer(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client - use mock.AnythingOfType to match the struct
	// Note: SDK expects ResponseComputer (not pointer) as parameter
	mockClient.On("CreateComputer", mock.AnythingOfType("jamfpro.ResponseComputer")).Return(&jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			ID:   1,
			Name: "Test Computer",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "create_computer", map[string]interface{}{
		"name": "Test Computer",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully created computer 'Test Computer'")
	assert.Contains(t, result, "ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestUpdateComputerByID tests the update_computer_by_id tool
func TestUpdateComputerByID(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Mock the initial GetComputerByID call (toolset gets current data first)
	mockClient.On("GetComputerByID", "1").Return(&jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			ID:   1,
			Name: "Original Computer",
		},
	}, nil)

	// Mock the UpdateComputerByID call
	mockClient.On("UpdateComputerByID", "1", mock.AnythingOfType("jamfpro.ResponseComputer")).Return(&jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			ID:   1,
			Name: "Updated Computer",
		},
	}, nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "update_computer_by_id", map[string]interface{}{
		"id":   "1",
		"name": "Updated Computer",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully updated computer with ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestDeleteComputerByID tests the delete_computer_by_id tool
func TestDeleteComputerByID(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("DeleteComputerByID", "1").Return(nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "delete_computer_by_id", map[string]interface{}{
		"id": "1",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully deleted computer with ID 1")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestDeleteComputerByName tests the delete_computer_by_name tool
func TestDeleteComputerByName(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("DeleteComputerByName", "Test Computer").Return(nil)

	// Call the tool via ExecuteTool
	result, err := toolset.ExecuteTool(context.Background(), "delete_computer_by_name", map[string]interface{}{
		"name": "Test Computer",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully deleted computer with name 'Test Computer'")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}

// TestGetComputerTemplate tests the get_computer_template tool
func TestGetComputerTemplate(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Call the tool via ExecuteTool (no mock needed for template)
	result, err := toolset.ExecuteTool(context.Background(), "get_computer_template", map[string]interface{}{})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Computer template:")
	assert.Contains(t, result, "Example-MacBook-Pro")
}

// TestExecuteToolInvalidTool tests calling an invalid tool
func TestExecuteToolInvalidTool(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Call an invalid tool
	_, err := toolset.ExecuteTool(context.Background(), "invalid_tool", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown tool")
}

// TestMissingRequiredArgument tests calling a tool without required arguments
func TestMissingRequiredArgument(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Call get_computer_by_id without the required id argument
	_, err := toolset.ExecuteTool(context.Background(), "get_computer_by_id", map[string]interface{}{})

	// Verify we get an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required argument")
}

// TestCreateComputerWithOptionalFields tests create_computer with optional fields
func TestCreateComputerWithOptionalFields(t *testing.T) {
	// Create a mock client
	mockClient := new(MockJamfProClient)

	// Create a logger
	logger, _ := zap.NewDevelopment()

	// Create a computers toolset
	toolset := NewComputersToolset(mockClient, logger)

	// Set up the mock client
	mockClient.On("CreateComputer", mock.AnythingOfType("jamfpro.ResponseComputer")).Return(&jamfpro.ResponseComputer{
		General: jamfpro.ComputerSubsetGeneral{
			ID:           1,
			Name:         "Test Computer",
			SerialNumber: "ABC123",
			AssetTag:     "ASSET001",
		},
		Location: jamfpro.ComputerSubsetLocation{
			Username:     "testuser",
			EmailAddress: "test@example.com",
		},
	}, nil)

	// Call the tool with optional fields
	result, err := toolset.ExecuteTool(context.Background(), "create_computer", map[string]interface{}{
		"name":          "Test Computer",
		"serial_number": "ABC123",
		"asset_tag":     "ASSET001",
		"username":      "testuser",
		"email_address": "test@example.com",
	})

	// Verify the results
	assert.NoError(t, err)
	assert.Contains(t, result, "Successfully created computer 'Test Computer'")

	// Verify the mock was called
	mockClient.AssertExpectations(t)
}
