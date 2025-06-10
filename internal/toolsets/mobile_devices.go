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
func NewMobileDevicesToolset(client JamfProClient, logger *zap.Logger) *MobileDevicesToolset {
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

	// Create Mobile Device
	m.AddTool(mcp.Tool{
		Name:        "create_mobile_device",
		Description: "Create a new mobile device in Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				// General Section
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Device name (required)",
				},
				"serial_number": map[string]interface{}{
					"type":        "string",
					"description": "Serial number of the device (required)",
				},
				"udid": map[string]interface{}{
					"type":        "string",
					"description": "UDID of the device (required)",
				},
				"asset_tag": map[string]interface{}{
					"type":        "string",
					"description": "Asset tag of the device",
				},
				"phone_number": map[string]interface{}{
					"type":        "string",
					"description": "Phone number associated with the device",
				},
				"wifi_mac_address": map[string]interface{}{
					"type":        "string",
					"description": "WiFi MAC address of the device",
				},
				"bluetooth_mac_address": map[string]interface{}{
					"type":        "string",
					"description": "Bluetooth MAC address of the device",
				},

				// Location Section
				"username": map[string]interface{}{
					"type":        "string",
					"description": "Username of the device owner",
				},
				"real_name": map[string]interface{}{
					"type":        "string",
					"description": "Real name of the device owner",
				},
				"email_address": map[string]interface{}{
					"type":        "string",
					"description": "Email address of the device owner",
				},
				"position": map[string]interface{}{
					"type":        "string",
					"description": "Position of the device owner",
				},
				"department": map[string]interface{}{
					"type":        "string",
					"description": "Department of the device owner",
				},
				"building": map[string]interface{}{
					"type":        "string",
					"description": "Building where the device is located",
				},
				"room": map[string]interface{}{
					"type":        "string",
					"description": "Room where the device is located",
				},

				// Purchasing Section
				"po_number": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order number",
				},
				"vendor": map[string]interface{}{
					"type":        "string",
					"description": "Vendor from whom the device was purchased",
				},
				"applecare_id": map[string]interface{}{
					"type":        "string",
					"description": "AppleCare ID for the device",
				},
				"purchase_price": map[string]interface{}{
					"type":        "string",
					"description": "Purchase price of the device",
				},
				"purchasing_account": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing account used",
				},
				"po_date": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order date (YYYY-MM-DD format)",
				},
				"warranty_expires": map[string]interface{}{
					"type":        "string",
					"description": "Warranty expiration date (YYYY-MM-DD format)",
				},
				"lease_expires": map[string]interface{}{
					"type":        "string",
					"description": "Lease expiration date (YYYY-MM-DD format)",
				},
				"purchasing_contact": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing contact person",
				},
				"is_purchased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device was purchased",
				},
				"is_leased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device is leased",
				},
			},
			Required: []string{"name", "serial_number", "udid"},
		},
	})

	// Update Mobile Device by ID
	m.AddTool(mcp.Tool{
		Name:        "update_mobile_device_by_id",
		Description: "Update an existing mobile device by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the mobile device to update (required)",
				},
				// General Section
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Device name",
				},
				"serial_number": map[string]interface{}{
					"type":        "string",
					"description": "Serial number of the device",
				},
				"udid": map[string]interface{}{
					"type":        "string",
					"description": "UDID of the device",
				},
				"asset_tag": map[string]interface{}{
					"type":        "string",
					"description": "Asset tag of the device",
				},
				"phone_number": map[string]interface{}{
					"type":        "string",
					"description": "Phone number associated with the device",
				},
				"wifi_mac_address": map[string]interface{}{
					"type":        "string",
					"description": "WiFi MAC address of the device",
				},
				"bluetooth_mac_address": map[string]interface{}{
					"type":        "string",
					"description": "Bluetooth MAC address of the device",
				},

				// Location Section
				"username": map[string]interface{}{
					"type":        "string",
					"description": "Username of the device owner",
				},
				"real_name": map[string]interface{}{
					"type":        "string",
					"description": "Real name of the device owner",
				},
				"email_address": map[string]interface{}{
					"type":        "string",
					"description": "Email address of the device owner",
				},
				"position": map[string]interface{}{
					"type":        "string",
					"description": "Position of the device owner",
				},
				"department": map[string]interface{}{
					"type":        "string",
					"description": "Department of the device owner",
				},
				"building": map[string]interface{}{
					"type":        "string",
					"description": "Building where the device is located",
				},
				"room": map[string]interface{}{
					"type":        "string",
					"description": "Room where the device is located",
				},

				// Purchasing Section
				"po_number": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order number",
				},
				"vendor": map[string]interface{}{
					"type":        "string",
					"description": "Vendor from whom the device was purchased",
				},
				"applecare_id": map[string]interface{}{
					"type":        "string",
					"description": "AppleCare ID for the device",
				},
				"purchase_price": map[string]interface{}{
					"type":        "string",
					"description": "Purchase price of the device",
				},
				"purchasing_account": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing account used",
				},
				"po_date": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order date (YYYY-MM-DD format)",
				},
				"warranty_expires": map[string]interface{}{
					"type":        "string",
					"description": "Warranty expiration date (YYYY-MM-DD format)",
				},
				"lease_expires": map[string]interface{}{
					"type":        "string",
					"description": "Lease expiration date (YYYY-MM-DD format)",
				},
				"purchasing_contact": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing contact person",
				},
				"is_purchased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device was purchased",
				},
				"is_leased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device is leased",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Mobile Device Template
	m.AddTool(mcp.Tool{
		Name:        "get_mobile_device_template",
		Description: "Get a reference template for a mobile device resource showing all available fields",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
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
	case "create_mobile_device":
		return m.createMobileDevice(ctx, arguments)
	case "update_mobile_device_by_id":
		return m.updateMobileDeviceByID(ctx, arguments)
	case "get_mobile_device_template":
		return m.getMobileDeviceTemplate(ctx)
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

	return fmt.Sprintf("Mobile device with ID %s has been successfully deleted", id), nil
}

func (m *MobileDevicesToolset) createMobileDevice(ctx context.Context, args map[string]interface{}) (string, error) {
	// Required fields
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	serialNumber, err := GetStringArgument(args, "serial_number", true)
	if err != nil {
		return "", err
	}

	udid, err := GetStringArgument(args, "udid", true)
	if err != nil {
		return "", err
	}

	// Create the mobile device structure
	device := &jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			DisplayName:  name,
			DeviceName:   name,
			Name:         name,
			SerialNumber: serialNumber,
			UDID:         udid,
		},
	}

	// Add optional General fields
	if assetTag, ok := args["asset_tag"].(string); ok && assetTag != "" {
		device.General.AssetTag = assetTag
	}

	if phoneNumber, ok := args["phone_number"].(string); ok && phoneNumber != "" {
		device.General.PhoneNumber = phoneNumber
	}

	if wifiMacAddress, ok := args["wifi_mac_address"].(string); ok && wifiMacAddress != "" {
		device.General.WifiMacAddress = wifiMacAddress
	}

	if bluetoothMacAddress, ok := args["bluetooth_mac_address"].(string); ok && bluetoothMacAddress != "" {
		device.General.BluetoothMacAddress = bluetoothMacAddress
	}

	// Location fields - FIXED: Initialize struct properly and set fields directly
	if username, ok := args["username"].(string); ok && username != "" {
		device.Location.Username = username
	}

	if realName, ok := args["real_name"].(string); ok && realName != "" {
		device.Location.RealName = realName
	}

	if emailAddress, ok := args["email_address"].(string); ok && emailAddress != "" {
		device.Location.EmailAddress = emailAddress
	}

	if position, ok := args["position"].(string); ok && position != "" {
		device.Location.Position = position
	}

	if department, ok := args["department"].(string); ok && department != "" {
		device.Location.Department = department
	}

	if building, ok := args["building"].(string); ok && building != "" {
		device.Location.Building = building
	}

	if roomStr, ok := args["room"].(string); ok && roomStr != "" {
		// Convert room string to int if possible
		var room int
		if _, err := fmt.Sscanf(roomStr, "%d", &room); err == nil {
			device.Location.Room = room
		}
	}

	// Purchasing fields - FIXED: Set fields directly without condition checks
	if poNumber, ok := args["po_number"].(string); ok && poNumber != "" {
		device.Purchasing.PONumber = poNumber
	}

	if vendor, ok := args["vendor"].(string); ok && vendor != "" {
		device.Purchasing.Vendor = vendor
	}

	if applecareId, ok := args["applecare_id"].(string); ok && applecareId != "" {
		device.Purchasing.ApplecareID = applecareId
	}

	if purchasePrice, ok := args["purchase_price"].(string); ok && purchasePrice != "" {
		device.Purchasing.PurchasePrice = purchasePrice
	}

	if purchasingAccount, ok := args["purchasing_account"].(string); ok && purchasingAccount != "" {
		device.Purchasing.PurchasingAccount = purchasingAccount
	}

	if poDate, ok := args["po_date"].(string); ok && poDate != "" {
		device.Purchasing.PODate = poDate
	}

	if warrantyExpires, ok := args["warranty_expires"].(string); ok && warrantyExpires != "" {
		device.Purchasing.WarrantyExpires = warrantyExpires
	}

	if leaseExpires, ok := args["lease_expires"].(string); ok && leaseExpires != "" {
		device.Purchasing.LeaseExpires = leaseExpires
	}

	if purchasingContact, ok := args["purchasing_contact"].(string); ok && purchasingContact != "" {
		device.Purchasing.PurchasingContact = purchasingContact
	}

	if isPurchased, ok := args["is_purchased"].(bool); ok {
		device.Purchasing.IsPurchased = isPurchased
	}

	if isLeased, ok := args["is_leased"].(bool); ok {
		device.Purchasing.IsLeased = isLeased
	}

	// Create the mobile device
	createdDevice, err := m.GetClient().CreateMobileDevice(device)
	if err != nil {
		return "", fmt.Errorf("failed to create mobile device: %w", err)
	}

	return fmt.Sprintf("Successfully created mobile device '%s' with ID %d", name, createdDevice.General.ID), nil
}

