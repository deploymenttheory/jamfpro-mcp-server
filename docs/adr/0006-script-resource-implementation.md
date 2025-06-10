# ADR 0006: Script Resource Implementation

## Status

Accepted

## Context

The Jamf Pro MCP Server needs to provide tools for interacting with scripts managed by Jamf Pro. Scripts are an important resource in the Jamf Pro ecosystem, allowing organizations to automate tasks, perform system configurations, and execute custom code on managed devices.

The Jamf Pro SDK provides access to scripts through the Pro API, with the following key capabilities:
- Listing all scripts with pagination, sorting, and filtering
- Retrieving detailed information about specific scripts
- Creating new script resources
- Updating existing script resources
- Deleting script resources

When implementing the script toolset, we need to consider the following factors:
1. The data model for scripts, including all available fields
2. The required fields for creating and updating scripts
3. The need for LLM models to understand the structure of the script resource
4. The capabilities and limitations of the Pro API for scripts

## Decision

We will implement a comprehensive script toolset with the following features:

1. **Full Data Model Support**: We will expose the complete data model through our toolset, including all fields in the script resource:
   - Basic information (name, category, info, notes)
   - Execution details (OS requirements, priority)
   - Script contents
   - Script parameters (4-11)

2. **Reference Template**: We will provide a template function and tool that LLMs can use to understand the structure of a script resource. This will help them generate valid requests for creating and updating scripts.

3. **Field Validation**: We will validate required fields during creation to ensure that API calls have the necessary data to succeed. For script creation, this includes:
   - Name
   - Script contents

4. **Comprehensive Update Capabilities**: Our update function will support modifying all fields in the script resource, with the ability to update by ID or name.

5. **Pagination and Filtering Support**: When listing scripts, we will support pagination, sorting, and filtering to enable efficient retrieval of script resources.

6. **MCP Protocol Alignment**: All tools will follow the MCP protocol standard, with clear input schemas and documentation that helps LLMs understand what each tool does and what parameters it accepts.

## Consequences

### Positive

- LLMs will have comprehensive access to all script features in Jamf Pro
- The reference template will make it easier for LLMs to understand the structure of script resources
- Required field validation will reduce errors in API calls
- Support for pagination, sorting, and filtering will enable efficient resource management
- The toolset will be extensible for future enhancements

### Negative

- The Pro API has some limitations compared to the Classic API in certain edge cases
- Some fields may have specific validation requirements that aren't immediately obvious

## Implementation Notes

The script toolset implementation includes:

1. **Tools for CRUD Operations**:
   - `get_scripts`: List all scripts with pagination, sorting, and filtering
   - `get_script_by_id` / `get_script_by_name`: Get specific scripts
   - `create_script`: Create new scripts
   - `update_script_by_id` / `update_script_by_name`: Update existing scripts
   - `delete_script_by_id` / `delete_script_by_name`: Delete scripts

2. **Reference Tool**:
   - `get_script_template`: Provides a complete example of a script resource

The input schemas for creation and update operations include all available fields with clear descriptions of each field's purpose, and indicate which fields are required versus optional.

This implementation aligns with our decision in ADR-0005 to provide reference templates for complex resources, making it easier for LLMs to understand and work with the Jamf Pro API. 