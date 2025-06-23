package toolsets

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"go.uber.org/zap"
)

// JamfProClient defines the interface for Jamf Pro API client
// This allows for dependency injection and easier testing
type JamfProClient interface {
	// Common methods
	GetJamfProInformation() (*jamfpro.ResponseJamfProInformation, error)

	// Computer methods (Classic API)
	GetComputers() (*jamfpro.ResponseComputersList, error)
	GetComputerByID(id string) (*jamfpro.ResponseComputer, error)
	GetComputerByName(name string) (*jamfpro.ResponseComputer, error)
	GetComputerGroups() (*jamfpro.ResponseComputerGroupsList, error)
	GetComputerGroupByID(id string) (*jamfpro.ResourceComputerGroup, error)
	CreateComputer(computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error)
	UpdateComputerByID(id string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error)
	UpdateComputerByName(name string, computer jamfpro.ResponseComputer) (*jamfpro.ResponseComputer, error)
	DeleteComputerByID(id string) error
	DeleteComputerByName(name string) error

	// Computer Inventory methods (Pro API)
	GetComputersInventory(params url.Values) (*jamfpro.ResponseComputerInventoryList, error)
	GetComputerInventoryByID(id string) (*jamfpro.ResourceComputerInventory, error)
	GetComputerInventoryByName(name string) (*jamfpro.ResourceComputerInventory, error)
	UpdateComputerInventoryByID(id string, inventory *jamfpro.ResourceComputerInventory) (*jamfpro.ResourceComputerInventory, error)
	DeleteComputerInventoryByID(id string) error

	// FileVault methods
	GetComputersFileVaultInventory(params url.Values) (*jamfpro.FileVaultInventoryList, error)
	GetComputerFileVaultInventoryByID(id string) (*jamfpro.FileVaultInventory, error)
	GetComputerRecoveryLockPasswordByID(id string) (*jamfpro.ResponseRecoveryLockPassword, error)

	// Device management methods
	RemoveComputerMDMProfile(id string) (*jamfpro.ResponseRemoveMDMProfile, error)
	EraseComputerByID(id string, request jamfpro.RequestEraseDeviceComputer) error

	// Attachment methods
	UploadAttachmentAndAssignToComputerByID(computerID string, filePaths []string) (*jamfpro.ResponseUploadAttachment, error)
	DeleteAttachmentByIDAndComputerID(computerID, attachmentID string) error

	// Mobile Device methods (Classic API)
	GetMobileDevices() (*jamfpro.ResponseMobileDeviceList, error)
	GetMobileDeviceByID(id string) (*jamfpro.ResourceMobileDevice, error)
	GetMobileDeviceByName(name string) (*jamfpro.ResourceMobileDevice, error)
	GetMobileDeviceGroups() (*jamfpro.ResponseMobileDeviceGroupsList, error)
	GetMobileDeviceGroupByID(id string) (*jamfpro.ResourceMobileDeviceGroup, error)
	GetMobileDeviceApplications() (*jamfpro.ResponseMobileDeviceApplicationsList, error)
	GetMobileDeviceConfigurationProfiles() (*jamfpro.ResponseMobileDeviceConfigurationProfilesList, error)
	CreateMobileDevice(device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error)
	UpdateMobileDeviceByID(id string, device *jamfpro.ResourceMobileDevice) (*jamfpro.ResourceMobileDevice, error)
	DeleteMobileDeviceByID(id string) error

	// Policies methods (Classic API) - FIXED: Corrected return types
	GetPolicies() (*jamfpro.ResponsePoliciesList, error)
	GetPolicyByID(id string) (*jamfpro.ResourcePolicy, error)
	GetPolicyByName(name string) (*jamfpro.ResourcePolicy, error)
	GetPolicyByCategory(category string) (*jamfpro.ResponsePoliciesList, error)
	GetPoliciesByType(createdBy string) (*jamfpro.ResponsePoliciesList, error)
	CreatePolicy(policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error)
	UpdatePolicyByID(id string, policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error)
	UpdatePolicyByName(name string, policy *jamfpro.ResourcePolicy) (*jamfpro.ResponsePolicyCreateAndUpdate, error)
	DeletePolicyByID(id string) error
	DeletePolicyByName(name string) error

	// Scripts methods (Pro API)
	GetScripts(params url.Values) (*jamfpro.ResponseScriptsList, error)
	GetScriptByID(id string) (*jamfpro.ResourceScript, error)
	GetScriptByName(name string) (*jamfpro.ResourceScript, error)
	CreateScript(script *jamfpro.ResourceScript) (*jamfpro.ResponseScriptCreate, error)
	UpdateScriptByID(id string, script *jamfpro.ResourceScript) (*jamfpro.ResourceScript, error)
	UpdateScriptByName(name string, script *jamfpro.ResourceScript) (*jamfpro.ResourceScript, error)
	DeleteScriptByID(id string) error
	DeleteScriptByName(name string) error
}

