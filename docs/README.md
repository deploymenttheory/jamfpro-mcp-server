# Jamf Pro MCP Server - Project Structure

This document provides a complete overview of the Jamf Pro MCP Server project structure and architecture.

## Project Overview

The Jamf Pro MCP Server is a comprehensive Model Context Protocol server that provides seamless integration with Jamf Pro APIs. It's built in Go and follows modern software development practices with Docker support, comprehensive configuration management, and modular toolset architecture.

## Complete File Structure

```
jamfpro-mcp-server/
├── cmd/
│   └── jamfpro-mcp-server/
│       └── main.go                 # Main application entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── mcp/
│   │   ├── protocol.go             # MCP protocol implementation
│   │   └── translations.go         # Tool description translations
│   ├── server/
│   │   └── server.go               # Main MCP server implementation
│   └── toolsets/
│       ├── toolsets.go             # Toolset interface and factory
│       ├── computers.go            # Computers toolset
│       ├── mobile_devices.go       # Mobile devices toolset
│       └── placeholder_toolsets.go # Other toolset implementations
├── examples/
│   └── config.example.json         # Example configuration file
├── go.mod                          # Go module definition
├── go.sum                          # Go module checksums
├── Dockerfile                      # Production Docker image
├── Dockerfile.dev                  # Development Docker image
├── docker-compose.yml              # Docker Compose configuration
├── Makefile                        # Build and development commands
├── README.md                       # Project documentation
├── .gitignore                      # Git ignore rules
├── .air.toml.example              # Hot reload configuration
└── PROJECT_STRUCTURE.md           # This file
```

## Architecture Overview

### 1. **Command Layer** (`cmd/`)
- **main.go**: Application entry point with CLI interface using Cobra
- Handles command-line arguments, environment variables, and configuration loading
- Manages graceful shutdown and signal handling

### 2. **Configuration Layer** (`internal/config/`)
- **config.go**: Comprehensive configuration management
- Supports multiple configuration sources: CLI flags, environment variables, JSON files
- Validates configuration and provides defaults
- Maps to Jamf Pro SDK configuration format

### 3. **MCP Protocol Layer** (`internal/mcp/`)
- **protocol.go**: Full MCP 2024-11-05 protocol implementation
- **translations.go**: Tool description management and customization
- Handles message parsing, routing, and response formatting
- Supports tool registration and execution

### 4. **Server Layer** (`internal/server/`)
- **server.go**: Main server orchestration
- Manages Jamf Pro client initialization
- Handles stdio communication for MCP protocol
- Coordinates toolset registration and execution

### 5. **Toolsets Layer** (`internal/toolsets/`)
- **toolsets.go**: Base toolset interface and factory pattern
- **computers.go**: Full computers management implementation
- **mobile_devices.go**: Mobile device management tools
- **placeholder_toolsets.go**: Framework for additional toolsets

## Key Features Implemented

### ✅ **Core MCP Protocol**
- Full MCP 2024-11-05 specification compliance
- Tool registration and execution
- Message handling and error management
- Client capability negotiation

### ✅ **Jamf Pro Integration**
- OAuth2 and Basic authentication support
- Comprehensive API coverage through go-api-sdk-jamfpro
- Rate limiting and retry mechanisms
- Connection health checking

### ✅ **Toolset Architecture**
- Modular toolset design for easy extension
- Factory pattern for toolset creation
- Base toolset with common functionality
- Tool description customization

### ✅ **Configuration Management**
- Multiple configuration sources
- Environment variable support
- JSON configuration files
- Command-line argument parsing
- Configuration validation

### ✅ **Development & Deployment**
- Docker support with multi-stage builds
- Docker Compose for development
- Hot reloading with Air
- Comprehensive Makefile
- Development and production environments

### ✅ **Operational Features**
- Structured logging with configurable levels
- Health checks and monitoring
- Graceful shutdown handling
- Sensitive data protection
- Error handling and recovery

## Implemented Toolsets

### **Computers Toolset** (Fully Implemented)
- `get_computers` - List all computers
- `get_computer_by_id` - Get computer details by ID
- `get_computer_by_name` - Get computer details by name
- `get_computers_inventory` - Get inventory with filtering
- `get_computer_inventory_by_id` - Get specific computer inventory
- `update_computer_inventory` - Update computer information
- `get_computer_groups` - List computer groups
- `get_computer_group_by_id` - Get computer group details
- `delete_computer` - Delete computer by ID

### **Mobile Devices Toolset** (Fully Implemented)
- `get_mobile_devices` - List all mobile devices
- `get_mobile_device_by_id` - Get device details by ID
- `get_mobile_device_by_name` - Get device details by name
- `get_mobile_device_groups` - List mobile device groups
- `get_mobile_device_group_by_id` - Get group details
- `get_mobile_device_applications` - List mobile apps
- `get_mobile_device_configuration_profiles` - List profiles
- `delete_mobile_device` - Delete mobile device

### **Additional Toolsets** (Framework Ready)
- Policies, Users, Groups, Configuration Profiles
- Scripts, Buildings, Departments, Categories
- Sites, API Roles, API Integrations
- Inventory, Packages, Printers, Network Segments
- Webhooks, VPP, Advanced Searches
- Extension Attributes, LDAP, Self Service
- Patch Management, Applications, Restrictions
- Disk Encryption, Enrollment, Server Information

## Usage Examples

### **Basic Deployment**
```bash
# Using environment variables
export JAMF_INSTANCE_URL="https://company.jamfcloud.com"
export JAMF_CLIENT_ID="your-client-id"
export JAMF_CLIENT_SECRET="your-client-secret"
./jamfpro-mcp-server stdio
```

### **Docker Deployment**
```bash
# Build and run with Docker
make docker-run
```

### **Development**
```bash
# Set up development environment
make dev-setup

# Run with hot reloading
make watch

# Run tests
make test
```

### **MCP Client Integration**
```json
{
  "mcp": {
    "servers": {
      "jamfpro": {
        "command": "./jamfpro-mcp-server",
        "args": ["stdio"],
        "env": {
          "JAMF_INSTANCE_URL": "https://company.jamfcloud.com",
          "JAMF_CLIENT_ID": "your-client-id",
          "JAMF_CLIENT_SECRET": "your-client-secret"
        }
      }
    }
  }
}
```

## Next Steps for Enhancement

### **High Priority**
1. Complete implementation of remaining toolsets
2. Add comprehensive test coverage
3. Implement caching for frequently accessed data
4. Add metrics and monitoring endpoints

### **Medium Priority**
1. Support for webhook endpoints
2. Batch operation support
3. Advanced filtering and search capabilities
4. Configuration validation UI

### **Future Considerations**
1. Plugin architecture for custom toolsets
2. GraphQL API support
3. Real-time event streaming
4. Multi-tenant support

## Building and Running

### **Prerequisites**
- Go 1.21+
- Docker (optional)
- Jamf Pro instance with API access

### **Quick Start**
```bash
# Clone and build
git clone <repository>
cd jamfpro-mcp-server
make build

# Configure
cp examples/config.example.json jamfpro-mcp-server-config.json
# Edit configuration file with your Jamf Pro details

# Run
./jamfpro-mcp-server stdio
```

### **Development**
```bash
# Set up development environment
make dev-setup

# Run with hot reloading
cp .air.toml.example .air.toml
make watch
```

This project provides a solid foundation for Jamf Pro integration via the Model Context Protocol, with room for extensive customization and enhancement based on specific needs.