func (m *MobileDevicesToolset) updateMobileDeviceByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	// First, get the existing device
	existingDevice, err := m.GetClient().GetMobileDeviceByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get mobile device with ID %s: %w", id, err)
	}

	// Update General fields if provided
	if name, ok := args["name"].(string); ok && name != "" {
		existingDevice.General.DisplayName = name
		existingDevice.General.DeviceName = name
		existingDevice.General.Name = name
	}

	if assetTag, ok := args["asset_tag"].(string); ok {
		existingDevice.General.AssetTag = assetTag
	}

	if phoneNumber, ok := args["phone_number"].(string); ok {
		existingDevice.General.PhoneNumber = phoneNumber
	}

	if serialNumber, ok := args["serial_number"].(string); ok && serialNumber != "" {
		existingDevice.General.SerialNumber = serialNumber
	}

	if udid, ok := args["udid"].(string); ok && udid != "" {
		existingDevice.General.UDID = udid
	}

	if wifiMacAddress, ok := args["wifi_mac_address"].(string); ok {
		existingDevice.General.WifiMacAddress = wifiMacAddress
	}

	if bluetoothMacAddress, ok := args["bluetooth_mac_address"].(string); ok {
		existingDevice.General.BluetoothMacAddress = bluetoothMacAddress
	}

	// Update Location fields if provided
	if username, ok := args["username"].(string); ok {
		existingDevice.Location.Username = username
	}

	if realName, ok := args["real_name"].(string); ok {
		existingDevice.Location.RealName = realName
	}

	if emailAddress, ok := args["email_address"].(string); ok {
		existingDevice.Location.EmailAddress = emailAddress
	}

	if position, ok := args["position"].(string); ok {
		existingDevice.Location.Position = position
	}

	if department, ok := args["department"].(string); ok {
		existingDevice.Location.Department = department
	}

	if building, ok := args["building"].(string); ok {
		existingDevice.Location.Building = building
	}

	if roomStr, ok := args["room"].(string); ok && roomStr != "" {
		// Convert room string to int if possible
		var room int
		if _, err := fmt.Sscanf(roomStr, "%d", &room); err == nil {
			existingDevice.Location.Room = room
		}
	}

	// Update Purchasing fields if provided
	if poNumber, ok := args["po_number"].(string); ok {
		existingDevice.Purchasing.PONumber = poNumber
	}

	if vendor, ok := args["vendor"].(string); ok {
		existingDevice.Purchasing.Vendor = vendor
	}

	if applecareId, ok := args["applecare_id"].(string); ok {
		existingDevice.Purchasing.ApplecareID = applecareId
	}

	if purchasePrice, ok := args["purchase_price"].(string); ok {
		existingDevice.Purchasing.PurchasePrice = purchasePrice
	}

	if purchasingAccount, ok := args["purchasing_account"].(string); ok {
		existingDevice.Purchasing.PurchasingAccount = purchasingAccount
	}

	if poDate, ok := args["po_date"].(string); ok {
		existingDevice.Purchasing.PODate = poDate
	}

	if warrantyExpires, ok := args["warranty_expires"].(string); ok {
		existingDevice.Purchasing.WarrantyExpires = warrantyExpires
	}

	if leaseExpires, ok := args["lease_expires"].(string); ok {
		existingDevice.Purchasing.LeaseExpires = leaseExpires
	}

	if purchasingContact, ok := args["purchasing_contact"].(string); ok {
		existingDevice.Purchasing.PurchasingContact = purchasingContact
	}

	if isPurchased, ok := args["is_purchased"].(bool); ok {
		existingDevice.Purchasing.IsPurchased = isPurchased
	}

	if isLeased, ok := args["is_leased"].(bool); ok {
		existingDevice.Purchasing.IsLeased = isLeased
	}

	// Update the mobile device
	_, err = m.GetClient().UpdateMobileDeviceByID(id, existingDevice)
	if err != nil {
		return "", fmt.Errorf("failed to update mobile device with ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully updated mobile device with ID %s", id), nil
}