// Toolset represents a collection of related tools
type Toolset interface {
	// GetName returns the name of the toolset
	GetName() string

	// GetDescription returns the description of the toolset
	GetDescription() string

	// GetTools returns the list of tools provided by this toolset
	GetTools() []mcp.Tool

	// ExecuteTool executes a specific tool with the given arguments
	ExecuteTool(ctx context.Context, toolName string, arguments map[string]interface{}) (string, error)
}

// BaseToolset provides common functionality for all toolsets
type BaseToolset struct {
	name        string
	description string
	client      JamfProClient
	logger      *zap.Logger
	tools       map[string]mcp.Tool
}

// NewBaseToolset creates a new base toolset
func NewBaseToolset(name, description string, client JamfProClient, logger *zap.Logger) *BaseToolset {
	return &BaseToolset{
		name:        name,
		description: description,
		client:      client,
		logger:      logger,
		tools:       make(map[string]mcp.Tool),
	}
}

// GetName returns the name of the toolset
func (b *BaseToolset) GetName() string {
	return b.name
}

// GetDescription returns the description of the toolset
func (b *BaseToolset) GetDescription() string {
	return b.description
}

// GetTools returns the list of tools
func (b *BaseToolset) GetTools() []mcp.Tool {
	tools := make([]mcp.Tool, 0, len(b.tools))
	for _, tool := range b.tools {
		tools = append(tools, tool)
	}
	return tools
}

// AddTool adds a tool to the toolset
func (b *BaseToolset) AddTool(tool mcp.Tool) {
	b.tools[tool.Name] = tool
}

// GetClient returns the Jamf Pro client
func (b *BaseToolset) GetClient() JamfProClient {
	return b.client
}

// GetLogger returns the logger
func (b *BaseToolset) GetLogger() *zap.Logger {
	return b.logger
}

// Factory creates toolsets
type Factory struct {
	client JamfProClient
	logger *zap.Logger
}

// NewFactory creates a new toolset factory
func NewFactory(client JamfProClient, logger *zap.Logger) *Factory {
	return &Factory{
		client: client,
		logger: logger,
	}
}

