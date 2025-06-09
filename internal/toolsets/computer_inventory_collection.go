package toolsets

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"go.uber.org/zap"
)

// ComputerInventoryToolset handles computer inventory operations using Jamf Pro API
type ComputerInventoryToolset struct {
	*BaseToolset
}

// NewComputerInventoryToolset creates a new computer inventory toolset
func NewComputerInventoryToolset(client *jamfpro.Client, logger *zap.Logger) *ComputerInventoryToolset {
	base := NewBaseToolset(
		"computer-inventory",
		"Tools for managing computer inventory using the Jamf Pro API, including detailed hardware/software inventory, FileVault, and device management",
		client,
		logger,
	)

	toolset := &ComputerInventoryToolset{
		BaseToolset: base,
	}

	// Add tools based on actual Pro API capabilities
	toolset.addTools()

	return toolset
}

// addTools adds all computer inventory-related tools
func (c *ComputerInventoryToolset) addTools() {
	// Basic inventory operations
	c.addBasicInventoryTools()

	// FileVault operations
	c.addFileVaultTools()

	// Device management operations
	c.addDeviceManagementTools()

	// Attachment operations
	c.addAttachmentTools()
}

// addBasicInventoryTools adds basic computer inventory CRUD operations
func (c *ComputerInventoryToolset) addBasicInventoryTools() {
	// Get Computers Inventory
	c.AddTool(mcp.Tool{
		Name:        "get_computers_inventory",
		Description: "Retrieve computer inventory information for all computers with optional filtering, sorting, and section selection",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"page": map[string]interface{}{
					"type":        "integer",
					"description": "Page number for pagination (default: 0)",
					"minimum":     0,
				},
				"page_size": map[string]interface{}{
					"type":        "integer",
					"description": "Number of items per page (default: 100, max: 2000)",
					"minimum":     1,
					"maximum":     2000,
				},
				"sort": map[string]interface{}{
					"type":        "string",
					"description": "Sort field and direction (e.g., 'general.name:asc', 'general.lastContactTime:desc')",
				},
				"filter": map[string]interface{}{
					"type":        "string",
					"description": "Filter criteria (e.g., 'general.name==\"John's MacBook\"')",
				},
				"sections": map[string]interface{}{
					"type":        "array",
					"description": "Specific inventory sections to include. Available sections: GENERAL, DISK_ENCRYPTION, PURCHASING, APPLICATIONS, STORAGE, USER_AND_LOCATION, CONFIGURATION_PROFILES, PRINTERS, SERVICES, HARDWARE, LOCAL_USER_ACCOUNTS, CERTIFICATES, ATTACHMENTS, PLUGINS, PACKAGE_RECEIPTS, FONTS, SECURITY, OPERATING_SYSTEM, LICENSED_SOFTWARE, IBEACONS, SOFTWARE_UPDATES, EXTENSION_ATTRIBUTES, CONTENT_CACHING, GROUP_MEMBERSHIPS",
					"items": map[string]interface{}{
						"type": "string",
						"enum": []string{
							"GENERAL", "DISK_ENCRYPTION", "PURCHASING", "APPLICATIONS",
							"STORAGE", "USER_AND_LOCATION", "CONFIGURATION_PROFILES",
							"PRINTERS", "SERVICES", "HARDWARE", "LOCAL_USER_ACCOUNTS",
							"CERTIFICATES", "ATTACHMENTS", "PLUGINS", "PACKAGE_RECEIPTS",
							"FONTS", "SECURITY", "OPERATING_SYSTEM", "LICENSED_SOFTWARE",
							"IBEACONS", "SOFTWARE_UPDATES", "EXTENSION_ATTRIBUTES",
							"CONTENT_CACHING", "GROUP_MEMBERSHIPS",
						},
					},
				},
			},
			Required: []string{},
		},
	})

	// Get Computer Inventory by ID
	c.AddTool(mcp.Tool{
		Name:        "get_computer_inventory_by_id",
		Description: "Retrieve detailed inventory information for a specific computer by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to retrieve inventory for",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Computer Inventory by Name
	c.AddTool(mcp.Tool{
		Name:        "get_computer_inventory_by_name",
		Description: "Retrieve detailed inventory information for a specific computer by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the computer to retrieve inventory for",
				},
			},
			Required: []string{"name"},
		},
	})

	// Update Computer Inventory
	c.AddTool(mcp.Tool{
		Name:        "update_computer_inventory",
		Description: "Update computer inventory information using PATCH method. Only specified fields will be updated.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to update (required)",
				},
				"general": map[string]interface{}{
					"type":        "object",
					"description": "General information updates",
					"properties": map[string]interface{}{
						"name": map[string]interface{}{
							"type":        "string",
							"description": "Computer name",
						},
						"barcode1": map[string]interface{}{
							"type":        "string",
							"description": "Primary barcode",
						},
						"barcode2": map[string]interface{}{
							"type":        "string",
							"description": "Secondary barcode",
						},
						"assetTag": map[string]interface{}{
							"type":        "string",
							"description": "Asset tag",
						},
						"site": map[string]interface{}{
							"type":        "object",
							"description": "Site information",
							"properties": map[string]interface{}{
								"id": map[string]interface{}{
									"type":        "string",
									"description": "Site ID",
								},
								"name": map[string]interface{}{
									"type":        "string",
									"description": "Site name",
								},
							},
						},
					},
				},
				"userAndLocation": map[string]interface{}{
					"type":        "object",
					"description": "User and location information updates",
					"properties": map[string]interface{}{
						"username": map[string]interface{}{
							"type":        "string",
							"description": "Username",
						},
						"realname": map[string]interface{}{
							"type":        "string",
							"description": "Real name",
						},
						"email": map[string]interface{}{
							"type":        "string",
							"description": "Email address",
						},
						"position": map[string]interface{}{
							"type":        "string",
							"description": "Position/Title",
						},
						"phone": map[string]interface{}{
							"type":        "string",
							"description": "Phone number",
						},
						"departmentId": map[string]interface{}{
							"type":        "string",
							"description": "Department ID",
						},
						"buildingId": map[string]interface{}{
							"type":        "string",
							"description": "Building ID",
						},
						"room": map[string]interface{}{
							"type":        "string",
							"description": "Room",
						},
					},
				},
				"purchasing": map[string]interface{}{
					"type":        "object",
					"description": "Purchasing information updates",
					"properties": map[string]interface{}{
						"leased": map[string]interface{}{
							"type":        "boolean",
							"description": "Whether the device is leased",
						},
						"purchased": map[string]interface{}{
							"type":        "boolean",
							"description": "Whether the device was purchased",
						},
						"poNumber": map[string]interface{}{
							"type":        "string",
							"description": "Purchase order number",
						},
						"poDate": map[string]interface{}{
							"type":        "string",
							"description": "Purchase order date (YYYY-MM-DD)",
						},
						"vendor": map[string]interface{}{
							"type":        "string",
							"description": "Vendor name",
						},
						"warrantyDate": map[string]interface{}{
							"type":        "string",
							"description": "Warranty expiration date (YYYY-MM-DD)",
						},
						"appleCareId": map[string]interface{}{
							"type":        "string",
							"description": "AppleCare ID",
						},
						"leaseDate": map[string]interface{}{
							"type":        "string",
							"description": "Lease date (YYYY-MM-DD)",
						},
						"purchasePrice": map[string]interface{}{
							"type":        "string",
							"description": "Purchase price",
						},
						"lifeExpectancy": map[string]interface{}{
							"type":        "integer",
							"description": "Life expectancy in years",
						},
						"purchasingAccount": map[string]interface{}{
							"type":        "string",
							"description": "Purchasing account",
						},
						"purchasingContact": map[string]interface{}{
							"type":        "string",
							"description": "Purchasing contact",
						},
					},
				},
				"extensionAttributes": map[string]interface{}{
					"type":        "array",
					"description": "Extension attributes to update",
					"items": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"definitionId": map[string]interface{}{
								"type":        "string",
								"description": "Extension attribute definition ID",
							},
							"values": map[string]interface{}{
								"type": "array",
								"items": map[string]interface{}{
									"type": "string",
								},
								"description": "Extension attribute values",
							},
						},
						"required": []string{"definitionId", "values"},
					},
				},
			},
			Required: []string{"id"},
		},
	})

	// Delete Computer Inventory
	c.AddTool(mcp.Tool{
		Name:        "delete_computer_inventory",
		Description: "Delete a computer's inventory information by its ID (removes computer from inventory)",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to delete from inventory",
				},
			},
			Required: []string{"id"},
		},
	})
}

