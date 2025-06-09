package toolsets

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"go.uber.org/zap"
)

// ComputersToolset handles computer-related operations using Jamf Pro Classic API
type ComputersToolset struct {
	*BaseToolset
}

// NewComputersToolset creates a new computers toolset
func NewComputersToolset(client *jamfpro.Client, logger *zap.Logger) *ComputersToolset {
	base := NewBaseToolset(
		"computers",
		"Tools for managing computers in Jamf Pro using the Classic API, including CRUD operations and detailed computer information",
		client,
		logger,
	)

	toolset := &ComputersToolset{
		BaseToolset: base,
	}

	// Add tools based on actual SDK capabilities
	toolset.addTools()

	return toolset
}

// addTools adds all computer-related tools based on the actual Classic API SDK
func (c *ComputersToolset) addTools() {
	// Get Computers List
	c.AddTool(mcp.Tool{
		Name:        "get_computers",
		Description: "Retrieve a list of all computers from Jamf Pro (returns basic info: ID and name only)",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
	})

	// Get Computer by ID (Full Details)
	c.AddTool(mcp.Tool{
		Name:        "get_computer_by_id",
		Description: "Retrieve complete detailed information about a specific computer by its ID (includes all sections: general, location, purchasing, hardware, software, etc.)",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to retrieve",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Computer by Name (Full Details)
	c.AddTool(mcp.Tool{
		Name:        "get_computer_by_name",
		Description: "Retrieve complete detailed information about a specific computer by its name (includes all sections: general, location, purchasing, hardware, software, etc.)",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the computer to retrieve",
				},
			},
			Required: []string{"name"},
		},
	})

	// Create Computer
	c.AddTool(mcp.Tool{
		Name:        "create_computer",
		Description: "Create a new computer record in Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Computer name (required)",
				},
				"serial_number": map[string]interface{}{
					"type":        "string",
					"description": "Serial number of the computer",
				},
				"udid": map[string]interface{}{
					"type":        "string",
					"description": "UDID of the computer",
				},
				"mac_address": map[string]interface{}{
					"type":        "string",
					"description": "MAC address of the computer",
				},
				"asset_tag": map[string]interface{}{
					"type":        "string",
					"description": "Asset tag for the computer",
				},
				"barcode_1": map[string]interface{}{
					"type":        "string",
					"description": "Primary barcode",
				},
				"barcode_2": map[string]interface{}{
					"type":        "string",
					"description": "Secondary barcode",
				},
				"site_id": map[string]interface{}{
					"type":        "integer",
					"description": "Site ID for the computer (-1 for none)",
				},
				"site_name": map[string]interface{}{
					"type":        "string",
					"description": "Site name for the computer",
				},
				// Location information
				"username": map[string]interface{}{
					"type":        "string",
					"description": "Username",
				},
				"real_name": map[string]interface{}{
					"type":        "string",
					"description": "Real name of the user",
				},
				"email_address": map[string]interface{}{
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
				"phone_number": map[string]interface{}{
					"type":        "string",
					"description": "Phone number (alternative field)",
				},
				"department": map[string]interface{}{
					"type":        "string",
					"description": "Department name",
				},
				"building": map[string]interface{}{
					"type":        "string",
					"description": "Building name",
				},
				"room": map[string]interface{}{
					"type":        "string",
					"description": "Room",
				},
				// Purchasing information
				"is_purchased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device was purchased",
				},
				"is_leased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device is leased",
				},
				"po_number": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order number",
				},
				"vendor": map[string]interface{}{
					"type":        "string",
					"description": "Vendor name",
				},
				"applecare_id": map[string]interface{}{
					"type":        "string",
					"description": "AppleCare ID",
				},
				"purchase_price": map[string]interface{}{
					"type":        "string",
					"description": "Purchase price",
				},
				"purchasing_account": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing account",
				},
				"purchasing_contact": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing contact",
				},
				"po_date": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order date",
				},
				"warranty_expires": map[string]interface{}{
					"type":        "string",
					"description": "Warranty expiration date",
				},
				"lease_expires": map[string]interface{}{
					"type":        "string",
					"description": "Lease expiration date",
				},
				"life_expectancy": map[string]interface{}{
					"type":        "integer",
					"description": "Life expectancy in years",
				},
			},
			Required: []string{"name"},
		},
	})

	// Update Computer by ID
	c.AddTool(mcp.Tool{
		Name:        "update_computer_by_id",
		Description: "Update computer information by ID. Can update general info, location, purchasing, and other details",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to update (required)",
				},
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Computer name",
				},
				"asset_tag": map[string]interface{}{
					"type":        "string",
					"description": "Asset tag",
				},
				"barcode_1": map[string]interface{}{
					"type":        "string",
					"description": "Primary barcode",
				},
				"barcode_2": map[string]interface{}{
					"type":        "string",
					"description": "Secondary barcode",
				},
				"site_id": map[string]interface{}{
					"type":        "integer",
					"description": "Site ID (-1 for none)",
				},
				"site_name": map[string]interface{}{
					"type":        "string",
					"description": "Site name",
				},
				// Location fields
				"username": map[string]interface{}{
					"type":        "string",
					"description": "Username",
				},
				"real_name": map[string]interface{}{
					"type":        "string",
					"description": "Real name of the user",
				},
				"email_address": map[string]interface{}{
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
				"phone_number": map[string]interface{}{
					"type":        "string",
					"description": "Phone number (alternative field)",
				},
				"department": map[string]interface{}{
					"type":        "string",
					"description": "Department name",
				},
				"building": map[string]interface{}{
					"type":        "string",
					"description": "Building name",
				},
				"room": map[string]interface{}{
					"type":        "string",
					"description": "Room",
				},
				// Purchasing fields
				"is_purchased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device was purchased",
				},
				"is_leased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device is leased",
				},
				"po_number": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order number",
				},
				"vendor": map[string]interface{}{
					"type":        "string",
					"description": "Vendor name",
				},
				"applecare_id": map[string]interface{}{
					"type":        "string",
					"description": "AppleCare ID",
				},
				"purchase_price": map[string]interface{}{
					"type":        "string",
					"description": "Purchase price",
				},
				"purchasing_account": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing account",
				},
				"purchasing_contact": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing contact",
				},
				"po_date": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order date",
				},
				"warranty_expires": map[string]interface{}{
					"type":        "string",
					"description": "Warranty expiration date",
				},
				"lease_expires": map[string]interface{}{
					"type":        "string",
					"description": "Lease expiration date",
				},
				"life_expectancy": map[string]interface{}{
					"type":        "integer",
					"description": "Life expectancy in years",
				},
			},
			Required: []string{"id"},
		},
	})

	// Update Computer by Name
	c.AddTool(mcp.Tool{
		Name:        "update_computer_by_name",
		Description: "Update computer information by name. Can update general info, location, purchasing, and other details",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the computer to update (required)",
				},
				"new_name": map[string]interface{}{
					"type":        "string",
					"description": "New computer name (if changing the name)",
				},
				"asset_tag": map[string]interface{}{
					"type":        "string",
					"description": "Asset tag",
				},
				"barcode_1": map[string]interface{}{
					"type":        "string",
					"description": "Primary barcode",
				},
				"barcode_2": map[string]interface{}{
					"type":        "string",
					"description": "Secondary barcode",
				},
				"site_id": map[string]interface{}{
					"type":        "integer",
					"description": "Site ID (-1 for none)",
				},
				"site_name": map[string]interface{}{
					"type":        "string",
					"description": "Site name",
				},
				// Location fields (same as update by ID)
				"username": map[string]interface{}{
					"type":        "string",
					"description": "Username",
				},
				"real_name": map[string]interface{}{
					"type":        "string",
					"description": "Real name of the user",
				},
				"email_address": map[string]interface{}{
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
				"phone_number": map[string]interface{}{
					"type":        "string",
					"description": "Phone number (alternative field)",
				},
				"department": map[string]interface{}{
					"type":        "string",
					"description": "Department name",
				},
				"building": map[string]interface{}{
					"type":        "string",
					"description": "Building name",
				},
				"room": map[string]interface{}{
					"type":        "string",
					"description": "Room",
				},
				// Purchasing fields (same as update by ID)
				"is_purchased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device was purchased",
				},
				"is_leased": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the device is leased",
				},
				"po_number": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order number",
				},
				"vendor": map[string]interface{}{
					"type":        "string",
					"description": "Vendor name",
				},
				"applecare_id": map[string]interface{}{
					"type":        "string",
					"description": "AppleCare ID",
				},
				"purchase_price": map[string]interface{}{
					"type":        "string",
					"description": "Purchase price",
				},
				"purchasing_account": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing account",
				},
				"purchasing_contact": map[string]interface{}{
					"type":        "string",
					"description": "Purchasing contact",
				},
				"po_date": map[string]interface{}{
					"type":        "string",
					"description": "Purchase order date",
				},
				"warranty_expires": map[string]interface{}{
					"type":        "string",
					"description": "Warranty expiration date",
				},
				"lease_expires": map[string]interface{}{
					"type":        "string",
					"description": "Lease expiration date",
				},
				"life_expectancy": map[string]interface{}{
					"type":        "integer",
					"description": "Life expectancy in years",
				},
			},
			Required: []string{"name"},
		},
	})

	// Delete Computer by ID
	c.AddTool(mcp.Tool{
		Name:        "delete_computer_by_id",
		Description: "Delete a computer from Jamf Pro by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the computer to delete",
				},
			},
			Required: []string{"id"},
		},
	})

	// Delete Computer by Name
	c.AddTool(mcp.Tool{
		Name:        "delete_computer_by_name",
		Description: "Delete a computer from Jamf Pro by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the computer to delete",
				},
			},
			Required: []string{"name"},
		},
	})
}

