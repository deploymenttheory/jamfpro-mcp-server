package toolsets

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"go.uber.org/zap"
)

// PoliciesToolset handles policy-related operations using Jamf Pro Classic API
type PoliciesToolset struct {
	*BaseToolset
}

// NewPoliciesToolset creates a new policies toolset
func NewPoliciesToolset(client JamfProClient, logger *zap.Logger) *PoliciesToolset {
	base := NewBaseToolset(
		"policies",
		"Tools for managing policies in Jamf Pro using the Classic API, including CRUD operations for deployment configurations",
		client,
		logger,
	)

	toolset := &PoliciesToolset{
		BaseToolset: base,
	}

	// Add tools based on actual SDK capabilities
	toolset.addTools()

	return toolset
}

// addTools adds all policy-related tools
func (p *PoliciesToolset) addTools() {
	// Get Policies List
	p.AddTool(mcp.Tool{
		Name:        "get_policies",
		Description: "Retrieve a list of all policies from Jamf Pro",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
	})

	// Get Policy by ID
	p.AddTool(mcp.Tool{
		Name:        "get_policy_by_id",
		Description: "Retrieve a policy by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the policy to retrieve",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Policy by Name
	p.AddTool(mcp.Tool{
		Name:        "get_policy_by_name",
		Description: "Retrieve a policy by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the policy to retrieve",
				},
			},
			Required: []string{"name"},
		},
	})

	// Get Policies by Category
	p.AddTool(mcp.Tool{
		Name:        "get_policies_by_category",
		Description: "Retrieve policies by their category",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"category": map[string]interface{}{
					"type":        "string",
					"description": "The category name to filter policies by",
				},
			},
			Required: []string{"category"},
		},
	})

	// Get Policies by Type
	p.AddTool(mcp.Tool{
		Name:        "get_policies_by_type",
		Description: "Retrieve policies by the type of entity that created them (either 'casper' for Casper Remote or 'jss' for GUI/API created policies)",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"created_by": map[string]interface{}{
					"type":        "string",
					"description": "The entity type that created the policies (either 'casper' or 'jss')",
					"enum":        []string{"casper", "jss"},
				},
			},
			Required: []string{"created_by"},
		},
	})

	// Create Policy
	p.AddTool(mcp.Tool{
		Name:        "create_policy",
		Description: "Create a new policy in Jamf Pro with basic configuration",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the policy (required)",
				},
				"enabled": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is enabled",
					"default":     false,
				},
				"trigger_checkin": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is triggered on check-in",
					"default":     false,
				},
				"trigger_enrollment_complete": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is triggered on enrollment completion",
					"default":     false,
				},
				"trigger_login": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is triggered on login",
					"default":     false,
				},
				"trigger_logout": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is triggered on logout",
					"default":     false,
				},
				"trigger_network_state_changed": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is triggered when network state changes",
					"default":     false,
				},
				"trigger_startup": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is triggered on startup",
					"default":     false,
				},
				"trigger_other": map[string]interface{}{
					"type":        "string",
					"description": "Custom trigger for the policy",
					"default":     "",
				},
				"frequency": map[string]interface{}{
					"type":        "string",
					"description": "Frequency of the policy execution",
					"default":     "Once per computer",
					"enum": []string{
						"Once per computer",
						"Once per user per computer",
						"Once per user",
						"Once every day",
						"Once every week",
						"Once every month",
						"Ongoing",
					},
				},
				"category_id": map[string]interface{}{
					"type":        "integer",
					"description": "ID of the category for the policy",
					"default":     -1,
				},
				"category_name": map[string]interface{}{
					"type":        "string",
					"description": "Name of the category for the policy",
					"default":     "No category assigned",
				},
				"site_id": map[string]interface{}{
					"type":        "integer",
					"description": "ID of the site for the policy",
					"default":     -1,
				},
				"site_name": map[string]interface{}{
					"type":        "string",
					"description": "Name of the site for the policy",
					"default":     "None",
				},
				"all_computers": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy applies to all computers",
					"default":     false,
				},
				"self_service": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the policy is available in Self Service",
					"default":     false,
				},
				"run_maintenance": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether to run maintenance tasks",
					"default":     false,
				},
			},
			Required: []string{"name"},
		},
	})

	// Delete Policy by ID
	p.AddTool(mcp.Tool{
		Name:        "delete_policy_by_id",
		Description: "Delete a policy from Jamf Pro by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the policy to delete",
				},
			},
			Required: []string{"id"},
		},
	})

	// Delete Policy by Name
	p.AddTool(mcp.Tool{
		Name:        "delete_policy_by_name",
		Description: "Delete a policy from Jamf Pro by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the policy to delete",
				},
			},
			Required: []string{"name"},
		},
	})
}

