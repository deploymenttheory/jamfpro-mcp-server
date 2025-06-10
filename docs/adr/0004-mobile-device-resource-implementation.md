# ADR 0004: Mobile Device Resource Implementation

## Status

Accepted

## Context

The Jamf Pro MCP Server needs to provide tools for interacting with mobile devices managed by Jamf Pro. Mobile devices represent an important resource in the Jamf Pro ecosystem, allowing organizations to manage iOS, iPadOS, and tvOS devices.

The Jamf Pro SDK provides access to mobile devices through the Classic API, with the following key capabilities:
- Listing all mobile devices
- Retrieving detailed information about specific devices
- Creating new mobile device records
- Updating existing mobile device records
- Deleting mobile device records
- Accessing related resources like mobile device groups, applications, and configuration profiles

When implementing the mobile device toolset, we need to consider the following factors:
1. The complex and extensive data model for mobile devices
2. The required fields for creating and updating mobile devices
3. The need for LLM models to understand the structure and relationships of the data
4. The limitations of the Classic API for mobile devices

## Decision

We will implement a comprehensive mobile devices toolset with the following features:

1. **Full Data Model Support**: We will expose the complete data model through our toolset, including all fields in:
   - General device information
   - Location information
   - Purchasing information
   - Related resources (applications, certificates, configuration profiles, etc.)

2. **Reference Template**: We will provide a template function and tool that LLMs can use to understand the structure of a mobile device resource. This will help them generate valid requests for creating and updating mobile devices.

3. **Field Validation**: We will validate required fields during creation and updates to ensure that API calls have the necessary data to succeed. For mobile device creation, this includes:
   - Name
   - Serial Number
   - UDID

4. **Comprehensive Update Capabilities**: Our update function will support modifying all fields in the mobile device resource, including nested fields in location and purchasing sections.

5. **Clear Error Handling**: We will provide clear error messages that help identify issues with requests, especially related to required fields or API limitations.

6. **MCP Protocol Alignment**: All tools will follow the MCP protocol standard, with clear input schemas and documentation that helps LLMs understand what each tool does and what parameters it accepts.

## Consequences

### Positive

- LLMs will have comprehensive access to all mobile device features in Jamf Pro
- The reference template will make it easier for LLMs to understand the structure of mobile device resources
- Required field validation will reduce errors in API calls
- Clear error messages will help troubleshoot issues
- The toolset will be extensible for future enhancements

### Negative

- The extensive data model increases complexity
- Some fields might not be modifiable through the API despite being exposed in the data model
- The Classic API has some limitations compared to the newer Pro API

## Implementation Notes

The mobile device toolset implementation includes:

1. **Tools for CRUD Operations**:
   - `get_mobile_devices`: List all mobile devices
   - `get_mobile_device_by_id` / `get_mobile_device_by_name`: Get specific devices
   - `create_mobile_device`: Create new devices
   - `update_mobile_device_by_id`: Update existing devices
   - `delete_mobile_device`: Delete devices

2. **Tools for Related Resources**:
   - `get_mobile_device_groups`: Get all device groups
   - `get_mobile_device_group_by_id`: Get a specific device group
   - `get_mobile_device_applications`: Get all mobile applications
   - `get_mobile_device_configuration_profiles`: Get all configuration profiles

3. **Reference Tool**:
   - `get_mobile_device_template`: Provides a complete example of a mobile device resource

The input schemas for creation and update operations include all available fields, categorized by their section in the data model (General, Location, Purchasing), with clear descriptions of each field's purpose. 