// addFileVaultTools adds FileVault-related tools
func (c *ComputerInventoryToolset) addFileVaultTools() {
	// Get Computers FileVault Inventory
	c.AddTool(mcp.Tool{
		Name:        "get_computers_filevault_inventory",
		Description: "Retrieve FileVault encryption information for all computers",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"page": map[string]interface{}{
					"type":        "integer",
					"description": "Page number for pagination (default: 0)",
					"minimum":     0,
				},
				"page_size": map[string]interface{}{
					"type":        "integer",
					"description": "Number of items per page (default: 100, max: 2000)",
					"minimum":     1,
					"maximum":     2000,
				},
				"sort": map[string]interface{}{
					"type":        "string",
					"description": "Sort field and direction",
				},
				"filter": map[string]interface{}{
					"type":        "string",
					"description": "Filter criteria for FileVault inventory",
				},
			},
			Required: []string{},
		},
	})

	// Get Computer FileVault Inventory by ID
	c.AddTool(mcp.Tool{
		Name:        "get_computer_filevault_inventory_by_id",
		Description: "Retrieve FileVault encryption information for a specific computer by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to retrieve FileVault information for",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Computer Recovery Lock Password
	c.AddTool(mcp.Tool{
		Name:        "get_computer_recovery_lock_password",
		Description: "Retrieve the recovery lock password for a specific computer by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to retrieve recovery lock password for",
				},
			},
			Required: []string{"id"},
		},
	})
}

