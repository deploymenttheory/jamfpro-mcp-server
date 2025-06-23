# Template

This repository serves as a **Default Template Repository** according official [GitHub Contributing Guidelines][ProjectSetup] for healthy contributions. It brings you clean default Templates for several areas:

- [Azure DevOps Pull Requests](.azuredevops/PULL_REQUEST_TEMPLATE.md) ([`.azuredevops\PULL_REQUEST_TEMPLATE.md`](`.azuredevops\PULL_REQUEST_TEMPLATE.md`))
- [Azure Pipelines](.pipelines/pipeline.yml) ([`.pipelines/pipeline.yml`](`.pipelines/pipeline.yml`))
- [GitHub Workflows](.github/workflows/)
  - [Super Linter](.github/workflows/linter.yml) ([`.github/workflows/linter.yml`](`.github/workflows/linter.yml`))
  - [Sample Workflows](.github/workflows/workflow.yml) ([`.github/workflows/workflow.yml`](`.github/workflows/workflow.yml`))
- [GitHub Pull Requests](.github/PULL_REQUEST_TEMPLATE.md) ([`.github/PULL_REQUEST_TEMPLATE.md`](`.github/PULL_REQUEST_TEMPLATE.md`))
- [GitHub Issues](.github/ISSUE_TEMPLATE/)
  - [Feature Requests](.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md) ([`.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md`](`.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md`))
  - [Bug Reports](.github/ISSUE_TEMPLATE/BUG_REPORT.md) ([`.github/ISSUE_TEMPLATE/BUG_REPORT.md`](`.github/ISSUE_TEMPLATE/BUG_REPORT.md`))
- [Codeowners](.github/CODEOWNERS) ([`.github/CODEOWNERS`](`.github/CODEOWNERS`)) _adjust usernames once cloned_
- [Wiki and Documentation](docs/) ([`docs/`](`docs/`))
- [gitignore](.gitignore) ([`.gitignore`](.gitignore))
- [gitattributes](.gitattributes) ([`.gitattributes`](.gitattributes))
- [Changelog](CHANGELOG.md) ([`CHANGELOG.md`](`CHANGELOG.md`))
- [Code of Conduct](CODE_OF_CONDUCT.md) ([`CODE_OF_CONDUCT.md`](`CODE_OF_CONDUCT.md`))
- [Contribution](CONTRIBUTING.md) ([`CONTRIBUTING.md`](`CONTRIBUTING.md`))
- [License](LICENSE) ([`LICENSE`](`LICENSE`)) _adjust projectname once cloned_
- [Readme](README.md) ([`README.md`](`README.md`))
- [Security](SECURITY.md) ([`SECURITY.md`](`SECURITY.md`))


## Status