// ExecuteTool executes a policy-related tool
func (p *PoliciesToolset) ExecuteTool(ctx context.Context, toolName string, arguments map[string]interface{}) (string, error) {
	p.GetLogger().Debug("Executing policies tool", zap.String("tool", toolName))

	switch toolName {
	case "get_policies":
		return p.getPolicies(ctx)
	case "get_policy_by_id":
		return p.getPolicyByID(ctx, arguments)
	case "get_policy_by_name":
		return p.getPolicyByName(ctx, arguments)
	case "get_policies_by_category":
		return p.getPoliciesByCategory(ctx, arguments)
	case "get_policies_by_type":
		return p.getPoliciesByType(ctx, arguments)
	case "create_policy":
		return p.createPolicy(ctx, arguments)
	case "delete_policy_by_id":
		return p.deletePolicyByID(ctx, arguments)
	case "delete_policy_by_name":
		return p.deletePolicyByName(ctx, arguments)
	default:
		return "", fmt.Errorf("unknown tool: %s", toolName)
	}
}

// getPolicies retrieves all policies
func (p *PoliciesToolset) getPolicies(ctx context.Context) (string, error) {
	policies, err := p.GetClient().GetPolicies()
	if err != nil {
		return "", fmt.Errorf("failed to get policies: %w", err)
	}

	response, err := FormatJSONResponse(policies)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d policies:\n\n%s", policies.Size, response), nil
}

// getPolicyByID retrieves a policy by ID
func (p *PoliciesToolset) getPolicyByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	policy, err := p.GetClient().GetPolicyByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get policy with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(policy)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Policy details for ID %s:\n\n%s", id, response), nil
}

// getPolicyByName retrieves a policy by name
func (p *PoliciesToolset) getPolicyByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	policy, err := p.GetClient().GetPolicyByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get policy with name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(policy)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Policy details for name '%s':\n\n%s", name, response), nil
}

// getPoliciesByCategory retrieves policies by category
func (p *PoliciesToolset) getPoliciesByCategory(ctx context.Context, args map[string]interface{}) (string, error) {
	category, err := GetStringArgument(args, "category", true)
	if err != nil {
		return "", err
	}

	policies, err := p.GetClient().GetPolicyByCategory(category)
	if err != nil {
		return "", fmt.Errorf("failed to get policies with category %s: %w", category, err)
	}

	response, err := FormatJSONResponse(policies)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d policies in category '%s':\n\n%s", policies.Size, category, response), nil
}

// getPoliciesByType retrieves policies by type
func (p *PoliciesToolset) getPoliciesByType(ctx context.Context, args map[string]interface{}) (string, error) {
	createdBy, err := GetStringArgument(args, "created_by", true)
	if err != nil {
		return "", err
	}

	policies, err := p.GetClient().GetPoliciesByType(createdBy)
	if err != nil {
		return "", fmt.Errorf("failed to get policies created by %s: %w", createdBy, err)
	}

	response, err := FormatJSONResponse(policies)
	if err != nil {
		return "", err
	}

	typeDesc := "Jamf Pro GUI/API"
	if createdBy == "casper" {
		typeDesc = "Casper Remote"
	}

	return fmt.Sprintf("Found %d policies created by %s:\n\n%s", policies.Size, typeDesc, response), nil
}