// addDeviceManagementTools adds device management tools
func (c *ComputerInventoryToolset) addDeviceManagementTools() {
	// Remove Computer MDM Profile
	c.AddTool(mcp.Tool{
		Name:        "remove_computer_mdm_profile",
		Description: "Remove the MDM profile from a computer, effectively unenrolling it from management",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to remove MDM profile from",
				},
			},
			Required: []string{"id"},
		},
	})

	// Erase Computer
	c.AddTool(mcp.Tool{
		Name:        "erase_computer",
		Description: "Erase a computer by sending a remote wipe command. This will completely wipe the device.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to erase",
				},
				"pin": map[string]interface{}{
					"type":        "string",
					"description": "Optional device PIN for the erase command (required for some device types)",
				},
			},
			Required: []string{"id"},
		},
	})
}

// addAttachmentTools adds attachment management tools
func (c *ComputerInventoryToolset) addAttachmentTools() {
	// Upload Attachment to Computer
	c.AddTool(mcp.Tool{
		Name:        "upload_computer_attachment",
		Description: "Upload a file attachment to a computer. API supports single file upload only.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to upload attachment to",
				},
				"file_path": map[string]interface{}{
					"type":        "string",
					"description": "Path to the file to upload (single file only)",
				},
			},
			Required: []string{"id", "file_path"},
		},
	})

	// Delete Computer Attachment
	c.AddTool(mcp.Tool{
		Name:        "delete_computer_attachment",
		Description: "Delete a specific attachment from a computer",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"computer_id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer",
				},
				"attachment_id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the attachment to delete",
				},
			},
			Required: []string{"computer_id", "attachment_id"},
		},
	})
}

