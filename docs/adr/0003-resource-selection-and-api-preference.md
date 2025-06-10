# ADR 0003: Resource Selection and API Preference

## Status

Accepted

## Context

The Jamf Pro ecosystem offers access to various resources through two distinct API styles:

1. **Classic API** - XML-based, older but with broader coverage of resources
2. **Pro API** - JSON-based, newer with improved design but not covering all resources

We need to decide which resources to prioritize for implementation in the MCP server and determine which API style to use for each resource. This decision affects the capabilities, performance, and maintainability of the MCP server.

## Decision

We will implement the following core resources in the initial release:

1. **Computers**
   - API: Classic API
   - Rationale: While available in both APIs, the Classic API provides more comprehensive management capabilities.

2. **Computer Inventory**
   - API: Pro API
   - Rationale: The Pro API provides a more structured and detailed inventory model with better filtering and section selection capabilities.

3. **Scripts**
   - API: Pro API
   - Rationale: The Pro API offers a cleaner interface for script management with better parameter handling.

4. **Mobile Devices**
   - API: Classic API
   - Rationale: More comprehensive mobile device management capabilities in the Classic API.

For API selection in general, we will follow these guidelines:

1. **Feature Completeness**: Choose the API that provides the most complete set of operations for the resource.

2. **Data Model Quality**: Prefer APIs with clearer, more consistent data models.

3. **Performance**: Consider API performance, especially for operations that may involve large datasets.

4. **Future Direction**: Give preference to the Pro API where capabilities are equivalent, as it represents Jamf's strategic direction.

5. **Client Needs**: Prioritize the API that best meets the anticipated needs of MCP clients.

## Consequences

### Positive

- Optimized API selection for each resource type
- Better performance and capabilities by leveraging the strengths of each API
- Future-proofing by preferring Pro API where appropriate
- Comprehensive coverage of core Jamf Pro functionality

### Negative

- Need to maintain code that works with both API styles
- Potential inconsistency in behavior between resources using different APIs
- Migration effort required if we need to switch API styles for a resource in the future
- Learning curve for developers who need to understand both API styles

## Implementation Notes

The implementation will abstract away the differences between API styles from the MCP client perspective. Each toolset will handle the complexities of working with its selected API internally, providing a consistent interface regardless of the underlying API used.

For resources with partial implementations in the Pro API, we may need to use a hybrid approach, calling different APIs for different operations on the same resource type. In such cases, we will carefully document this behavior and ensure consistent responses regardless of the API used. 