// CreateToolset creates a toolset by name
func (f *Factory) CreateToolset(name string) (Toolset, error) {
	switch name {

	// ========== DEVICE MANAGEMENT ==========
	// Based on jamfproapi_* and classicapi_* files

	case "computers":
		return NewComputersToolset(f.client, f.logger), nil
	case "computer-inventory":
		return NewComputerInventoryToolset(f.client, f.logger), nil
	case "mobile-devices":
		return NewMobileDevicesToolset(f.client, f.logger), nil
	case "mobile-device-inventory":
		return nil, fmt.Errorf("mobile-device-inventory toolset not yet implemented - based on jamfproapi_mobile_device_inventory.go")
	case "computer-groups":
		return nil, fmt.Errorf("computer-groups toolset not yet implemented - based on classicapi_computer_groups.go")
	case "mobile-device-groups":
		return nil, fmt.Errorf("mobile-device-groups toolset not yet implemented - based on classicapi_mobile_device_groups.go")
	case "smart-computer-groups":
		return nil, fmt.Errorf("smart-computer-groups toolset not yet implemented - based on classicapi_smart_computer_groups.go")
	case "smart-mobile-device-groups":
		return nil, fmt.Errorf("smart-mobile-device-groups toolset not yet implemented - based on classicapi_smart_mobile_device_groups.go")

	// ========== POLICIES & CONFIGURATION ==========

	case "policies":
		return NewPoliciesToolset(f.client, f.logger), nil
	case "configuration-profiles":
		return nil, fmt.Errorf("configuration-profiles toolset not yet implemented - based on classicapi_os_x_configuration_profiles.go")
	case "mobile-device-configuration-profiles":
		return nil, fmt.Errorf("mobile-device-configuration-profiles toolset not yet implemented - based on classicapi_mobile_device_configuration_profiles.go")
	case "computer-configuration-profiles":
		return nil, fmt.Errorf("computer-configuration-profiles toolset not yet implemented - based on jamfproapi_computer_configuration_profiles.go")
	case "restricted-software":
		return nil, fmt.Errorf("restricted-software toolset not yet implemented - based on classicapi_restricted_software.go")
	case "patch-policies":
		return nil, fmt.Errorf("patch-policies toolset not yet implemented - based on classicapi_patch_policies.go")

	// ========== USER & ACCESS MANAGEMENT ==========

	case "users":
		return nil, fmt.Errorf("users toolset not yet implemented - based on classicapi_users.go")
	case "groups":
		return nil, fmt.Errorf("groups toolset not yet implemented - based on classicapi_groups.go")
	case "user-groups":
		return nil, fmt.Errorf("user-groups toolset not yet implemented - based on jamfproapi_user_groups.go")
	case "accounts":
		return nil, fmt.Errorf("accounts toolset not yet implemented - based on classicapi_accounts.go")
	case "ldap-servers":
		return nil, fmt.Errorf("ldap-servers toolset not yet implemented - based on classicapi_ldap_servers.go")
	case "ldap":
		return nil, fmt.Errorf("ldap toolset not yet implemented - based on jamfproapi_ldap.go")
	case "api-roles":
		return nil, fmt.Errorf("api-roles toolset not yet implemented - based on jamfproapi_api_roles.go")
	case "api-integrations":
		return nil, fmt.Errorf("api-integrations toolset not yet implemented - based on jamfproapi_api_integrations.go")
	case "api-authentication":
		return nil, fmt.Errorf("api-authentication toolset not yet implemented - based on jamfproapi_api_authentication.go")

		// ========== APPLICATIONS & SOFTWARE ==========

	case "scripts":
		return NewScriptsToolset(f.client, f.logger), nil
	case "mobile-device-applications":
		return nil, fmt.Errorf("mobile-device-applications toolset not yet implemented - based on classicapi_mobile_device_applications.go")
	case "mac-applications":
		return nil, fmt.Errorf("mac-applications toolset not yet implemented - based on classicapi_mac_applications.go")
	case "mobile-applications":
		return nil, fmt.Errorf("mobile-applications toolset not yet implemented - based on jamfproapi_mobile_applications.go")
	case "packages":
		return nil, fmt.Errorf("packages toolset not yet implemented - based on classicapi_packages.go")
	case "patch-software-title-configurations":
		return nil, fmt.Errorf("patch-software-title-configurations toolset not yet implemented - based on classicapi_patch_software_title_configurations.go")
	case "patch-management":
		return nil, fmt.Errorf("patch-management toolset not yet implemented - based on jamfproapi_patch_management.go")
	case "licensed-software":
		return nil, fmt.Errorf("licensed-software toolset not yet implemented - based on classicapi_licensed_software.go")
	case "software-update-servers":
		return nil, fmt.Errorf("software-update-servers toolset not yet implemented - based on classicapi_software_update_servers.go")

	// ========== INFRASTRUCTURE & NETWORKING ==========

	case "buildings":
		return nil, fmt.Errorf("buildings toolset not yet implemented - based on classicapi_buildings.go")
	case "departments":
		return nil, fmt.Errorf("departments toolset not yet implemented - based on classicapi_departments.go")
	case "categories":
		return nil, fmt.Errorf("categories toolset not yet implemented - based on classicapi_categories.go")
	case "sites":
		return nil, fmt.Errorf("sites toolset not yet implemented - based on classicapi_sites.go")
	case "network-segments":
		return nil, fmt.Errorf("network-segments toolset not yet implemented - based on classicapi_network_segments.go")
	case "printers":
		return nil, fmt.Errorf("printers toolset not yet implemented - based on classicapi_printers.go")
	case "distribution-points":
		return nil, fmt.Errorf("distribution-points toolset not yet implemented - based on classicapi_distribution_points.go")
	case "directory-bindings":
		return nil, fmt.Errorf("directory-bindings toolset not yet implemented - based on classicapi_directory_bindings.go")
	case "dock-items":
		return nil, fmt.Errorf("dock-items toolset not yet implemented - based on classicapi_dock_items.go")
	case "removable-mac-addresses":
		return nil, fmt.Errorf("removable-mac-addresses toolset not yet implemented - based on classicapi_removable_mac_addresses.go")

	// ========== ENROLLMENT & PROVISIONING ==========

	case "enrollment":
		return nil, fmt.Errorf("enrollment toolset not yet implemented - based on jamfproapi_enrollment.go")
	case "device-enrollment-program":
		return nil, fmt.Errorf("device-enrollment-program toolset not yet implemented - based on classicapi_device_enrollment_program.go")
	case "computer-prestage-enrollments":
		return nil, fmt.Errorf("computer-prestage-enrollments toolset not yet implemented - based on jamfproapi_computer_prestage_enrollments.go")
	case "mobile-device-prestage-enrollments":
		return nil, fmt.Errorf("mobile-device-prestage-enrollments toolset not yet implemented - based on jamfproapi_mobile_device_prestage_enrollments.go")
	case "enrollment-customization":
		return nil, fmt.Errorf("enrollment-customization toolset not yet implemented - based on jamfproapi_enrollment_customization.go")
	case "byo-profiles":
		return nil, fmt.Errorf("byo-profiles toolset not yet implemented - based on classicapi_byo_profiles.go")

	// ========== SEARCHING & REPORTING ==========

	case "advanced-computer-searches":
		return nil, fmt.Errorf("advanced-computer-searches toolset not yet implemented - based on classicapi_advanced_computer_searches.go")
	case "advanced-mobile-device-searches":
		return nil, fmt.Errorf("advanced-mobile-device-searches toolset not yet implemented - based on classicapi_advanced_mobile_device_searches.go")
	case "advanced-user-searches":
		return nil, fmt.Errorf("advanced-user-searches toolset not yet implemented - based on classicapi_advanced_user_searches.go")
	case "advanced-searches":
		return nil, fmt.Errorf("advanced-searches toolset not yet implemented - combined search functionality")
	case "computer-reports":
		return nil, fmt.Errorf("computer-reports toolset not yet implemented - based on classicapi_computer_reports.go")

	// ========== EXTENSION ATTRIBUTES ==========

	case "computer-extension-attributes":
		return nil, fmt.Errorf("computer-extension-attributes toolset not yet implemented - based on classicapi_computer_extension_attributes.go")
	case "mobile-device-extension-attributes":
		return nil, fmt.Errorf("mobile-device-extension-attributes toolset not yet implemented - based on classicapi_mobile_device_extension_attributes.go")
	case "user-extension-attributes":
		return nil, fmt.Errorf("user-extension-attributes toolset not yet implemented - based on classicapi_user_extension_attributes.go")
	case "extension-attributes":
		return nil, fmt.Errorf("extension-attributes toolset not yet implemented - combined extension attributes functionality")

	// ========== SECURITY & ENCRYPTION ==========

	case "disk-encryption":
		return nil, fmt.Errorf("disk-encryption toolset not yet implemented - based on classicapi_disk_encryption.go")
	case "filevault":
		return nil, fmt.Errorf("filevault toolset not yet implemented - part of computer inventory")
	case "activation-code":
		return nil, fmt.Errorf("activation-code toolset not yet implemented - based on classicapi_activation_code.go")
	case "computer-checkin":
		return nil, fmt.Errorf("computer-checkin toolset not yet implemented - based on classicapi_computer_checkin.go")
	case "gsx-connection":
		return nil, fmt.Errorf("gsx-connection toolset not yet implemented - based on classicapi_gsx_connection.go")

	// ========== VOLUME PURCHASING & LICENSING ==========

	case "vpp-accounts":
		return nil, fmt.Errorf("vpp-accounts toolset not yet implemented - based on classicapi_vpp_accounts.go")
	case "vpp-assignments":
		return nil, fmt.Errorf("vpp-assignments toolset not yet implemented - based on classicapi_vpp_assignments.go")
	case "vpp-invitations":
		return nil, fmt.Errorf("vpp-invitations toolset not yet implemented - based on classicapi_vpp_invitations.go")
	case "vpp":
		return nil, fmt.Errorf("vpp toolset not yet implemented - combined VPP functionality")
	case "volume-purchasing":
		return nil, fmt.Errorf("volume-purchasing toolset not yet implemented - based on classicapi_volume_purchasing.go")

	// ========== SELF SERVICE & AUTOMATION ==========

	case "self-service":
		return nil, fmt.Errorf("self-service toolset not yet implemented - based on jamfproapi_self_service.go")
	case "self-service-branding":
		return nil, fmt.Errorf("self-service-branding toolset not yet implemented - based on jamfproapi_self_service_branding.go")
	case "webhooks":
		return nil, fmt.Errorf("webhooks toolset not yet implemented - based on classicapi_webhooks.go")
	case "computer-commands":
		return nil, fmt.Errorf("computer-commands toolset not yet implemented - based on jamfproapi_computer_commands.go")
	case "mobile-device-commands":
		return nil, fmt.Errorf("mobile-device-commands toolset not yet implemented - based on jamfproapi_mobile_device_commands.go")

	// ========== SYSTEM ADMINISTRATION ==========

	case "jamf-pro-information":
		return nil, fmt.Errorf("jamf-pro-information toolset not yet implemented - based on classicapi_jamf_pro_information.go")
	case "jamf-pro-server-information":
		return nil, fmt.Errorf("jamf-pro-server-information toolset not yet implemented - based on jamfproapi_jamf_pro_server_information.go")
	case "jamf-pro-version":
		return nil, fmt.Errorf("jamf-pro-version toolset not yet implemented - based on jamfproapi_jamf_pro_version.go")
	case "smtp-server":
		return nil, fmt.Errorf("smtp-server toolset not yet implemented - based on classicapi_smtp_server.go")
	case "smtp":
		return nil, fmt.Errorf("smtp toolset not yet implemented - based on smtp server functionality")
	case "sso":
		return nil, fmt.Errorf("sso toolset not yet implemented - based on sso functionality")
	case "sso-certificate":
		return nil, fmt.Errorf("sso-certificate toolset not yet implemented - based on jamfproapi_sso_certificate.go")
	case "sso-failover":
		return nil, fmt.Errorf("sso-failover toolset not yet implemented - based on jamfproapi_sso_failover.go")

	// ========== MOBILE DEVICE SPECIFIC ==========

	case "mobile-device-enrollment-profiles":
		return nil, fmt.Errorf("mobile-device-enrollment-profiles toolset not yet implemented - based on classicapi_mobile_device_enrollment_profiles.go")
	case "mobile-device-provisioning-profiles":
		return nil, fmt.Errorf("mobile-device-provisioning-profiles toolset not yet implemented - based on classicapi_mobile_device_provisioning_profiles.go")
	case "ibeacons":
		return nil, fmt.Errorf("ibeacons toolset not yet implemented - based on classicapi_ibeacons.go")
	case "ebooks":
		return nil, fmt.Errorf("ebooks toolset not yet implemented - based on classicapi_ebooks.go")

	// ========== PERIPHERAL & HARDWARE ==========

	case "peripherals":
		return nil, fmt.Errorf("peripherals toolset not yet implemented - based on classicapi_peripherals.go")
	case "peripheral-types":
		return nil, fmt.Errorf("peripheral-types toolset not yet implemented - based on classicapi_peripheral_types.go")
	case "computer-hardware":
		return nil, fmt.Errorf("computer-hardware toolset not yet implemented - part of computer inventory")

	// ========== FILE & CONTENT MANAGEMENT ==========

	case "allowed-file-extensions":
		return nil, fmt.Errorf("allowed-file-extensions toolset not yet implemented - based on classicapi_allowed_file_extensions.go")
	case "file-uploads":
		return nil, fmt.Errorf("file-uploads toolset not yet implemented - based on jamfproapi_file_uploads.go")
	case "cloud-azure":
		return nil, fmt.Errorf("cloud-azure toolset not yet implemented - based on jamfproapi_cloud_azure.go")
	case "cloud-identity-providers":
		return nil, fmt.Errorf("cloud-identity-providers toolset not yet implemented - based on jamfproapi_cloud_identity_providers.go")
	case "cloud-information":
		return nil, fmt.Errorf("cloud-information toolset not yet implemented - based on jamfproapi_cloud_information.go")

	// ========== LEGACY & DEPRECATED ==========

	case "save-computer-reports":
		return nil, fmt.Errorf("save-computer-reports toolset not yet implemented - based on classicapi_save_computer_reports.go")
	case "computer-inventory-collection-settings":
		return nil, fmt.Errorf("computer-inventory-collection-settings toolset not yet implemented - based on jamfproapi_computer_inventory_collection_settings.go")
	case "inventory":
		return nil, fmt.Errorf("inventory toolset not yet implemented - generic inventory functionality")

	// ========== CLUSTERING & INFRASTRUCTURE ==========

	case "jcds":
		return nil, fmt.Errorf("jcds toolset not yet implemented - based on classicapi_jcds.go")
	case "jcds2":
		return nil, fmt.Errorf("jcds2 toolset not yet implemented - based on classicapi_jcds2.go")
	case "cache-settings":
		return nil, fmt.Errorf("cache-settings toolset not yet implemented - based on jamfproapi_cache_settings.go")
	case "client-check-in":
		return nil, fmt.Errorf("client-check-in toolset not yet implemented - based on jamfproapi_client_check_in.go")

	// ========== ALERTS & NOTIFICATIONS ==========

	case "alerts":
		return nil, fmt.Errorf("alerts toolset not yet implemented - based on jamfproapi_alerts.go")
	case "app-store-country-codes":
		return nil, fmt.Errorf("app-store-country-codes toolset not yet implemented - based on jamfproapi_app_store_country_codes.go")

	// ========== SECURITY COMPLIANCE ==========

	case "computer-security":
		return nil, fmt.Errorf("computer-security toolset not yet implemented - part of computer inventory")
	case "compliance-vendor-device-information":
		return nil, fmt.Errorf("compliance-vendor-device-information toolset not yet implemented - based on jamfproapi_compliance_vendor_device_information.go")
	case "conditional-access":
		return nil, fmt.Errorf("conditional-access toolset not yet implemented - based on jamfproapi_conditional_access.go")

	// ========== MAINTENANCE & UTILITIES ==========

	case "re-enrollment":
		return nil, fmt.Errorf("re-enrollment toolset not yet implemented - based on jamfproapi_re_enrollment.go")
	case "returntoservice":
		return nil, fmt.Errorf("returntoservice toolset not yet implemented - based on jamfproapi_returntoservice.go")
	case "team-viewer":
		return nil, fmt.Errorf("team-viewer toolset not yet implemented - based on jamfproapi_team_viewer.go")
	case "engage":
		return nil, fmt.Errorf("engage toolset not yet implemented - based on jamfproapi_engage.go")
	case "supervision-identities":
		return nil, fmt.Errorf("supervision-identities toolset not yet implemented - based on jamfproapi_supervision_identities.go")
	case "supervision-identity-certificate":
		return nil, fmt.Errorf("supervision-identity-certificate toolset not yet implemented - based on jamfproapi_supervision_identity_certificate.go")

	default:
		return nil, fmt.Errorf("unknown toolset: %s", name)
	}
}