// ExecuteTool executes a computer inventory-related tool
func (c *ComputerInventoryToolset) ExecuteTool(ctx context.Context, toolName string, arguments map[string]interface{}) (string, error) {
	c.GetLogger().Debug("Executing computer inventory tool", zap.String("tool", toolName))

	switch toolName {
	// Basic inventory operations
	case "get_computers_inventory":
		return c.getComputersInventory(ctx, arguments)
	case "get_computer_inventory_by_id":
		return c.getComputerInventoryByID(ctx, arguments)
	case "get_computer_inventory_by_name":
		return c.getComputerInventoryByName(ctx, arguments)
	case "update_computer_inventory":
		return c.updateComputerInventory(ctx, arguments)
	case "delete_computer_inventory":
		return c.deleteComputerInventory(ctx, arguments)

	// FileVault operations
	case "get_computers_filevault_inventory":
		return c.getComputersFileVaultInventory(ctx, arguments)
	case "get_computer_filevault_inventory_by_id":
		return c.getComputerFileVaultInventoryByID(ctx, arguments)
	case "get_computer_recovery_lock_password":
		return c.getComputerRecoveryLockPassword(ctx, arguments)

	// Device management operations
	case "remove_computer_mdm_profile":
		return c.removeComputerMDMProfile(ctx, arguments)
	case "erase_computer":
		return c.eraseComputer(ctx, arguments)

	// Attachment operations
	case "upload_computer_attachment":
		return c.uploadComputerAttachment(ctx, arguments)
	case "delete_computer_attachment":
		return c.deleteComputerAttachment(ctx, arguments)

	default:
		return "", fmt.Errorf("unknown tool: %s", toolName)
	}
}

// Implementation of basic inventory operations

func (c *ComputerInventoryToolset) getComputersInventory(ctx context.Context, args map[string]interface{}) (string, error) {
	params := url.Values{}

	// Handle pagination
	if page, _ := GetIntArgument(args, "page", false); page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if pageSize, _ := GetIntArgument(args, "page_size", false); pageSize > 0 {
		params.Set("page-size", strconv.Itoa(pageSize))
	}

	// Handle sorting
	if sort, _ := GetStringArgument(args, "sort", false); sort != "" {
		params.Set("sort", sort)
	}

	// Handle filtering
	if filter, _ := GetStringArgument(args, "filter", false); filter != "" {
		params.Set("filter", filter)
	}

	// Handle sections
	if sections, _ := GetStringSliceArgument(args, "sections", false); len(sections) > 0 {
		params.Set("section", strings.Join(sections, ","))
	}

	inventory, err := c.GetClient().GetComputersInventory(params)
	if err != nil {
		return "", fmt.Errorf("failed to get computers inventory: %w", err)
	}

	response, err := FormatJSONResponse(inventory)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d computers in inventory:\n\n%s", inventory.TotalCount, response), nil
}

func (c *ComputerInventoryToolset) getComputerInventoryByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	inventory, err := c.GetClient().GetComputerInventoryByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get computer inventory for ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(inventory)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Computer inventory for ID %s:\n\n%s", id, response), nil
}

func (c *ComputerInventoryToolset) getComputerInventoryByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	inventory, err := c.GetClient().GetComputerInventoryByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get computer inventory for name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(inventory)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Computer inventory for name '%s':\n\n%s", name, response), nil
}

