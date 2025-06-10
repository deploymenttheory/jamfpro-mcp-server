# ADR 0002: Toolset and Resource Implementation Strategy

## Status

Accepted

## Context

The Jamf Pro MCP Server needs to provide tools for interacting with various Jamf Pro resources. These resources include computers, mobile devices, scripts, and computer inventory. We need to determine the best approach for implementing these tools while maintaining consistency, extensibility, and adherence to the MCP protocol standards.

Jamf Pro offers two distinct API styles:
1. The Classic API (XML-based)
2. The Pro API (JSON-based)

Each resource type may be available through one or both of these APIs, with varying capabilities and data models.

## Decision

We will implement a consistent toolset pattern for all Jamf Pro resources with the following approach:

1. **Toolset Structure**:
   - Each resource type will have its own dedicated toolset (e.g., ComputersToolset, MobileDevicesToolset)
   - All toolsets will implement the same base interface (Toolset)
   - Toolsets will extend a BaseToolset that provides common functionality

2. **Tool Definition**:
   - Each tool will have a clear name, description, and input schema
   - Input schemas will be defined according to JSON Schema standards
   - Required and optional parameters will be clearly specified

3. **API Preference**:
   - When a resource is available in both Classic and Pro APIs, we will prefer the Pro API implementation
   - The Classic API will be used only when necessary functionality is not available in the Pro API
   - The API selection will be abstracted away from the tool user

4. **Resource Operations**:
   - Standard CRUD operations will be implemented for each resource type
   - Additional specialized operations will be added as needed
   - Operations will be mapped to appropriate SDK functions

5. **Response Formatting**:
   - All responses will be formatted consistently using the MCP protocol's content structure
   - Errors will be handled uniformly across all toolsets
   - Response data will be sanitized to remove sensitive information

6. **Implementation Priority**:
   - Focus first on core resources: Computers, Computer Inventory, Scripts, and Mobile Devices
   - Additional resources will be added incrementally

## Consequences

### Positive

- Consistent interface for all resource types
- Clear separation between tool definition and implementation
- Flexibility to choose the most appropriate API for each operation
- Extensible design allowing new tools to be added easily
- Uniform error handling and response formatting

### Negative

- Some duplication in tool definitions across similar resources
- Need to maintain mapping between tool inputs and SDK parameters
- Potential complexity in handling differences between API versions
- Challenge in representing complex nested data structures in tool schemas

## Implementation Notes

For the initial implementation, we will focus on the following resources:

1. **Computers**: Basic computer management operations using the Classic API
2. **Computer Inventory**: Detailed inventory management using the Pro API
3. **Scripts**: Script management using the Pro API
4. **Mobile Devices**: Mobile device management using the Classic API

Each toolset will follow this general structure:
- Define the available tools and their schemas
- Implement tool execution by mapping to SDK calls
- Format responses according to MCP protocol standards
- Handle errors appropriately

We will use a factory pattern to create the appropriate toolset based on configuration options, allowing for dynamic toolset selection and initialization. 