[![Super Linter](<https://github.com/segraef/Template/actions/workflows/linter.yml/badge.svg>)](<https://github.com/segraef/Template/actions/workflows/linter.yml>)

[![Sample Workflow](<https://github.com/segraef/Template/actions/workflows/workflow.yml/badge.svg>)](<https://github.com/segraef/Template/actions/workflows/workflow.yml>)

## Creating a repository from a template

You can [generate](https://github.com/segraef/Template/generate) a new repository with the same directory structure and files as an existing repository. More details can be found [here][CreateFromTemplate].

## Reporting Issues and Feedback

### Issues and Bugs

If you find any bugs, please file an issue in the [GitHub Issues][GitHubIssues] page. Please fill out the provided template with the appropriate information.

If you are taking the time to mention a problem, even a seemingly minor one, it is greatly appreciated, and a totally valid contribution to this project. **Thank you!**

## Feedback

If there is a feature you would like to see in here, please file an issue or feature request in the [GitHub Issues][GitHubIssues] page to provide direct feedback.

## Contribution

If you would like to become an active contributor to this repository or project, please follow the instructions provided in [`CONTRIBUTING.md`][Contributing].


# Jamf Pro MCP Server

The Jamf Pro MCP Server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction)
server that provides seamless integration with Jamf Pro APIs, enabling advanced
automation and interaction capabilities for AI tools and applications.

### Use Cases

- Automating Jamf Pro workflows and processes
- Extracting and analyzing data from Jamf Pro
- Building AI-powered tools that interact with the Jamf Pro ecosystem
- Managing Apple devices through AI assistants

---

## Local Jamf Pro MCP Server

[![Install with Docker in VS Code](https://img.shields.io/badge/VS_Code-Install_Server-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=jamfpro&inputs=%5B%7B%22id%22%3A%22jamf_client_id%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22Jamf%20Pro%20Client%20ID%22%2C%22password%22%3Atrue%7D%2C%7B%22id%22%3A%22jamf_client_secret%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22Jamf%20Pro%20Client%20Secret%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22JAMF_INSTANCE_URL%3Dhttps%3A%2F%2Fyour-instance.jamfcloud.com%22%2C%22-e%22%2C%22JAMF_CLIENT_ID%3D%24%7Binput%3Ajamf_client_id%7D%22%2C%22-e%22%2C%22JAMF_CLIENT_SECRET%3D%24%7Binput%3Ajamf_client_secret%7D%22%2C%22ghcr.io%2Fdeploymenttheory%2Fjamfpro-mcp-server%3Alatest%22%5D%2C%22env%22%3A%7B%22JAMF_CLIENT_ID%22%3A%22%24%7Binput%3Ajamf_client_id%7D%22%2C%22JAMF_CLIENT_SECRET%22%3A%22%24%7Binput%3Ajamf_client_secret%7D%22%7D%7D) 

## Prerequisites

1. To run the server in a container, you will need to have [Docker](https://www.docker.com/) installed.
2. Once Docker is installed, you will also need to ensure Docker is running.
3. You will need a Jamf Pro instance with API credentials:
   - For OAuth authentication: Client ID and Client Secret
   - For Basic authentication: Username and Password

## API Integration

The Jamf Pro MCP Server integrates with both the [Jamf Pro API](https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-overview) and the Classic API. It leverages the [go-api-sdk-jamfpro](https://github.com/deploymenttheory/go-api-sdk-jamfpro) library to provide comprehensive access to Jamf Pro's functionality.

### API Features

- **RESTful Interface**: Uses standard HTTP methods (GET, POST, PUT, DELETE, PATCH) to interact with Jamf Pro resources
- **JSON Data Format**: Most endpoints use JSON for data exchange
- **Authentication**: Supports both OAuth2 and Basic authentication methods
- **Comprehensive Coverage**: Access to over 435 operations across both the Jamf Pro API and Classic API

### API Authentication

The MCP server supports two authentication methods for connecting to Jamf Pro:

#### OAuth2 Authentication (Recommended)

OAuth2 is the recommended authentication method for Jamf Pro API access. It provides better security through token-based authentication and supports automatic token refresh.

To set up OAuth2 authentication:

1. In Jamf Pro, navigate to Settings > API roles and clients > 
2. Create an api role and delegate the appropriate permissions to the role
3. Create a new API Client and assign the asforemented api role to the client
4. Note the Client ID and Client Secret for use with the MCP server

#### Basic Authentication

Basic authentication uses standard Jamf Pro username and password credentials. While simpler to set up, it's less secure and doesn't support advanced features like token refresh.

## Installation

### Usage with VS Code

For quick installation, use one of the one-click install buttons. Once you complete that flow, toggle Agent mode (located by the Copilot Chat text input) and the server will start.

### Usage in other MCP Hosts

Add the following JSON block to your IDE MCP settings:

```json
{
  "mcpServers": {
    "jamfpro": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e", "JAMF_INSTANCE_URL=https://your-instance.jamfcloud.com",
        "-e", "JAMF_CLIENT_ID=${input:jamf_client_id}",
        "-e", "JAMF_CLIENT_SECRET=${input:jamf_client_secret}",
        "ghcr.io/deploymenttheory/jamfpro-mcp-server:latest"
      ],
      "env": {
        "JAMF_CLIENT_ID": "${input:jamf_client_id}",
        "JAMF_CLIENT_SECRET": "${input:jamf_client_secret}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "jamf_client_id",
      "description": "Jamf Pro Client ID",
      "password": true
    },
    {
      "type": "promptString",
      "id": "jamf_client_secret",
      "description": "Jamf Pro Client Secret",
      "password": true
    }
  ]
}
```

Optionally, you can add a similar configuration to a file called `.vscode/mcp.json` in your workspace to share the configuration with others.

### Build from source

If you don't want to use Docker, you can build the binary directly:

```bash
go build -o jamfpro-mcp-server cmd/jamfpro-mcp-server/main.go
```

Then run it with:

```bash
./jamfpro-mcp-server --jamf-instance-url=https://your-instance.jamfcloud.com --jamf-client-id=your-client-id --jamf-client-secret=your-client-secret
```

## Tool Configuration

The Jamf Pro MCP Server supports enabling or disabling specific groups of functionalities via the `--toolsets` flag. This allows you to control which Jamf Pro API capabilities are available to your AI tools.

### Available Toolsets

The following sets of tools are available (all are on by default):

| Toolset                 | Description                                                   |
| ----------------------- | ------------------------------------------------------------- |
| `computers`             | Computer management and inventory operations                  |
| `mobile-devices`        | Mobile device management operations                           |
| `policies`              | Policy management operations                                  |
| `users`                 | User management operations                                    |
| `groups`                | Group management operations                                   |
| `configuration-profiles`| Configuration profile operations                              |
| `scripts`               | Script management operations                                  |
| `buildings`             | Building management operations                                |
| `departments`           | Department management operations                              |
| `categories`            | Category management operations                                |
| `sites`                 | Site management operations                                    |
| `inventory`             | Inventory collection operations                               |
| `all`                   | Enable all available toolsets                                 |

#### Specifying Toolsets

To specify toolsets you want available, you can pass an allow-list in two ways:

1. **Using Command Line Argument**:

   ```bash
   jamfpro-mcp-server --toolsets computers,mobile-devices,policies
   ```

2. **Using Environment Variable**:
   ```bash
   JAMF_TOOLSETS="computers,mobile-devices,policies" ./jamfpro-mcp-server
   ```

### Using Toolsets With Docker

When using Docker, you can pass the toolsets as environment variables:

```bash
docker run -i --rm \
  -e JAMF_INSTANCE_URL=https://your-instance.jamfcloud.com \
  -e JAMF_CLIENT_ID=your-client-id \
  -e JAMF_CLIENT_SECRET=your-client-secret \
  -e JAMF_TOOLSETS="computers,mobile-devices,policies" \
  ghcr.io/deploymenttheory/jamfpro-mcp-server:latest
```

## Dynamic Tool Discovery

The MCP server supports dynamic toolset discovery. This allows the MCP host to list and enable toolsets in response to a user prompt, which helps avoid overwhelming the model with too many tools at once.

### Using Dynamic Tool Discovery

When using the binary, you can pass the `--dynamic-toolsets` flag:

```bash
./jamfpro-mcp-server --dynamic-toolsets
```

When using Docker, you can pass the flag as an environment variable:

```bash
docker run -i --rm \
  -e JAMF_INSTANCE_URL=https://your-instance.jamfcloud.com \
  -e JAMF_CLIENT_ID=your-client-id \
  -e JAMF_CLIENT_SECRET=your-client-secret \
  -e JAMF_DYNAMIC_TOOLSETS=1 \
  ghcr.io/deploymenttheory/jamfpro-mcp-server:latest
```

## Authentication Methods

The Jamf Pro MCP Server supports two authentication methods:

### OAuth2 Authentication (Recommended)

```bash
./jamfpro-mcp-server \
  --jamf-instance-url=https://your-instance.jamfcloud.com \
  --jamf-client-id=your-client-id \
  --jamf-client-secret=your-client-secret
```

### Basic Authentication

```bash
./jamfpro-mcp-server \
  --jamf-instance-url=https://your-instance.jamfcloud.com \
  --auth-method=basic \
  --jamf-username=your-username \
  --jamf-password=your-password
```

## i18n / Overriding Descriptions

The descriptions of the tools can be overridden by creating a
`jamfpro-mcp-server-config.json` file in the same directory as the binary.

The file should contain a JSON object with the tool names as keys and the new
descriptions as values. For example:

```json
{
  "TOOL_GET_COMPUTER_DESCRIPTION": "Get detailed information about a computer",
  "TOOL_CREATE_POLICY_DESCRIPTION": "Create a new policy in Jamf Pro"
}
```

You can create an export of the current translations by running the binary with
the `--export-translations` flag.

## Available Tools

### Computers

- **get_computer** - Get detailed information about a computer
  - `id`: Computer ID (number, required)

- **list_computers** - List all computers
  - `page`: Page number (number, optional)
  - `page_size`: Results per page (number, optional)

- **update_computer** - Update computer information
  - `id`: Computer ID (number, required)
  - `name`: New computer name (string, optional)
  - `serial_number`: Serial number (string, optional)
  - `udid`: UDID (string, optional)

- **delete_computer** - Delete a computer
  - `id`: Computer ID (number, required)

### Mobile Devices

- **get_mobile_device** - Get detailed information about a mobile device
  - `id`: Mobile device ID (number, required)

- **list_mobile_devices** - List all mobile devices
  - `page`: Page number (number, optional)
  - `page_size`: Results per page (number, optional)

- **update_mobile_device** - Update mobile device information
  - `id`: Mobile device ID (number, required)
  - `name`: New device name (string, optional)
  - `asset_tag`: Asset tag (string, optional)

- **delete_mobile_device** - Delete a mobile device
  - `id`: Mobile device ID (number, required)

### Policies

- **get_policy** - Get detailed information about a policy
  - `id`: Policy ID (number, required)

- **list_policies** - List all policies
  - `page`: Page number (number, optional)
  - `page_size`: Results per page (number, optional)

- **create_policy** - Create a new policy
  - `name`: Policy name (string, required)
  - `enabled`: Whether the policy is enabled (boolean, optional)
  - `trigger`: Trigger type (string, optional)
  - `frequency`: How often to execute the policy (string, optional)

- **update_policy** - Update a policy
  - `id`: Policy ID (number, required)
  - `name`: New policy name (string, optional)
  - `enabled`: Whether the policy is enabled (boolean, optional)

- **delete_policy** - Delete a policy
  - `id`: Policy ID (number, required)

### Scripts

- **get_script** - Get detailed information about a script
  - `id`: Script ID (number, required)

- **list_scripts** - List all scripts
  - `page`: Page number (number, optional)
  - `page_size`: Results per page (number, optional)

- **create_script** - Create a new script
  - `name`: Script name (string, required)
  - `contents`: Script contents (string, required)
  - `notes`: Script notes (string, optional)
  - `priority`: Script priority (string, optional)

- **update_script** - Update a script
  - `id`: Script ID (number, required)
  - `name`: New script name (string, optional)
  - `contents`: New script contents (string, optional)

- **delete_script** - Delete a script
  - `id`: Script ID (number, required)

## Resources

### Computer Inventory

- **Get Computer Inventory**
  Retrieves inventory information for all computers.

  - **Template**: `jamfpro://computers/inventory`

- **Get Computer Inventory by ID**
  Retrieves inventory information for a specific computer by ID.

  - **Template**: `jamfpro://computers/inventory/{id}`
  - **Parameters**:
    - `id`: Computer ID (number, required)

- **Get Computer Inventory by Name**
  Retrieves inventory information for a specific computer by name.

  - **Template**: `jamfpro://computers/inventory/name/{name}`
  - **Parameters**:
    - `name`: Computer name (string, required)

- **Get FileVault Inventory**
  Retrieves FileVault inventory information for all computers.

  - **Template**: `jamfpro://computers/inventory/filevault`

- **Get FileVault Inventory by ID**
  Retrieves FileVault inventory information for a specific computer by ID.

  - **Template**: `jamfpro://computers/inventory/filevault/{id}`
  - **Parameters**:
    - `id`: Computer ID (number, required)

## References

- [Jamf Pro API Overview](https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-overview)
- [go-api-sdk-jamfpro](https://github.com/deploymenttheory/go-api-sdk-jamfpro)

## License

This project is licensed under the terms of the MIT open source license. Please refer to [LICENSE](./LICENSE) for the full terms.