func (c *ComputerInventoryToolset) updateComputerInventory(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	// Build update object from provided arguments
	updateData := &jamfpro.ResourceComputerInventory{
		ID: id,
	}

	// Handle general updates
	if generalData, ok := args["general"].(map[string]interface{}); ok {
		if name, exists := generalData["name"].(string); exists {
			updateData.General.Name = name
		}
		if barcode1, exists := generalData["barcode1"].(string); exists {
			updateData.General.Barcode1 = barcode1
		}
		if barcode2, exists := generalData["barcode2"].(string); exists {
			updateData.General.Barcode2 = barcode2
		}
		if assetTag, exists := generalData["assetTag"].(string); exists {
			updateData.General.AssetTag = assetTag
		}
		if siteData, exists := generalData["site"].(map[string]interface{}); exists {
			if siteID, ok := siteData["id"].(string); ok {
				updateData.General.Site.ID = siteID
			}
			if siteName, ok := siteData["name"].(string); ok {
				updateData.General.Site.Name = siteName
			}
		}
	}

	// Handle user and location updates
	if userLocationData, ok := args["userAndLocation"].(map[string]interface{}); ok {
		if username, exists := userLocationData["username"].(string); exists {
			updateData.UserAndLocation.Username = username
		}
		if realname, exists := userLocationData["realname"].(string); exists {
			updateData.UserAndLocation.Realname = realname
		}
		if email, exists := userLocationData["email"].(string); exists {
			updateData.UserAndLocation.Email = email
		}
		if position, exists := userLocationData["position"].(string); exists {
			updateData.UserAndLocation.Position = position
		}
		if phone, exists := userLocationData["phone"].(string); exists {
			updateData.UserAndLocation.Phone = phone
		}
		if departmentId, exists := userLocationData["departmentId"].(string); exists {
			updateData.UserAndLocation.DepartmentId = departmentId
		}
		if buildingId, exists := userLocationData["buildingId"].(string); exists {
			updateData.UserAndLocation.BuildingId = buildingId
		}
		if room, exists := userLocationData["room"].(string); exists {
			updateData.UserAndLocation.Room = room
		}
	}

	// Handle purchasing updates
	if purchasingData, ok := args["purchasing"].(map[string]interface{}); ok {
		if leased, exists := purchasingData["leased"].(bool); exists {
			updateData.Purchasing.Leased = leased
		}
		if purchased, exists := purchasingData["purchased"].(bool); exists {
			updateData.Purchasing.Purchased = purchased
		}
		if poNumber, exists := purchasingData["poNumber"].(string); exists {
			updateData.Purchasing.PoNumber = poNumber
		}
		if poDate, exists := purchasingData["poDate"].(string); exists {
			updateData.Purchasing.PoDate = poDate
		}
		if vendor, exists := purchasingData["vendor"].(string); exists {
			updateData.Purchasing.Vendor = vendor
		}
		if warrantyDate, exists := purchasingData["warrantyDate"].(string); exists {
			updateData.Purchasing.WarrantyDate = warrantyDate
		}
		if appleCareId, exists := purchasingData["appleCareId"].(string); exists {
			updateData.Purchasing.AppleCareId = appleCareId
		}
		if leaseDate, exists := purchasingData["leaseDate"].(string); exists {
			updateData.Purchasing.LeaseDate = leaseDate
		}
		if purchasePrice, exists := purchasingData["purchasePrice"].(string); exists {
			updateData.Purchasing.PurchasePrice = purchasePrice
		}
		if lifeExpectancy, exists := purchasingData["lifeExpectancy"].(float64); exists {
			updateData.Purchasing.LifeExpectancy = int(lifeExpectancy)
		}
		if purchasingAccount, exists := purchasingData["purchasingAccount"].(string); exists {
			updateData.Purchasing.PurchasingAccount = purchasingAccount
		}
		if purchasingContact, exists := purchasingData["purchasingContact"].(string); exists {
			updateData.Purchasing.PurchasingContact = purchasingContact
		}
	}

	// Handle extension attributes
	if extAttrsData, ok := args["extensionAttributes"].([]interface{}); ok {
		for _, attrInterface := range extAttrsData {
			if attrData, ok := attrInterface.(map[string]interface{}); ok {
				attr := jamfpro.ComputerInventorySubsetExtensionAttribute{}
				if defID, exists := attrData["definitionId"].(string); exists {
					attr.DefinitionId = defID
				}
				if values, exists := attrData["values"].([]interface{}); exists {
					for _, value := range values {
						if strValue, ok := value.(string); ok {
							attr.Values = append(attr.Values, strValue)
						}
					}
				}
				updateData.ExtensionAttributes = append(updateData.ExtensionAttributes, attr)
			}
		}
	}

	result, err := c.GetClient().UpdateComputerInventoryByID(id, updateData)
	if err != nil {
		return "", fmt.Errorf("failed to update computer inventory for ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully updated computer inventory for ID %s:\n\n%s", id, response), nil
}

func (c *ComputerInventoryToolset) deleteComputerInventory(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	err = c.GetClient().DeleteComputerInventoryByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to delete computer inventory for ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully deleted computer inventory for ID %s", id), nil
}

// Implementation of FileVault operations

func (c *ComputerInventoryToolset) getComputersFileVaultInventory(ctx context.Context, args map[string]interface{}) (string, error) {
	params := url.Values{}

	// Handle pagination and filtering
	if page, _ := GetIntArgument(args, "page", false); page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if pageSize, _ := GetIntArgument(args, "page_size", false); pageSize > 0 {
		params.Set("page-size", strconv.Itoa(pageSize))
	}
	if sort, _ := GetStringArgument(args, "sort", false); sort != "" {
		params.Set("sort", sort)
	}
	if filter, _ := GetStringArgument(args, "filter", false); filter != "" {
		params.Set("filter", filter)
	}

	inventory, err := c.GetClient().GetComputersFileVaultInventory(params)
	if err != nil {
		return "", fmt.Errorf("failed to get computers FileVault inventory: %w", err)
	}

	response, err := FormatJSONResponse(inventory)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d computers with FileVault information:\n\n%s", inventory.TotalCount, response), nil
}

func (c *ComputerInventoryToolset) getComputerFileVaultInventoryByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	inventory, err := c.GetClient().GetComputerFileVaultInventoryByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get FileVault inventory for computer ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(inventory)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("FileVault inventory for computer ID %s:\n\n%s", id, response), nil
}

