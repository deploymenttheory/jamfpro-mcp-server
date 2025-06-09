package toolsets

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"go.uber.org/zap"
)

// ScriptsToolset handles script-related operations using Jamf Pro API
type ScriptsToolset struct {
	*BaseToolset
}

// NewScriptsToolset creates a new scripts toolset
func NewScriptsToolset(client *jamfpro.Client, logger *zap.Logger) *ScriptsToolset {
	base := NewBaseToolset(
		"scripts",
		"Tools for managing scripts in Jamf Pro using the Pro API, including full CRUD operations and script parameter management",
		client,
		logger,
	)

	toolset := &ScriptsToolset{
		BaseToolset: base,
	}

	// Add tools based on actual Pro API capabilities
	toolset.addTools()

	return toolset
}

// addTools adds all script-related tools
func (s *ScriptsToolset) addTools() {
	// Get Scripts List
	s.AddTool(mcp.Tool{
		Name:        "get_scripts",
		Description: "Retrieve a list of all scripts from Jamf Pro with optional pagination, sorting, and filtering",
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
					"description": "Sort field and direction (e.g., 'name:asc', 'categoryName:desc')",
				},
				"filter": map[string]interface{}{
					"type":        "string",
					"description": "Filter criteria for the search (e.g., 'name==\"My Script\"')",
				},
			},
			Required: []string{},
		},
	})

	// Get Script by ID
	s.AddTool(mcp.Tool{
		Name:        "get_script_by_id",
		Description: "Retrieve detailed information about a specific script by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the script to retrieve",
				},
			},
			Required: []string{"id"},
		},
	})

	// Get Script by Name
	s.AddTool(mcp.Tool{
		Name:        "get_script_by_name",
		Description: "Retrieve detailed information about a specific script by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the script to retrieve",
				},
			},
			Required: []string{"name"},
		},
	})

	// Create Script
	s.AddTool(mcp.Tool{
		Name:        "create_script",
		Description: "Create a new script in Jamf Pro with script contents and configuration",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Script name (required)",
				},
				"script_contents": map[string]interface{}{
					"type":        "string",
					"description": "The actual script code/contents (required)",
				},
				"category_name": map[string]interface{}{
					"type":        "string",
					"description": "Category name for the script",
				},
				"category_id": map[string]interface{}{
					"type":        "string",
					"description": "Category ID for the script",
				},
				"info": map[string]interface{}{
					"type":        "string",
					"description": "Script description/information",
				},
				"notes": map[string]interface{}{
					"type":        "string",
					"description": "Script notes",
				},
				"os_requirements": map[string]interface{}{
					"type":        "string",
					"description": "OS requirements for the script",
				},
				"priority": map[string]interface{}{
					"type":        "string",
					"description": "Script execution priority (Before, After, At Reboot)",
					"enum":        []string{"Before", "After", "At Reboot"},
				},
				"parameter_4": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 4 label",
				},
				"parameter_5": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 5 label",
				},
				"parameter_6": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 6 label",
				},
				"parameter_7": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 7 label",
				},
				"parameter_8": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 8 label",
				},
				"parameter_9": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 9 label",
				},
				"parameter_10": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 10 label",
				},
				"parameter_11": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 11 label",
				},
			},
			Required: []string{"name", "script_contents"},
		},
	})

	// Update Script by ID
	s.AddTool(mcp.Tool{
		Name:        "update_script_by_id",
		Description: "Update an existing script by its ID. Only specified fields will be updated.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the script to update (required)",
				},
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Script name",
				},
				"script_contents": map[string]interface{}{
					"type":        "string",
					"description": "The actual script code/contents",
				},
				"category_name": map[string]interface{}{
					"type":        "string",
					"description": "Category name for the script",
				},
				"category_id": map[string]interface{}{
					"type":        "string",
					"description": "Category ID for the script",
				},
				"info": map[string]interface{}{
					"type":        "string",
					"description": "Script description/information",
				},
				"notes": map[string]interface{}{
					"type":        "string",
					"description": "Script notes",
				},
				"os_requirements": map[string]interface{}{
					"type":        "string",
					"description": "OS requirements for the script",
				},
				"priority": map[string]interface{}{
					"type":        "string",
					"description": "Script execution priority",
					"enum":        []string{"Before", "After", "At Reboot"},
				},
				"parameter_4": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 4 label",
				},
				"parameter_5": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 5 label",
				},
				"parameter_6": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 6 label",
				},
				"parameter_7": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 7 label",
				},
				"parameter_8": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 8 label",
				},
				"parameter_9": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 9 label",
				},
				"parameter_10": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 10 label",
				},
				"parameter_11": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 11 label",
				},
			},
			Required: []string{"id"},
		},
	})

	// Update Script by Name
	s.AddTool(mcp.Tool{
		Name:        "update_script_by_name",
		Description: "Update an existing script by its name. Only specified fields will be updated.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the script to update (required)",
				},
				"new_name": map[string]interface{}{
					"type":        "string",
					"description": "New script name (if changing the name)",
				},
				"script_contents": map[string]interface{}{
					"type":        "string",
					"description": "The actual script code/contents",
				},
				"category_name": map[string]interface{}{
					"type":        "string",
					"description": "Category name for the script",
				},
				"category_id": map[string]interface{}{
					"type":        "string",
					"description": "Category ID for the script",
				},
				"info": map[string]interface{}{
					"type":        "string",
					"description": "Script description/information",
				},
				"notes": map[string]interface{}{
					"type":        "string",
					"description": "Script notes",
				},
				"os_requirements": map[string]interface{}{
					"type":        "string",
					"description": "OS requirements for the script",
				},
				"priority": map[string]interface{}{
					"type":        "string",
					"description": "Script execution priority",
					"enum":        []string{"Before", "After", "At Reboot"},
				},
				"parameter_4": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 4 label",
				},
				"parameter_5": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 5 label",
				},
				"parameter_6": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 6 label",
				},
				"parameter_7": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 7 label",
				},
				"parameter_8": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 8 label",
				},
				"parameter_9": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 9 label",
				},
				"parameter_10": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 10 label",
				},
				"parameter_11": map[string]interface{}{
					"type":        "string",
					"description": "Script parameter 11 label",
				},
			},
			Required: []string{"name"},
		},
	})

	// Delete Script by ID
	s.AddTool(mcp.Tool{
		Name:        "delete_script_by_id",
		Description: "Delete a script from Jamf Pro by its ID",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{
					"type":        "string",
					"description": "The ID of the script to delete",
				},
			},
			Required: []string{"id"},
		},
	})

	// Delete Script by Name
	s.AddTool(mcp.Tool{
		Name:        "delete_script_by_name",
		Description: "Delete a script from Jamf Pro by its name",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "The name of the script to delete",
				},
			},
			Required: []string{"name"},
		},
	})
}