// ExecuteTool executes a computer-related tool
func (c *ComputersToolset) ExecuteTool(ctx context.Context, toolName string, arguments map[string]interface{}) (string, error) {
	c.GetLogger().Debug("Executing computers tool", zap.String("tool", toolName))

	switch toolName {
	case "get_computers":
		return c.getComputers(ctx)
	case "get_computer_by_id":
		return c.getComputerByID(ctx, arguments)
	case "get_computer_by_name":
		return c.getComputerByName(ctx, arguments)
	case "create_computer":
		return c.createComputer(ctx, arguments)
	case "update_computer_by_id":
		return c.updateComputerByID(ctx, arguments)
	case "update_computer_by_name":
		return c.updateComputerByName(ctx, arguments)
	case "delete_computer_by_id":
		return c.deleteComputerByID(ctx, arguments)
	case "delete_computer_by_name":
		return c.deleteComputerByName(ctx, arguments)
	default:
		return "", fmt.Errorf("unknown tool: %s", toolName)
	}
}

// Implementation of computer operations based on actual SDK methods

// getComputers retrieves all computers (basic list)
func (c *ComputersToolset) getComputers(ctx context.Context) (string, error) {
	computers, err := c.GetClient().GetComputers()
	if err != nil {
		return "", fmt.Errorf("failed to get computers: %w", err)
	}

	response, err := FormatJSONResponse(computers)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d computers:\n\n%s", computers.TotalCount, response), nil
}

