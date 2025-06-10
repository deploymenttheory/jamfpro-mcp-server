# ADR 0001: MCP Server Architecture for Jamf Pro

## Status

Accepted

## Context

The Jamf Pro MCP (Model Context Protocol) Server needs to provide a standardized way for AI tools to interact with Jamf Pro APIs. The architecture should align with GitHub's MCP Server design while accommodating the specific needs of Jamf Pro's API structure.

The GitHub MCP Server implementation serves as a reference architecture, offering a tested approach to creating a Model Context Protocol server. However, we need to adapt this architecture to Jamf Pro's unique API structure, which includes both Classic API and Pro API endpoints.

## Decision

We will implement a Jamf Pro MCP Server that follows the core architectural patterns of GitHub's MCP Server while making appropriate adaptations for Jamf Pro:

1. **Protocol Implementation**:
   - Implement the standard MCP protocol for communication
   - Support JSON-RPC 2.0 format for requests and responses

2. **Toolset Organization**:
   - Organize functionality into logical toolsets based on Jamf Pro resource types
   - Each toolset will be a separate module focusing on a specific area of Jamf Pro (e.g., computers, mobile devices, scripts)

3. **API Integration**:
   - Utilize the go-api-sdk-jamfpro as the underlying SDK for all API interactions
   - Support both Classic API and Pro API endpoints as needed
   - Abstract the API complexity away from the MCP protocol layer

4. **Server Structure**:
   - Use a similar server structure to GitHub's MCP Server
   - Implement a main server component that handles MCP protocol messages
   - Process requests through appropriate toolsets based on the requested tool
   - Return standardized responses according to MCP protocol

5. **Authentication**:
   - Support both OAuth2 and Basic authentication methods for Jamf Pro
   - Handle authentication token management and refresh

6. **Tool Definition**:
   - Define tools with clear schemas for input parameters
   - Provide descriptive documentation for each tool

## Consequences

### Positive

- Clear separation of concerns between protocol handling and API operations
- Consistent interface for AI tools to interact with Jamf Pro
- Reuse of proven architectural patterns from GitHub's MCP Server
- Flexible structure that can accommodate both current and future Jamf Pro API endpoints
- Ability to add new toolsets and tools without changing core server architecture

### Negative

- Some complexity in handling different API styles (Classic vs Pro)
- Need to maintain compatibility with both MCP protocol standards and Jamf Pro API changes
- Potential performance considerations when translating between MCP protocol and Jamf Pro API

## Implementation Notes

The implementation will be structured as follows:

- **Server Package**: Handles the core MCP server functionality
- **MCP Package**: Implements the Model Context Protocol
- **Toolsets Package**: Contains individual toolsets for different Jamf Pro resource types
- **Config Package**: Manages configuration and settings

Each toolset will implement a common interface, allowing the server to delegate tool execution to the appropriate toolset based on the requested tool name. 