// ExecuteTool executes a script-related tool
func (s *ScriptsToolset) ExecuteTool(ctx context.Context, toolName string, arguments map[string]interface{}) (string, error) {
	s.GetLogger().Debug("Executing scripts tool", zap.String("tool", toolName))

	switch toolName {
	case "get_scripts":
		return s.getScripts(ctx, arguments)
	case "get_script_by_id":
		return s.getScriptByID(ctx, arguments)
	case "get_script_by_name":
		return s.getScriptByName(ctx, arguments)
	case "create_script":
		return s.createScript(ctx, arguments)
	case "update_script_by_id":
		return s.updateScriptByID(ctx, arguments)
	case "update_script_by_name":
		return s.updateScriptByName(ctx, arguments)
	case "delete_script_by_id":
		return s.deleteScriptByID(ctx, arguments)
	case "delete_script_by_name":
		return s.deleteScriptByName(ctx, arguments)
	default:
		return "", fmt.Errorf("unknown tool: %s", toolName)
	}
}

// Implementation of script operations

func (s *ScriptsToolset) getScripts(ctx context.Context, args map[string]interface{}) (string, error) {
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

	scripts, err := s.GetClient().GetScripts(params)
	if err != nil {
		return "", fmt.Errorf("failed to get scripts: %w", err)
	}

	response, err := FormatJSONResponse(scripts)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d scripts:\n\n%s", scripts.Size, response), nil
}

func (s *ScriptsToolset) getScriptByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	script, err := s.GetClient().GetScriptByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get script with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(script)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Script details for ID %s:\n\n%s", id, response), nil
}

func (s *ScriptsToolset) getScriptByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	script, err := s.GetClient().GetScriptByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get script with name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(script)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Script details for name '%s':\n\n%s", name, response), nil
}