// GetMobileDeviceTemplate returns an example template of a mobile device resource
func (m *MobileDevicesToolset) GetMobileDeviceTemplate() *jamfpro.ResourceMobileDevice {
	return &jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			DisplayName:         "Example iPad",
			DeviceName:          "Example iPad",
			Name:                "Example iPad",
			AssetTag:            "ASSET123",
			SerialNumber:        "C02Q7KHTGFW2",
			UDID:                "270aae10800b6e61a2ee2bbc285eb967050b6112",
			PhoneNumber:         "+1 (555) 123-4567",
			WifiMacAddress:      "AA:BB:CC:DD:EE:FF",
			BluetoothMacAddress: "AA:BB:CC:DD:EE:00",
			IPAddress:           "192.168.1.100",
			Model:               "iPad Pro",
			ModelIdentifier:     "iPad11,4",
			ModelDisplay:        "iPad Pro 12.9-inch (3rd generation)",
		},
		Location: jamfpro.MobileDeviceSubsetLocation{
			Username:     "jdoe",
			RealName:     "John Doe",
			EmailAddress: "john.doe@example.com",
			Position:     "Developer",
			Department:   "Engineering",
			Building:     "Main Campus",
			Room:         101,
		},
		Purchasing: jamfpro.MobileDeviceSubsetPurchasing{
			IsPurchased:       true,
			IsLeased:          false,
			PONumber:          "PO-12345",
			Vendor:            "Apple",
			ApplecareID:       "AC-987654",
			PurchasePrice:     "$999.00",
			PurchasingAccount: "IT-DEPT",
			PODate:            "2023-01-15",
			WarrantyExpires:   "2025-01-15",
			LeaseExpires:      "",
			PurchasingContact: "procurement@example.com",
		},
	}
}

func (m *MobileDevicesToolset) getMobileDeviceTemplate(ctx context.Context) (string, error) {
	template := m.GetMobileDeviceTemplate()
	response, err := FormatJSONResponse(template)
	if err != nil {
		return "", fmt.Errorf("failed to format mobile device template: %w", err)
	}
	return fmt.Sprintf("Mobile device template:\n\n%s", response), nil
}
