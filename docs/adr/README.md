# Architectural Decision Records (ADRs)

This directory contains Architectural Decision Records (ADRs) for the Jamf Pro MCP Server project.

## What are ADRs?

Architectural Decision Records are documents that capture important architectural decisions made in a project. Each ADR describes:

- The context (the problem being addressed)
- The decision (the solution chosen)
- The consequences (both positive and negative)

By documenting these decisions, we create a historical record that helps team members understand why certain approaches were chosen and the trade-offs involved.

## ADR List

| ID | Title | Status |
|----|-------|--------|
| [0001](0001-mcp-server-architecture.md) | MCP Server Architecture for Jamf Pro | Accepted |
| [0002](0002-toolset-resource-implementation.md) | Toolset and Resource Implementation Strategy | Accepted |
| [0003](0003-resource-selection-and-api-preference.md) | Resource Selection and API Preference | Accepted |
| [0004](0004-mobile-device-resource-implementation.md) | Mobile Device Resource Implementation | Accepted |
| [0005](0005-data-model-reference-templates.md) | Data Model Reference Templates for LLMs | Accepted |
| [0006](0006-script-resource-implementation.md) | Script Resource Implementation | Accepted |
| [0007](0007-computer-resource-implementation.md) | Computer Resource Implementation | Accepted |

## ADR Process

When making significant architectural decisions for the Jamf Pro MCP Server, please follow this process:

1. Create a new ADR file using the sequential numbering pattern
2. Use the format: `####-title-with-hyphens.md`
3. Follow the template structure:
   - Title
   - Status (Proposed, Accepted, Superseded, etc.)
   - Context
   - Decision
   - Consequences (positive and negative)
   - Implementation Notes (optional)
4. Reference related ADRs where appropriate
5. Submit the ADR as part of the related pull request

## ADR Statuses

- **Proposed**: A decision has been proposed but not yet reviewed or accepted
- **Accepted**: The decision has been accepted and is now in effect
- **Superseded**: The decision has been replaced by a newer decision (reference the newer ADR)
- **Deprecated**: The decision is no longer relevant but has not been explicitly replaced
- **Rejected**: The proposed decision was rejected and an alternative approach was chosen 