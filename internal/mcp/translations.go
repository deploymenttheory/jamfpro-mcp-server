package mcp

import (
	"encoding/json"
	"os"
)

// GetDefaultTranslations returns the default tool descriptions
func GetDefaultTranslations() map[string]string {
	return map[string]string{
		// Computers toolset
		"TOOL_GET_COMPUTERS_DESCRIPTION":                "Retrieve a list of all computers from Jamf Pro",
		"TOOL_GET_COMPUTER_BY_ID_DESCRIPTION":           "Retrieve detailed information about a specific computer by its ID",
		"TOOL_GET_COMPUTER_BY_NAME_DESCRIPTION":         "Retrieve detailed information about a specific computer by its name",
		"TOOL_GET_COMPUTERS_INVENTORY_DESCRIPTION":      "Retrieve computer inventory information with optional filtering and pagination",
		"TOOL_GET_COMPUTER_INVENTORY_BY_ID_DESCRIPTION": "Retrieve detailed inventory information for a specific computer by its ID",
		"TOOL_UPDATE_COMPUTER_INVENTORY_DESCRIPTION":    "Update computer inventory information for a specific computer",
		"TOOL_GET_COMPUTER_GROUPS_DESCRIPTION":          "Retrieve a list of all computer groups from Jamf Pro",
		"TOOL_GET_COMPUTER_GROUP_BY_ID_DESCRIPTION":     "Retrieve detailed information about a specific computer group by its ID",
		"TOOL_DELETE_COMPUTER_DESCRIPTION":              "Delete a computer from Jamf Pro by its ID",

		// Mobile devices toolset
		"TOOL_GET_MOBILE_DEVICES_DESCRIPTION":            "Retrieve a list of all mobile devices from Jamf Pro",
		"TOOL_GET_MOBILE_DEVICE_BY_ID_DESCRIPTION":       "Retrieve detailed information about a specific mobile device by its ID",
		"TOOL_GET_MOBILE_DEVICE_BY_NAME_DESCRIPTION":     "Retrieve detailed information about a specific mobile device by its name",
		"TOOL_GET_MOBILE_DEVICE_GROUPS_DESCRIPTION":      "Retrieve a list of all mobile device groups from Jamf Pro",
		"TOOL_GET_MOBILE_DEVICE_GROUP_BY_ID_DESCRIPTION": "Retrieve detailed information about a specific mobile device group by its ID",
		"TOOL_DELETE_MOBILE_DEVICE_DESCRIPTION":          "Delete a mobile device from Jamf Pro by its ID",

		// Policies toolset
		"TOOL_GET_POLICIES_DESCRIPTION":       "Retrieve a list of all policies from Jamf Pro",
		"TOOL_GET_POLICY_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific policy by its ID",
		"TOOL_GET_POLICY_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific policy by its name",
		"TOOL_CREATE_POLICY_DESCRIPTION":      "Create a new policy in Jamf Pro",
		"TOOL_UPDATE_POLICY_DESCRIPTION":      "Update an existing policy in Jamf Pro",
		"TOOL_DELETE_POLICY_DESCRIPTION":      "Delete a policy from Jamf Pro by its ID",

		// Users toolset
		"TOOL_GET_USERS_DESCRIPTION":        "Retrieve a list of all users from Jamf Pro",
		"TOOL_GET_USER_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific user by their ID",
		"TOOL_GET_USER_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific user by their name",
		"TOOL_CREATE_USER_DESCRIPTION":      "Create a new user in Jamf Pro",
		"TOOL_UPDATE_USER_DESCRIPTION":      "Update an existing user in Jamf Pro",
		"TOOL_DELETE_USER_DESCRIPTION":      "Delete a user from Jamf Pro by their ID",

		// Groups toolset
		"TOOL_GET_USER_GROUPS_DESCRIPTION":        "Retrieve a list of all user groups from Jamf Pro",
		"TOOL_GET_USER_GROUP_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific user group by its ID",
		"TOOL_GET_USER_GROUP_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific user group by its name",
		"TOOL_CREATE_USER_GROUP_DESCRIPTION":      "Create a new user group in Jamf Pro",
		"TOOL_UPDATE_USER_GROUP_DESCRIPTION":      "Update an existing user group in Jamf Pro",
		"TOOL_DELETE_USER_GROUP_DESCRIPTION":      "Delete a user group from Jamf Pro by its ID",

		// Configuration profiles toolset
		"TOOL_GET_CONFIGURATION_PROFILES_DESCRIPTION":        "Retrieve a list of all configuration profiles from Jamf Pro",
		"TOOL_GET_CONFIGURATION_PROFILE_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific configuration profile by its ID",
		"TOOL_GET_CONFIGURATION_PROFILE_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific configuration profile by its name",
		"TOOL_CREATE_CONFIGURATION_PROFILE_DESCRIPTION":      "Create a new configuration profile in Jamf Pro",
		"TOOL_UPDATE_CONFIGURATION_PROFILE_DESCRIPTION":      "Update an existing configuration profile in Jamf Pro",
		"TOOL_DELETE_CONFIGURATION_PROFILE_DESCRIPTION":      "Delete a configuration profile from Jamf Pro by its ID",

		// Scripts toolset
		"TOOL_GET_SCRIPTS_DESCRIPTION":        "Retrieve a list of all scripts from Jamf Pro",
		"TOOL_GET_SCRIPT_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific script by its ID",
		"TOOL_GET_SCRIPT_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific script by its name",
		"TOOL_CREATE_SCRIPT_DESCRIPTION":      "Create a new script in Jamf Pro",
		"TOOL_UPDATE_SCRIPT_DESCRIPTION":      "Update an existing script in Jamf Pro",
		"TOOL_DELETE_SCRIPT_DESCRIPTION":      "Delete a script from Jamf Pro by its ID",

		// Buildings toolset
		"TOOL_GET_BUILDINGS_DESCRIPTION":      "Retrieve a list of all buildings from Jamf Pro",
		"TOOL_GET_BUILDING_BY_ID_DESCRIPTION": "Retrieve detailed information about a specific building by its ID",
		"TOOL_CREATE_BUILDING_DESCRIPTION":    "Create a new building in Jamf Pro",
		"TOOL_UPDATE_BUILDING_DESCRIPTION":    "Update an existing building in Jamf Pro",
		"TOOL_DELETE_BUILDING_DESCRIPTION":    "Delete a building from Jamf Pro by its ID",

		// Departments toolset
		"TOOL_GET_DEPARTMENTS_DESCRIPTION":        "Retrieve a list of all departments from Jamf Pro",
		"TOOL_GET_DEPARTMENT_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific department by its ID",
		"TOOL_GET_DEPARTMENT_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific department by its name",
		"TOOL_CREATE_DEPARTMENT_DESCRIPTION":      "Create a new department in Jamf Pro",
		"TOOL_UPDATE_DEPARTMENT_DESCRIPTION":      "Update an existing department in Jamf Pro",
		"TOOL_DELETE_DEPARTMENT_DESCRIPTION":      "Delete a department from Jamf Pro by its ID",

		// Categories toolset
		"TOOL_GET_CATEGORIES_DESCRIPTION":       "Retrieve a list of all categories from Jamf Pro",
		"TOOL_GET_CATEGORY_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific category by its ID",
		"TOOL_GET_CATEGORY_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific category by its name",
		"TOOL_CREATE_CATEGORY_DESCRIPTION":      "Create a new category in Jamf Pro",
		"TOOL_UPDATE_CATEGORY_DESCRIPTION":      "Update an existing category in Jamf Pro",
		"TOOL_DELETE_CATEGORY_DESCRIPTION":      "Delete a category from Jamf Pro by its ID",

		// Sites toolset
		"TOOL_GET_SITES_DESCRIPTION":        "Retrieve a list of all sites from Jamf Pro",
		"TOOL_GET_SITE_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific site by its ID",
		"TOOL_GET_SITE_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific site by its name",
		"TOOL_CREATE_SITE_DESCRIPTION":      "Create a new site in Jamf Pro",
		"TOOL_UPDATE_SITE_DESCRIPTION":      "Update an existing site in Jamf Pro",
		"TOOL_DELETE_SITE_DESCRIPTION":      "Delete a site from Jamf Pro by its ID",

		// API Roles toolset
		"TOOL_GET_API_ROLES_DESCRIPTION":      "Retrieve a list of all API roles from Jamf Pro",
		"TOOL_GET_API_ROLE_BY_ID_DESCRIPTION": "Retrieve detailed information about a specific API role by its ID",
		"TOOL_CREATE_API_ROLE_DESCRIPTION":    "Create a new API role in Jamf Pro",
		"TOOL_UPDATE_API_ROLE_DESCRIPTION":    "Update an existing API role in Jamf Pro",
		"TOOL_DELETE_API_ROLE_DESCRIPTION":    "Delete an API role from Jamf Pro by its ID",

		// API Integrations toolset
		"TOOL_GET_API_INTEGRATIONS_DESCRIPTION":      "Retrieve a list of all API integrations from Jamf Pro",
		"TOOL_GET_API_INTEGRATION_BY_ID_DESCRIPTION": "Retrieve detailed information about a specific API integration by its ID",
		"TOOL_CREATE_API_INTEGRATION_DESCRIPTION":    "Create a new API integration in Jamf Pro",
		"TOOL_UPDATE_API_INTEGRATION_DESCRIPTION":    "Update an existing API integration in Jamf Pro",
		"TOOL_DELETE_API_INTEGRATION_DESCRIPTION":    "Delete an API integration from Jamf Pro by its ID",

		// Jamf Pro Information toolset
		"TOOL_GET_JAMF_PRO_INFORMATION_DESCRIPTION": "Retrieve information about the Jamf Pro server and its capabilities",

		// Webhooks toolset
		"TOOL_GET_WEBHOOKS_DESCRIPTION":        "Retrieve a list of all webhooks from Jamf Pro",
		"TOOL_GET_WEBHOOK_BY_ID_DESCRIPTION":   "Retrieve detailed information about a specific webhook by its ID",
		"TOOL_GET_WEBHOOK_BY_NAME_DESCRIPTION": "Retrieve detailed information about a specific webhook by its name",
		"TOOL_CREATE_WEBHOOK_DESCRIPTION":      "Create a new webhook in Jamf Pro",
		"TOOL_UPDATE_WEBHOOK_DESCRIPTION":      "Update an existing webhook in Jamf Pro",
		"TOOL_DELETE_WEBHOOK_DESCRIPTION":      "Delete a webhook from Jamf Pro by its ID",

		// Advanced Searches toolset
		"TOOL_GET_ADVANCED_COMPUTER_SEARCHES_DESCRIPTION":      "Retrieve a list of all advanced computer searches from Jamf Pro",
		"TOOL_GET_ADVANCED_MOBILE_DEVICE_SEARCHES_DESCRIPTION": "Retrieve a list of all advanced mobile device searches from Jamf Pro",
		"TOOL_GET_ADVANCED_USER_SEARCHES_DESCRIPTION":          "Retrieve a list of all advanced user searches from Jamf Pro",

		// Extension Attributes toolset
		"TOOL_GET_COMPUTER_EXTENSION_ATTRIBUTES_DESCRIPTION":      "Retrieve a list of all computer extension attributes from Jamf Pro",
		"TOOL_GET_MOBILE_DEVICE_EXTENSION_ATTRIBUTES_DESCRIPTION": "Retrieve a list of all mobile device extension attributes from Jamf Pro",
		"TOOL_GET_USER_EXTENSION_ATTRIBUTES_DESCRIPTION":          "Retrieve a list of all user extension attributes from Jamf Pro",

		// Inventory toolset
		"TOOL_GET_COMPUTER_INVENTORY_COLLECTION_SETTINGS_DESCRIPTION":    "Retrieve computer inventory collection settings",
		"TOOL_UPDATE_COMPUTER_INVENTORY_COLLECTION_SETTINGS_DESCRIPTION": "Update computer inventory collection settings",

		// Network Segments toolset
		"TOOL_GET_NETWORK_SEGMENTS_DESCRIPTION":      "Retrieve a list of all network segments from Jamf Pro",
		"TOOL_GET_NETWORK_SEGMENT_BY_ID_DESCRIPTION": "Retrieve detailed information about a specific network segment by its ID",
		"TOOL_CREATE_NETWORK_SEGMENT_DESCRIPTION":    "Create a new network segment in Jamf Pro",
		"TOOL_UPDATE_NETWORK_SEGMENT_DESCRIPTION":    "Update an existing network segment in Jamf Pro",
		"TOOL_DELETE_NETWORK_SEGMENT_DESCRIPTION":    "Delete a network segment from Jamf Pro by its ID",

		// Printers toolset
		"TOOL_GET_PRINTERS_DESCRIPTION":      "Retrieve a list of all printers from Jamf Pro",
		"TOOL_GET_PRINTER_BY_ID_DESCRIPTION": "Retrieve detailed information about a specific printer by its ID",
		"TOOL_CREATE_PRINTER_DESCRIPTION":    "Create a new printer in Jamf Pro",
		"TOOL_UPDATE_PRINTER_DESCRIPTION":    "Update an existing printer in Jamf Pro",
		"TOOL_DELETE_PRINTER_DESCRIPTION":    "Delete a printer from Jamf Pro by its ID",
	}
}

// LoadTranslations loads tool description overrides from file or environment variables
func LoadTranslations() map[string]string {
	translations := GetDefaultTranslations()

	// Try to load from config file
	if data, err := os.ReadFile("jamfpro-mcp-server-config.json"); err == nil {
		var configTranslations map[string]string
		if err := json.Unmarshal(data, &configTranslations); err == nil {
			// Override defaults with config file values
			for key, value := range configTranslations {
				translations[key] = value
			}
		}
	}

	// Override with environment variables
	for key := range translations {
		if value := os.Getenv("JAMF_MCP_" + key); value != "" {
			translations[key] = value
		}
	}

	return translations
}

// GetToolDescription gets the description for a tool, with override support
func GetToolDescription(toolName string) string {
	translations := LoadTranslations()

	// Convert tool name to translation key
	key := "TOOL_" + toolName + "_DESCRIPTION"

	if description, exists := translations[key]; exists {
		return description
	}

	// Return a default description if no translation found
	return "Tool: " + toolName
}
