# ADR 0005: Data Model Reference Templates for LLMs

## Status

Accepted

## Context

When working with complex Jamf Pro resources through the MCP Server, Large Language Models (LLMs) need to understand the structure and relationships of the data models they're manipulating. This is particularly important for resources with deep and complex structures like mobile devices, computers, and other Jamf Pro entities.

LLMs perform best when they have concrete examples and references to work with, rather than having to infer structure from documentation alone. However, without explicit examples, LLMs may:

1. Miss required fields when creating resources
2. Misunderstand the nesting structure of complex objects
3. Use incorrect field names or data types
4. Fail to leverage the full capabilities of the API

## Decision

We will implement a "reference template" pattern for complex resources in the Jamf Pro MCP Server. This pattern includes:

1. **Template Getter Methods**: Each toolset for complex resources will include a method (e.g., `GetMobileDeviceTemplate()`) that returns a fully populated example of the resource.

2. **Template Reference Tools**: We will expose these templates through MCP tools (e.g., `get_mobile_device_template`) that return formatted JSON examples of the resources.

3. **Comprehensive Coverage**: Templates will include all major sections and fields of the resource, with realistic example values that demonstrate proper formatting.

4. **Documentation Comments**: Templates will be well-commented to explain the purpose and relationships of different sections.

5. **Consistent Naming**: All template-related methods and tools will follow a consistent naming pattern for discoverability.

This approach provides LLMs with concrete references when they need to generate code for creating or updating complex resources.

## Consequences

### Positive

- LLMs will have clear, concrete examples of each resource structure
- Resource creation and updates will be more likely to succeed on the first attempt
- The templates serve as implicit documentation for each resource
- Developers and users can refer to these templates when working with the API
- Reduces the need for trial-and-error when working with complex resources

### Negative

- Maintaining template examples requires updating them when the underlying data models change
- Templates increase the code surface area of each toolset
- There's a risk of the templates becoming outdated if not maintained alongside data model changes

## Implementation Notes

We have implemented this pattern for the mobile devices resource, providing:

1. A `GetMobileDeviceTemplate()` method that returns a fully populated example
2. A `get_mobile_device_template` tool that exposes this to LLMs through the MCP protocol
3. Comprehensive coverage of the General, Location, and Purchasing sections

We will extend this pattern to other complex resources like computers, scripts, and policies.

For each resource, the template will demonstrate:
- Required fields with appropriate example values
- Optional fields with representative values
- Proper nesting of sub-objects
- Appropriate data types for each field

We will document this pattern so that future resource implementations follow the same approach. 