// Helper functions for common argument validation and conversion

// GetStringArgument safely gets a string argument
func GetStringArgument(args map[string]interface{}, key string, required bool) (string, error) {
	value, exists := args[key]
	if !exists {
		if required {
			return "", fmt.Errorf("required argument %s is missing", key)
		}
		return "", nil
	}

	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("argument %s must be a string", key)
	}

	if required && str == "" {
		return "", fmt.Errorf("required argument %s cannot be empty", key)
	}

	return str, nil
}

// GetIntArgument safely gets an integer argument
func GetIntArgument(args map[string]interface{}, key string, required bool) (int, error) {
	value, exists := args[key]
	if !exists {
		if required {
			return 0, fmt.Errorf("required argument %s is missing", key)
		}
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		// Try to parse string as int - FIXED: Don't allow string conversion for int args
		if v == "" && !required {
			return 0, nil
		}
		return 0, fmt.Errorf("argument %s must be a number, got string: %s", key, v)
	default:
		return 0, fmt.Errorf("argument %s must be a number", key)
	}
}

// GetBoolArgument safely gets a boolean argument
func GetBoolArgument(args map[string]interface{}, key string, required bool) (bool, error) {
	value, exists := args[key]
	if !exists {
		if required {
			return false, fmt.Errorf("required argument %s is missing", key)
		}
		return false, nil
	}

	boolean, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("argument %s must be a boolean", key)
	}

	return boolean, nil
}