// getComputerByID retrieves a computer by ID with full details
func (c *ComputersToolset) getComputerByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	computer, err := c.GetClient().GetComputerByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get computer with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(computer)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Computer details for ID %s:\n\n%s", id, response), nil
}

// getComputerByName retrieves a computer by name with full details
func (c *ComputersToolset) getComputerByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	computer, err := c.GetClient().GetComputerByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get computer with name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(computer)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Computer details for name '%s':\n\n%s", name, response), nil
}

// createComputer creates a new computer record
func (c *ComputersToolset) createComputer(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	// Build the computer object
	computer := jamfpro.ResponseComputer{}

	// General information
	computer.General.Name = name

	if serialNumber, _ := GetStringArgument(args, "serial_number", false); serialNumber != "" {
		computer.General.SerialNumber = serialNumber
	}
	if udid, _ := GetStringArgument(args, "udid", false); udid != "" {
		computer.General.UDID = udid
	}
	if macAddress, _ := GetStringArgument(args, "mac_address", false); macAddress != "" {
		computer.General.MacAddress = macAddress
	}
	if assetTag, _ := GetStringArgument(args, "asset_tag", false); assetTag != "" {
		computer.General.AssetTag = assetTag
	}
	if barcode1, _ := GetStringArgument(args, "barcode_1", false); barcode1 != "" {
		computer.General.Barcode1 = barcode1
	}
	if barcode2, _ := GetStringArgument(args, "barcode_2", false); barcode2 != "" {
		computer.General.Barcode2 = barcode2
	}

	// Site information
	if siteID, _ := GetIntArgument(args, "site_id", false); siteID != 0 {
		computer.General.Site.ID = siteID
	}
	if siteName, _ := GetStringArgument(args, "site_name", false); siteName != "" {
		computer.General.Site.Name = siteName
	}

	// Location information
	if username, _ := GetStringArgument(args, "username", false); username != "" {
		computer.Location.Username = username
	}
	if realName, _ := GetStringArgument(args, "real_name", false); realName != "" {
		computer.Location.RealName = realName
	}
	if email, _ := GetStringArgument(args, "email_address", false); email != "" {
		computer.Location.EmailAddress = email
	}
	if position, _ := GetStringArgument(args, "position", false); position != "" {
		computer.Location.Position = position
	}
	if phone, _ := GetStringArgument(args, "phone", false); phone != "" {
		computer.Location.Phone = phone
	}
	if phoneNumber, _ := GetStringArgument(args, "phone_number", false); phoneNumber != "" {
		computer.Location.PhoneNumber = phoneNumber
	}
	if department, _ := GetStringArgument(args, "department", false); department != "" {
		computer.Location.Department = department
	}
	if building, _ := GetStringArgument(args, "building", false); building != "" {
		computer.Location.Building = building
	}
	if room, _ := GetStringArgument(args, "room", false); room != "" {
		computer.Location.Room = room
	}

	// Purchasing information
	if isPurchased, _ := GetBoolArgument(args, "is_purchased", false); isPurchased {
		computer.Purchasing.IsPurchased = isPurchased
	}
	if isLeased, _ := GetBoolArgument(args, "is_leased", false); isLeased {
		computer.Purchasing.IsLeased = isLeased
	}
	if poNumber, _ := GetStringArgument(args, "po_number", false); poNumber != "" {
		computer.Purchasing.PoNumber = poNumber
	}
	if vendor, _ := GetStringArgument(args, "vendor", false); vendor != "" {
		computer.Purchasing.Vendor = vendor
	}
	if applecareID, _ := GetStringArgument(args, "applecare_id", false); applecareID != "" {
		computer.Purchasing.ApplecareID = applecareID
	}
	if purchasePrice, _ := GetStringArgument(args, "purchase_price", false); purchasePrice != "" {
		computer.Purchasing.PurchasePrice = purchasePrice
	}
	if purchasingAccount, _ := GetStringArgument(args, "purchasing_account", false); purchasingAccount != "" {
		computer.Purchasing.PurchasingAccount = purchasingAccount
	}
	if purchasingContact, _ := GetStringArgument(args, "purchasing_contact", false); purchasingContact != "" {
		computer.Purchasing.PurchasingContact = purchasingContact
	}
	if poDate, _ := GetStringArgument(args, "po_date", false); poDate != "" {
		computer.Purchasing.PoDate = poDate
	}
	if warrantyExpires, _ := GetStringArgument(args, "warranty_expires", false); warrantyExpires != "" {
		computer.Purchasing.WarrantyExpires = warrantyExpires
	}
	if leaseExpires, _ := GetStringArgument(args, "lease_expires", false); leaseExpires != "" {
		computer.Purchasing.LeaseExpires = leaseExpires
	}
	if lifeExpectancy, _ := GetIntArgument(args, "life_expectancy", false); lifeExpectancy != 0 {
		computer.Purchasing.LifeExpectancy = lifeExpectancy
	}

	result, err := c.GetClient().CreateComputer(computer)
	if err != nil {
		return "", fmt.Errorf("failed to create computer '%s': %w", name, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully created computer '%s' with ID %d:\n\n%s", name, result.General.ID, response), nil
}

// updateComputerByID updates a computer by ID
func (c *ComputersToolset) updateComputerByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	// First get the current computer to preserve existing data
	currentComputer, err := c.GetClient().GetComputerByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get current computer data for ID %s: %w", id, err)
	}

	// Update fields that were provided
	computer := *currentComputer

	// General information updates
	if name, _ := GetStringArgument(args, "name", false); name != "" {
		computer.General.Name = name
	}
	if assetTag, _ := GetStringArgument(args, "asset_tag", false); assetTag != "" {
		computer.General.AssetTag = assetTag
	}
	if barcode1, _ := GetStringArgument(args, "barcode_1", false); barcode1 != "" {
		computer.General.Barcode1 = barcode1
	}
	if barcode2, _ := GetStringArgument(args, "barcode_2", false); barcode2 != "" {
		computer.General.Barcode2 = barcode2
	}

	// Site information updates
	if siteID, _ := GetIntArgument(args, "site_id", false); siteID != 0 {
		computer.General.Site.ID = siteID
	}
	if siteName, _ := GetStringArgument(args, "site_name", false); siteName != "" {
		computer.General.Site.Name = siteName
	}

	// Location information updates
	if username, _ := GetStringArgument(args, "username", false); username != "" {
		computer.Location.Username = username
	}
	if realName, _ := GetStringArgument(args, "real_name", false); realName != "" {
		computer.Location.RealName = realName
	}
	if email, _ := GetStringArgument(args, "email_address", false); email != "" {
		computer.Location.EmailAddress = email
	}
	if position, _ := GetStringArgument(args, "position", false); position != "" {
		computer.Location.Position = position
	}
	if phone, _ := GetStringArgument(args, "phone", false); phone != "" {
		computer.Location.Phone = phone
	}
	if phoneNumber, _ := GetStringArgument(args, "phone_number", false); phoneNumber != "" {
		computer.Location.PhoneNumber = phoneNumber
	}
	if department, _ := GetStringArgument(args, "department", false); department != "" {
		computer.Location.Department = department
	}
	if building, _ := GetStringArgument(args, "building", false); building != "" {
		computer.Location.Building = building
	}
	if room, _ := GetStringArgument(args, "room", false); room != "" {
		computer.Location.Room = room
	}

	// Purchasing information updates
	if _, ok := args["is_purchased"]; ok {
		if purchasedBool, _ := GetBoolArgument(args, "is_purchased", false); ok {
			computer.Purchasing.IsPurchased = purchasedBool
		}
	}
	if isLeased, ok := args["is_leased"]; ok {
		if leasedBool, _ := GetBoolArgument(args, "is_leased", false); ok {
			computer.Purchasing.IsLeased = leasedBool
		}
	}
	if poNumber, _ := GetStringArgument(args, "po_number", false); poNumber != "" {
		computer.Purchasing.PoNumber = poNumber
	}
	if vendor, _ := GetStringArgument(args, "vendor", false); vendor != "" {
		computer.Purchasing.Vendor = vendor
	}
	if applecareID, _ := GetStringArgument(args, "applecare_id", false); applecareID != "" {
		computer.Purchasing.ApplecareID = applecareID
	}
	if purchasePrice, _ := GetStringArgument(args, "purchase_price", false); purchasePrice != "" {
		computer.Purchasing.PurchasePrice = purchasePrice
	}
	if purchasingAccount, _ := GetStringArgument(args, "purchasing_account", false); purchasingAccount != "" {
		computer.Purchasing.PurchasingAccount = purchasingAccount
	}
	if purchasingContact, _ := GetStringArgument(args, "purchasing_contact", false); purchasingContact != "" {
		computer.Purchasing.PurchasingContact = purchasingContact
	}
	if poDate, _ := GetStringArgument(args, "po_date", false); poDate != "" {
		computer.Purchasing.PoDate = poDate
	}
	if warrantyExpires, _ := GetStringArgument(args, "warranty_expires", false); warrantyExpires != "" {
		computer.Purchasing.WarrantyExpires = warrantyExpires
	}
	if leaseExpires, _ := GetStringArgument(args, "lease_expires", false); leaseExpires != "" {
		computer.Purchasing.LeaseExpires = leaseExpires
	}
	if lifeExpectancy, _ := GetIntArgument(args, "life_expectancy", false); lifeExpectancy != 0 {
		computer.Purchasing.LifeExpectancy = lifeExpectancy
	}

	result, err := c.GetClient().UpdateComputerByID(id, computer)
	if err != nil {
		return "", fmt.Errorf("failed to update computer with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully updated computer with ID %s:\n\n%s", id, response), nil
}

// updateComputerByName updates a computer by name
func (c *ComputersToolset) updateComputerByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	// First get the current computer to preserve existing data
	currentComputer, err := c.GetClient().GetComputerByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get current computer data for name %s: %w", name, err)
	}

	// Update fields that were provided
	computer := *currentComputer

	// General information updates
	if newName, _ := GetStringArgument(args, "new_name", false); newName != "" {
		computer.General.Name = newName
	}
	if assetTag, _ := GetStringArgument(args, "asset_tag", false); assetTag != "" {
		computer.General.AssetTag = assetTag
	}
	if barcode1, _ := GetStringArgument(args, "barcode_1", false); barcode1 != "" {
		computer.General.Barcode1 = barcode1
	}
	if barcode2, _ := GetStringArgument(args, "barcode_2", false); barcode2 != "" {
		computer.General.Barcode2 = barcode2
	}

	// Site information updates
	if siteID, _ := GetIntArgument(args, "site_id", false); siteID != 0 {
		computer.General.Site.ID = siteID
	}
	if siteName, _ := GetStringArgument(args, "site_name", false); siteName != "" {
		computer.General.Site.Name = siteName
	}

	// Location information updates (same logic as updateByID)
	if username, _ := GetStringArgument(args, "username", false); username != "" {
		computer.Location.Username = username
	}
	if realName, _ := GetStringArgument(args, "real_name", false); realName != "" {
		computer.Location.RealName = realName
	}
	if email, _ := GetStringArgument(args, "email_address", false); email != "" {
		computer.Location.EmailAddress = email
	}
	if position, _ := GetStringArgument(args, "position", false); position != "" {
		computer.Location.Position = position
	}
	if phone, _ := GetStringArgument(args, "phone", false); phone != "" {
		computer.Location.Phone = phone
	}
	if phoneNumber, _ := GetStringArgument(args, "phone_number", false); phoneNumber != "" {
		computer.Location.PhoneNumber = phoneNumber
	}
	if department, _ := GetStringArgument(args, "department", false); department != "" {
		computer.Location.Department = department
	}
	if building, _ := GetStringArgument(args, "building", false); building != "" {
		computer.Location.Building = building
	}
	if room, _ := GetStringArgument(args, "room", false); room != "" {
		computer.Location.Room = room
	}

	// Purchasing information updates (same logic as updateByID)
	if isPurchased, ok := args["is_purchased"]; ok {
		if purchasedBool, _ := GetBoolArgument(args, "is_purchased", false); ok {
			computer.Purchasing.IsPurchased = purchasedBool
		}
	}
	if isLeased, ok := args["is_leased"]; ok {
		if leasedBool, _ := GetBoolArgument(args, "is_leased", false); ok {
			computer.Purchasing.IsLeased = leasedBool
		}
	}
	if poNumber, _ := GetStringArgument(args, "po_number", false); poNumber != "" {
		computer.Purchasing.PoNumber = poNumber
	}
	if vendor, _ := GetStringArgument(args, "vendor", false); vendor != "" {
		computer.Purchasing.Vendor = vendor
	}
	if applecareID, _ := GetStringArgument(args, "applecare_id", false); applecareID != "" {
		computer.Purchasing.ApplecareID = applecareID
	}
	if purchasePrice, _ := GetStringArgument(args, "purchase_price", false); purchasePrice != "" {
		computer.Purchasing.PurchasePrice = purchasePrice
	}
	if purchasingAccount, _ := GetStringArgument(args, "purchasing_account", false); purchasingAccount != "" {
		computer.Purchasing.PurchasingAccount = purchasingAccount
	}
	if purchasingContact, _ := GetStringArgument(args, "purchasing_contact", false); purchasingContact != "" {
		computer.Purchasing.PurchasingContact = purchasingContact
	}
	if poDate, _ := GetStringArgument(args, "po_date", false); poDate != "" {
		computer.Purchasing.PoDate = poDate
	}
	if warrantyExpires, _ := GetStringArgument(args, "warranty_expires", false); warrantyExpires != "" {
		computer.Purchasing.WarrantyExpires = warrantyExpires
	}
	if leaseExpires, _ := GetStringArgument(args, "lease_expires", false); leaseExpires != "" {
		computer.Purchasing.LeaseExpires = leaseExpires
	}
	if lifeExpectancy, _ := GetIntArgument(args, "life_expectancy", false); lifeExpectancy != 0 {
		computer.Purchasing.LifeExpectancy = lifeExpectancy
	}

	result, err := c.GetClient().UpdateComputerByName(name, computer)
	if err != nil {
		return "", fmt.Errorf("failed to update computer with name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully updated computer with name '%s':\n\n%s", name, response), nil
}

// deleteComputerByID deletes a computer by ID
func (c *ComputersToolset) deleteComputerByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	err = c.GetClient().DeleteComputerByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to delete computer with ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully deleted computer with ID %s", id), nil
}

// deleteComputerByName deletes a computer by name
func (c *ComputersToolset) deleteComputerByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	err = c.GetClient().DeleteComputerByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to delete computer with name %s: %w", name, err)
	}

	return fmt.Sprintf("Successfully deleted computer with name '%s'", name), nil
}