func (c *ComputerInventoryToolset) getComputerRecoveryLockPassword(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	password, err := c.GetClient().GetComputerRecoveryLockPasswordByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get recovery lock password for computer ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(password)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Recovery lock password for computer ID %s:\n\n%s", id, response), nil
}

// Implementation of device management operations

func (c *ComputerInventoryToolset) removeComputerMDMProfile(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	result, err := c.GetClient().RemoveComputerMDMProfile(id)
	if err != nil {
		return "", fmt.Errorf("failed to remove MDM profile for computer ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully initiated MDM profile removal for computer ID %s:\n\n%s", id, response), nil
}

func (c *ComputerInventoryToolset) eraseComputer(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	// Build erase request
	eraseRequest := jamfpro.RequestEraseDeviceComputer{}
	if pin, _ := GetStringArgument(args, "pin", false); pin != "" {
		eraseRequest.Pin = &pin
	}

	err = c.GetClient().EraseComputerByID(id, eraseRequest)
	if err != nil {
		return "", fmt.Errorf("failed to erase computer ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully initiated erase command for computer ID %s", id), nil
}

// Implementation of attachment operations

func (c *ComputerInventoryToolset) uploadComputerAttachment(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	filePath, err := GetStringArgument(args, "file_path", true)
	if err != nil {
		return "", err
	}

	result, err := c.GetClient().UploadAttachmentAndAssignToComputerByID(id, []string{filePath})
	if err != nil {
		return "", fmt.Errorf("failed to upload attachment to computer ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully uploaded attachment to computer ID %s:\n\n%s", id, response), nil
}

func (c *ComputerInventoryToolset) deleteComputerAttachment(ctx context.Context, args map[string]interface{}) (string, error) {
	computerID, err := GetStringArgument(args, "computer_id", true)
	if err != nil {
		return "", err
	}

	attachmentID, err := GetStringArgument(args, "attachment_id", true)
	if err != nil {
		return "", err
	}

	err = c.GetClient().DeleteAttachmentByIDAndComputerID(computerID, attachmentID)
	if err != nil {
		return "", fmt.Errorf("failed to delete attachment %s from computer %s: %w", attachmentID, computerID, err)
	}

	return fmt.Sprintf("Successfully deleted attachment %s from computer %s", attachmentID, computerID), nil
}