// GetStringSliceArgument safely gets a string slice argument
func GetStringSliceArgument(args map[string]interface{}, key string, required bool) ([]string, error) {
	value, exists := args[key]
	if !exists {
		if required {
			return nil, fmt.Errorf("required argument %s is missing", key)
		}
		return []string{}, nil
	}

	slice, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("argument %s must be an array", key)
	}

	result := make([]string, len(slice))
	for i, item := range slice {
		str, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("argument %s must be an array of strings", key)
		}
		result[i] = str
	}

	return result, nil
}

// FormatJSONResponse formats a response as pretty-printed JSON
func FormatJSONResponse(data interface{}) (string, error) {
	if data == nil {
		return "No data returned", nil
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format response as JSON: %w", err)
	}

	return string(jsonBytes), nil
}

// FormatListResponse formats a list response with count information - FIXED: Better type handling
func FormatListResponse(items interface{}, itemType string) (string, error) {
	jsonBytes, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format response as JSON: %w", err)
	}

	// Try to get count from different response types
	count := "unknown"

	// Handle different response types that might have count fields
	if respWithCount, ok := items.(interface{ GetSize() int }); ok {
		count = fmt.Sprintf("%d", respWithCount.GetSize())
	} else if respWithTotal, ok := items.(interface{ GetTotalCount() int }); ok {
		count = fmt.Sprintf("%d", respWithTotal.GetTotalCount())
	} else if slice, ok := items.([]interface{}); ok {
		count = fmt.Sprintf("%d", len(slice))
	}

	return fmt.Sprintf("Found %s %s:\n\n%s", count, itemType, string(jsonBytes)), nil
}
