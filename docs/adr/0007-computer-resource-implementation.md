# ADR 0007: Computer Resource Implementation

## Status

Accepted

## Context

The Jamf Pro MCP Server needs to provide tools for interacting with computers managed by Jamf Pro. Computers are a core resource in the Jamf Pro ecosystem, allowing organizations to manage macOS devices, track inventory, deploy configurations, and monitor compliance.

The Jamf Pro SDK provides access to computers through the Classic API, with the following key capabilities:
- Listing all computers
- Retrieving detailed information about specific computers
- Creating new computer records
- Updating existing computer records
- Deleting computer records

When implementing the computer toolset, we need to consider the following factors:
1. The complex and extensive data model for computers, which includes numerous nested objects and arrays
2. The required fields for creating and updating computers
3. The need for LLM models to understand the structure and relationships of the data
4. Balancing completeness with usability in the API surface

## Decision

We will implement a comprehensive computer toolset with the following features:

1. **Full Data Model Support**: We will expose the complete data model through our toolset, including all major sections:
   - General information (name, identifiers, network details)
   - Location information (user, department, building)
   - Purchasing information (PO details, warranty, lease)
   - Hardware details (model, specifications, storage)
   - Security information
   - Software inventory

2. **Reference Template**: We will provide a template function and tool that LLMs can use to understand the structure of a computer resource. This will help them generate valid requests for creating and updating computers.

3. **Field Validation**: We will validate required fields during creation to ensure that API calls have the necessary data to succeed. For computer creation, this includes:
   - Name (required)
   - Other key identifiers (serial number, UDID) should be strongly encouraged

4. **Comprehensive Update Capabilities**: Our update function will support modifying all fields in the computer resource, with the ability to update by ID or name.

5. **Consistent Error Handling**: We will provide clear error messages that help identify issues with requests, especially related to required fields or API limitations.

6. **MCP Protocol Alignment**: All tools will follow the MCP protocol standard, with clear input schemas and documentation that helps LLMs understand what each tool does and what parameters it accepts.

## Consequences

### Positive

- LLMs will have comprehensive access to all computer management features in Jamf Pro
- The reference template will make it easier for LLMs to understand the structure of computer resources
- Required field validation will reduce errors in API calls
- Clear error messages will help troubleshoot issues
- The toolset will be extensible for future enhancements

### Negative

- The extensive data model increases complexity
- Some fields in the data model are read-only and cannot be modified through the API
- The Classic API has some limitations compared to newer API designs
- The sheer number of fields and nested objects can be overwhelming

## Implementation Notes

The computer toolset implementation includes:

1. **Tools for CRUD Operations**:
   - `get_computers`: List all computers
   - `get_computer_by_id` / `get_computer_by_name`: Get specific computers
   - `create_computer`: Create new computers
   - `update_computer_by_id` / `update_computer_by_name`: Update existing computers
   - `delete_computer_by_id` / `delete_computer_by_name`: Delete computers

2. **Reference Tool**:
   - `get_computer_template`: Provides a complete example of a computer resource

The input schemas for creation and update operations include the most common fields with clear descriptions of each field's purpose, and indicate which fields are required versus optional.

This implementation aligns with our decision in ADR-0005 to provide reference templates for complex resources, making it easier for LLMs to understand and work with the Jamf Pro API.

The computer resource implementation is particularly important as it represents one of the most complex and extensively used resources in the Jamf Pro ecosystem. 