func (s *ScriptsToolset) createScript(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	scriptContents, err := GetStringArgument(args, "script_contents", true)
	if err != nil {
		return "", err
	}

	// Build the script object
	script := &jamfpro.ResourceScript{
		Name:           name,
		ScriptContents: scriptContents,
	}

	// Optional fields
	if categoryName, _ := GetStringArgument(args, "category_name", false); categoryName != "" {
		script.CategoryName = categoryName
	}
	if categoryID, _ := GetStringArgument(args, "category_id", false); categoryID != "" {
		script.CategoryId = categoryID
	}
	if info, _ := GetStringArgument(args, "info", false); info != "" {
		script.Info = info
	}
	if notes, _ := GetStringArgument(args, "notes", false); notes != "" {
		script.Notes = notes
	}
	if osRequirements, _ := GetStringArgument(args, "os_requirements", false); osRequirements != "" {
		script.OSRequirements = osRequirements
	}
	if priority, _ := GetStringArgument(args, "priority", false); priority != "" {
		script.Priority = priority
	}

	// Parameters 4-11
	if param4, _ := GetStringArgument(args, "parameter_4", false); param4 != "" {
		script.Parameter4 = param4
	}
	if param5, _ := GetStringArgument(args, "parameter_5", false); param5 != "" {
		script.Parameter5 = param5
	}
	if param6, _ := GetStringArgument(args, "parameter_6", false); param6 != "" {
		script.Parameter6 = param6
	}
	if param7, _ := GetStringArgument(args, "parameter_7", false); param7 != "" {
		script.Parameter7 = param7
	}
	if param8, _ := GetStringArgument(args, "parameter_8", false); param8 != "" {
		script.Parameter8 = param8
	}
	if param9, _ := GetStringArgument(args, "parameter_9", false); param9 != "" {
		script.Parameter9 = param9
	}
	if param10, _ := GetStringArgument(args, "parameter_10", false); param10 != "" {
		script.Parameter10 = param10
	}
	if param11, _ := GetStringArgument(args, "parameter_11", false); param11 != "" {
		script.Parameter11 = param11
	}

	result, err := s.GetClient().CreateScript(script)
	if err != nil {
		return "", fmt.Errorf("failed to create script '%s': %w", name, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully created script '%s':\n\n%s", name, response), nil
}

func (s *ScriptsToolset) updateScriptByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	// Build update object with only provided fields
	scriptUpdate := &jamfpro.ResourceScript{}

	if name, _ := GetStringArgument(args, "name", false); name != "" {
		scriptUpdate.Name = name
	}
	if scriptContents, _ := GetStringArgument(args, "script_contents", false); scriptContents != "" {
		scriptUpdate.ScriptContents = scriptContents
	}
	if categoryName, _ := GetStringArgument(args, "category_name", false); categoryName != "" {
		scriptUpdate.CategoryName = categoryName
	}
	if categoryID, _ := GetStringArgument(args, "category_id", false); categoryID != "" {
		scriptUpdate.CategoryId = categoryID
	}
	if info, _ := GetStringArgument(args, "info", false); info != "" {
		scriptUpdate.Info = info
	}
	if notes, _ := GetStringArgument(args, "notes", false); notes != "" {
		scriptUpdate.Notes = notes
	}
	if osRequirements, _ := GetStringArgument(args, "os_requirements", false); osRequirements != "" {
		scriptUpdate.OSRequirements = osRequirements
	}
	if priority, _ := GetStringArgument(args, "priority", false); priority != "" {
		scriptUpdate.Priority = priority
	}

	// Parameters 4-11
	if param4, _ := GetStringArgument(args, "parameter_4", false); param4 != "" {
		scriptUpdate.Parameter4 = param4
	}
	if param5, _ := GetStringArgument(args, "parameter_5", false); param5 != "" {
		scriptUpdate.Parameter5 = param5
	}
	if param6, _ := GetStringArgument(args, "parameter_6", false); param6 != "" {
		scriptUpdate.Parameter6 = param6
	}
	if param7, _ := GetStringArgument(args, "parameter_7", false); param7 != "" {
		scriptUpdate.Parameter7 = param7
	}
	if param8, _ := GetStringArgument(args, "parameter_8", false); param8 != "" {
		scriptUpdate.Parameter8 = param8
	}
	if param9, _ := GetStringArgument(args, "parameter_9", false); param9 != "" {
		scriptUpdate.Parameter9 = param9
	}
	if param10, _ := GetStringArgument(args, "parameter_10", false); param10 != "" {
		scriptUpdate.Parameter10 = param10
	}
	if param11, _ := GetStringArgument(args, "parameter_11", false); param11 != "" {
		scriptUpdate.Parameter11 = param11
	}

	result, err := s.GetClient().UpdateScriptByID(id, scriptUpdate)
	if err != nil {
		return "", fmt.Errorf("failed to update script with ID %s: %w", id, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully updated script with ID %s:\n\n%s", id, response), nil
}

func (s *ScriptsToolset) updateScriptByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	// Build update object with only provided fields
	scriptUpdate := &jamfpro.ResourceScript{}

	// Handle new name if provided
	if newName, _ := GetStringArgument(args, "new_name", false); newName != "" {
		scriptUpdate.Name = newName
	}
	if scriptContents, _ := GetStringArgument(args, "script_contents", false); scriptContents != "" {
		scriptUpdate.ScriptContents = scriptContents
	}
	if categoryName, _ := GetStringArgument(args, "category_name", false); categoryName != "" {
		scriptUpdate.CategoryName = categoryName
	}
	if categoryID, _ := GetStringArgument(args, "category_id", false); categoryID != "" {
		scriptUpdate.CategoryId = categoryID
	}
	if info, _ := GetStringArgument(args, "info", false); info != "" {
		scriptUpdate.Info = info
	}
	if notes, _ := GetStringArgument(args, "notes", false); notes != "" {
		scriptUpdate.Notes = notes
	}
	if osRequirements, _ := GetStringArgument(args, "os_requirements", false); osRequirements != "" {
		scriptUpdate.OSRequirements = osRequirements
	}
	if priority, _ := GetStringArgument(args, "priority", false); priority != "" {
		scriptUpdate.Priority = priority
	}

	// Parameters 4-11
	if param4, _ := GetStringArgument(args, "parameter_4", false); param4 != "" {
		scriptUpdate.Parameter4 = param4
	}
	if param5, _ := GetStringArgument(args, "parameter_5", false); param5 != "" {
		scriptUpdate.Parameter5 = param5
	}
	if param6, _ := GetStringArgument(args, "parameter_6", false); param6 != "" {
		scriptUpdate.Parameter6 = param6
	}
	if param7, _ := GetStringArgument(args, "parameter_7", false); param7 != "" {
		scriptUpdate.Parameter7 = param7
	}
	if param8, _ := GetStringArgument(args, "parameter_8", false); param8 != "" {
		scriptUpdate.Parameter8 = param8
	}
	if param9, _ := GetStringArgument(args, "parameter_9", false); param9 != "" {
		scriptUpdate.Parameter9 = param9
	}
	if param10, _ := GetStringArgument(args, "parameter_10", false); param10 != "" {
		scriptUpdate.Parameter10 = param10
	}
	if param11, _ := GetStringArgument(args, "parameter_11", false); param11 != "" {
		scriptUpdate.Parameter11 = param11
	}

	result, err := s.GetClient().UpdateScriptByName(name, scriptUpdate)
	if err != nil {
		return "", fmt.Errorf("failed to update script with name %s: %w", name, err)
	}

	response, err := FormatJSONResponse(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully updated script with name '%s':\n\n%s", name, response), nil
}

func (s *ScriptsToolset) deleteScriptByID(ctx context.Context, args map[string]interface{}) (string, error) {
	id, err := GetStringArgument(args, "id", true)
	if err != nil {
		return "", err
	}

	err = s.GetClient().DeleteScriptByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to delete script with ID %s: %w", id, err)
	}

	return fmt.Sprintf("Successfully deleted script with ID %s", id), nil
}

func (s *ScriptsToolset) deleteScriptByName(ctx context.Context, args map[string]interface{}) (string, error) {
	name, err := GetStringArgument(args, "name", true)
	if err != nil {
		return "", err
	}

	err = s.GetClient().DeleteScriptByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to delete script with name %s: %w", name, err)
	}

	return fmt.Sprintf("Successfully deleted script with name '%s'", name), nil
}
