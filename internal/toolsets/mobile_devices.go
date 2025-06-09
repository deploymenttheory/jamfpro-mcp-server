package toolsets

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"go.uber.org/zap"
)

// MobileDevicesToolset handles mobile device-related operations
type MobileDevicesToolset struct {
	*BaseToolset
}

// NewMobileDevicesToolset creates a new mobile devices toolset
func NewMobileDevicesToolset(client *jamfpro.Client, logger *zap.Logger) *MobileDevicesToolset {
	base := NewBaseToolset(
		"mobile-devices",
		"Tools for managing mobile devices in Jamf Pro, including iOS and iPadOS devices",
		client,
		logger,
	)

	toolset := &MobileDevicesToolset{
		BaseToolset: base,
	}

	toolset.addTools()
	return toolset
}

func (m *MobileDevicesToolset) addTools() {
	// Get Mobile Devices
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_devices",
		Description: "Retrieve a list of all mobile devices from Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
	})

	// Get Mobile Device by ID
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_by_id",
		Description: "Retrieve detailed information about a specific mobile device by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the mobile device to retrieve",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Mobile Device by Name
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_by_name",
		Description: "Retrieve detailed information about a specific mobile device by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the mobile device to retrieve",
				},
			},
			Required: []string{"name"},
		},
	})

	// Get Mobile Device Groups
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_groups",
		Description: "Retrieve a list of all mobile device groups from Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
	})

	// Get Mobile Device Group by ID
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_group_by_id",
		Description: "Retrieve detailed information about a specific mobile device group by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the mobile device group to retrieve",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Mobile Device Applications
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_applications",
		Description: "Retrieve a list of all mobile device applications from Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
	})

	// Get Mobile Device Configuration Profiles
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_configuration_profiles",
		Description: "Retrieve a list of all mobile device configuration profiles from Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
	})

	// Delete Mobile Device
	m.AddTool(mcp.Tool{
		Name:        "delete_mobile_device",
		Description: "Delete a mobile device from Jamf Pro by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the mobile device to delete",
				},
			},
			Required: []string{"id"},
		},
	})
}

func (m *MobileDevicesToolset) ExecuteTool(ctx context.Context, toolName string, arguments map[string]interface{}) (string, error) {
	m.GetLogger().Debug("Executing mobile devices tool", zap.String("tool", toolName))

	switch toolName {
	case "get_mobile_devices":
		return m.getMobileDevices(ctx)
	case "get_mobile_device_by_id":
		return m.getMobileDeviceByID(ctx, arguments)
	case "get_mobile_device_by_name":
		return m.getMobileDeviceByName(ctx, arguments)
	case "get_mobile_device_groups":
		return m.getMobileDeviceGroups(ctx)
	case "get_mobile_device_group_by_id":
		return m.getMobileDeviceGroupByID(ctx, arguments)
	case "get_mobile_device_applications":
		return m.getMobileDeviceApplications(ctx)
	case "get_mobile_device_configuration_profiles":
		return m.getMobileDeviceConfigurationProfiles(ctx)
	case "delete_mobile_device":
		return m.deleteMobileDevice(ctx, arguments)
	default:
		return "", fmt.Errorf("unknown tool: %s", toolName)
	}
}

func (m *MobileDevicesToolset) getMobileDevices(ctx context.Context) (string, error) {
	devices, err := m.GetClient().GetMobileDevices()
	if err != nil {
		return "", fmt.Errorf("failed to get mobile devices: %w", err)
	}

	return FormatListResponse(devices, "mobile devices")
}

func (m *MobileDevicesToolset) getMobileDeviceByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	device, err := m.GetClient().GetMobileDeviceByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(device)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Mobile device details for ID %s:\n\n%s", id, response), nil
}

func (m *MobileDevicesToolset) getMobileDeviceByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	device, err := m.GetClient().GetMobileDeviceByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device with name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(device)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Mobile device details for name '%s':\n\n%s", name, response), nil
}

func (m *MobileDevicesToolset) getMobileDeviceGroups(ctx context.Context) (string, error) {
	groups, err := m.GetClient().GetMobileDeviceGroups()
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device groups: %w", err)
	}

	return FormatListResponse(groups, "mobile device groups")
}

func (m *MobileDevicesToolset) getMobileDeviceGroupByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	group, err := m.GetClient().GetMobileDeviceGroupByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device group with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(group)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Mobile device group details for ID %s:\n\n%s", id, response), nil
}

func (m *MobileDevicesToolset) getMobileDeviceApplications(ctx context.Context) (string, error) {
	apps, err := m.GetClient().GetMobileDeviceApplications()
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device applications: %w", err)
	}

	return FormatListResponse(apps, "mobile device applications")
}

func (m *MobileDevicesToolset) getMobileDeviceConfigurationProfiles(ctx context.Context) (string, error) {
	profiles, err := m.GetClient().GetMobileDeviceConfigurationProfiles()
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device configuration profiles: %w", err)
	}

	return FormatListResponse(profiles, "mobile device configuration profiles")
}

func (m *MobileDevicesToolset) deleteMobileDevice(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	err = m.GetClient().DeleteMobileDeviceByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to delete mobile device with ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully deleted mobile device with ID %s", id), nil
}