// createPolicy creates a new policy
func (p *PoliciesToolset) createPolicy(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	// Build the policy with the minimum required configuration
	policy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name: name,
		},
	}

	// Set optional fields if provided
	if enabled, err := GetBoolArgument(args, "enabled", false); err == nil {
		policy.General.Enabled = enabled
	}

	// Trigger settings
	if triggerCheckin, err := GetBoolArgument(args, "trigger_checkin", false); err == nil {
		policy.General.TriggerCheckin = triggerCheckin
	}
	if triggerEnrollmentComplete, err := GetBoolArgument(args, "trigger_enrollment_complete", false); err == nil {
		policy.General.TriggerEnrollmentComplete = triggerEnrollmentComplete
	}
	if triggerLogin, err := GetBoolArgument(args, "trigger_login", false); err == nil {
		policy.General.TriggerLogin = triggerLogin
	}
	if triggerLogout, err := GetBoolArgument(args, "trigger_logout", false); err == nil {
		policy.General.TriggerLogout = triggerLogout
	}
	if triggerNetworkStateChanged, err := GetBoolArgument(args, "trigger_network_state_changed", false); err == nil {
		policy.General.TriggerNetworkStateChanged = triggerNetworkStateChanged
	}
	if triggerStartup, err := GetBoolArgument(args, "trigger_startup", false); err == nil {
		policy.General.TriggerStartup = triggerStartup
	}
	if triggerOther, err := GetStringArgument(args, "trigger_other", false); err == nil && triggerOther != "" {
		policy.General.TriggerOther = triggerOther
	} else {
		policy.General.TriggerOther = "EVENT" // Default trigger
	}

	// Frequency
	if frequency, err := GetStringArgument(args, "frequency", false); err == nil && frequency != "" {
		policy.General.Frequency = frequency
	} else {
		policy.General.Frequency = "Once per computer" // Default frequency
	}

	// Add default settings for retry and offline
	policy.General.RetryEvent = "none"
	policy.General.RetryAttempts = -1
	policy.General.NotifyOnEachFailedRetry = false
	policy.General.LocationUserOnly = false
	policy.General.TargetDrive = "/"
	policy.General.Offline = false

	// Category
	categoryID, err := GetIntArgument(args, "category_id", false)
	if err == nil {
		categoryName, _ := GetStringArgument(args, "category_name", false)
		if categoryName == "" {
			categoryName = "No category assigned"
		}
		policy.General.Category = &jamfpro.SharedResourceCategory{
			ID:   categoryID,
			Name: categoryName,
		}
	}

	// Site
	siteID, err := GetIntArgument(args, "site_id", false)
	if err == nil {
		siteName, _ := GetStringArgument(args, "site_name", false)
		if siteName == "" {
			siteName = "None"
		}
		policy.General.Site = &jamfpro.SharedResourceSite{
			ID:   siteID,
			Name: siteName,
		}
	}

	// Scope
	allComputers, _ := GetBoolArgument(args, "all_computers", false)
	policy.Scope = jamfpro.PolicySubsetScope{
		AllComputers: allComputers,
		AllJSSUsers:  false,
	}

	// Self Service
	selfService, _ := GetBoolArgument(args, "self_service", false)
	if selfService {
		policy.SelfService = jamfpro.PolicySubsetSelfService{
			UseForSelfService:           true,
			InstallButtonText:           "Install",
			ReinstallButtonText:         "Reinstall",
			ForceUsersToViewDescription: false,
			FeatureOnMainPage:           false,
			Notification:                false,
		}
	}

	// Maintenance
	runMaintenance, _ := GetBoolArgument(args, "run_maintenance", false)
	policy.Maintenance = jamfpro.PolicySubsetMaintenance{
		Recon:                    runMaintenance,
		ResetName:                false,
		InstallAllCachedPackages: false,
		Heal:                     false,
		Prebindings:              false,
		Permissions:              false,
		Byhost:                   false,
		SystemCache:              false,
		UserCache:                false,
		Verify:                   false,
	}

	// Create the policy
	createdPolicy, err := p.GetClient().CreatePolicy(policy)
	if err != nil {
		return "", fmt.Errorf("failed to create policy '%s': %w", name, err)
	}

	response, err := FormatJSONResponse(createdPolicy)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully created policy '%s' with ID %d:\n\n%s", name, createdPolicy.ID, response), nil
}

// deletePolicyByID deletes a policy by ID
func (p *PoliciesToolset) deletePolicyByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	err = p.GetClient().DeletePolicyByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to delete policy with ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully deleted policy with ID %s", id), nil
}

// deletePolicyByName deletes a policy by name
func (p *PoliciesToolset) deletePolicyByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	err = p.GetClient().DeletePolicyByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to delete policy with name %s: %w", name, err)
	}

	return fmt.Sprintf("Successfully deleted policy with name '%s'", name), nil
}
