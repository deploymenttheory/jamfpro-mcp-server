package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/config"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/toolsets"
	"go.uber.org/zap"
)

// Server represents the main MCP server
type Server struct {
	config     *config.Config
	logger     *zap.Logger
	mcpServer  *mcp.Server
	jamfClient *jamfpro.Client
	toolsets   map[string]toolsets.Toolset
}

// New creates a new server instance
func New(cfg *config.Config, logger *zap.Logger) (*Server, error) {
	// Create MCP server
	mcpServer := mcp.NewServer("jamfpro-mcp-server", "1.0.0")

	// Initialize Jamf Pro client
	jamfClient, err := initializeJamfClient(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Jamf Pro client: %w", err)
	}

	server := &Server{
		config:     cfg,
		logger:     logger,
		mcpServer:  mcpServer,
		jamfClient: jamfClient,
		toolsets:   make(map[string]toolsets.Toolset),
	}

	// Initialize toolsets
	if err := server.initializeToolsets(); err != nil {
		return nil, fmt.Errorf("failed to initialize toolsets: %w", err)
	}

	// Initialize resource provider
	if err := server.initializeResourceProvider(); err != nil {
		logger.Warn("Failed to initialize resource provider", zap.Error(err))
		// Continue anyway, resources are optional
	}

	return server, nil
}

// Start starts the MCP server
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting MCP server")

	// Create a scanner for reading from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Process messages from stdin
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line := scanner.Text()
			if line == "" {
				continue
			}

			s.logger.Debug("Received message", zap.String("message", line))

			if err := s.processMessage(ctx, line); err != nil {
				s.logger.Error("Failed to process message", zap.Error(err))

				// Send error response
				errorMsg := &mcp.Message{
					JSONRPC: "2.0",
					Error: &mcp.Error{
						Code:    mcp.InternalError,
						Message: err.Error(),
					},
				}

				if err := s.sendMessage(errorMsg); err != nil {
					s.logger.Error("Failed to send error message", zap.Error(err))
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	return nil
}

// processMessage processes an incoming message
func (s *Server) processMessage(ctx context.Context, messageText string) error {
	var msg mcp.Message
	if err := json.Unmarshal([]byte(messageText), &msg); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	// Handle the message
	response, err := s.mcpServer.HandleMessage(ctx, &msg)
	if err != nil {
		return fmt.Errorf("failed to handle message: %w", err)
	}

	// Send the response
	return s.sendMessage(response)
}

// sendMessage sends a message to stdout
func (s *Server) sendMessage(msg *mcp.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Write to stdout with newline
	if _, err := fmt.Fprintln(os.Stdout, string(data)); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	s.logger.Debug("Sent message", zap.String("message", string(data)))
	return nil
}

// initializeJamfClient initializes the Jamf Pro client
func initializeJamfClient(cfg *config.Config, logger *zap.Logger) (*jamfpro.Client, error) {
	logger.Info("Initializing Jamf Pro client",
		zap.String("instance_url", cfg.JamfInstanceURL),
		zap.String("auth_method", cfg.AuthMethod),
	)

	// Set environment variables for the Jamf Pro SDK
	if err := setJamfEnvironmentVariables(cfg); err != nil {
		return nil, fmt.Errorf("failed to set environment variables: %w", err)
	}

	// Build client using environment variables
	client, err := jamfpro.BuildClientWithEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to build Jamf Pro client: %w", err)
	}

	// Test the connection
	logger.Info("Testing Jamf Pro connection")
	_, err = client.GetJamfProInformation()
	if err != nil {
		logger.Warn("Failed to test Jamf Pro connection, but continuing", zap.Error(err))
		// Don't fail here as the connection might work for other operations
	} else {
		logger.Info("Successfully connected to Jamf Pro")
	}

	return client, nil
}

// setJamfEnvironmentVariables sets environment variables for the Jamf Pro SDK - FIXED with validation
func setJamfEnvironmentVariables(cfg *config.Config) error {
	// Validate URL format first
	if !isValidURL(cfg.JamfInstanceURL) {
		return fmt.Errorf("invalid Jamf instance URL format: %s", cfg.JamfInstanceURL)
	}

	envVars := map[string]string{
		"INSTANCE_DOMAIN":                     cfg.JamfInstanceURL,
		"AUTH_METHOD":                         cfg.AuthMethod,
		"LOG_LEVEL":                           cfg.LogLevel,
		"LOG_OUTPUT_FORMAT":                   "json",
		"HIDE_SENSITIVE_DATA":                 fmt.Sprintf("%t", cfg.HideSensitiveData),
		"MAX_RETRY_ATTEMPTS":                  fmt.Sprintf("%d", cfg.MaxRetryAttempts),
		"ENABLE_DYNAMIC_RATE_LIMITING":        fmt.Sprintf("%t", cfg.EnableDynamicRateLimiting),
		"MAX_CONCURRENT_REQUESTS":             fmt.Sprintf("%d", cfg.MaxConcurrentRequests),
		"TOKEN_REFRESH_BUFFER_PERIOD_SECONDS": fmt.Sprintf("%d", cfg.TokenRefreshBufferSeconds),
		"TOTAL_RETRY_DURATION_SECONDS":        fmt.Sprintf("%d", cfg.TotalRetryDurationSeconds),
		"CUSTOM_TIMEOUT_SECONDS":              fmt.Sprintf("%d", cfg.CustomTimeoutSeconds),
		"FOLLOW_REDIRECTS":                    fmt.Sprintf("%t", cfg.FollowRedirects),
		"MAX_REDIRECTS":                       fmt.Sprintf("%d", cfg.MaxRedirects),
		"ENABLE_CONCURRENCY_MANAGEMENT":       fmt.Sprintf("%t", cfg.EnableConcurrencyManagement),
		"JAMF_LOAD_BALANCER_LOCK":             fmt.Sprintf("%t", cfg.JamfLoadBalancerLock),
	}

	// Add auth-specific variables with validation
	if cfg.AuthMethod == "oauth2" {
		if cfg.JamfClientID == "" {
			return fmt.Errorf("client ID is required for OAuth2 authentication")
		}
		if cfg.JamfClientSecret == "" {
			return fmt.Errorf("client secret is required for OAuth2 authentication")
		}
		envVars["CLIENT_ID"] = cfg.JamfClientID
		envVars["CLIENT_SECRET"] = cfg.JamfClientSecret
	} else {
		if cfg.JamfUsername == "" {
			return fmt.Errorf("username is required for basic authentication")
		}
		if cfg.JamfPassword == "" {
			return fmt.Errorf("password is required for basic authentication")
		}
		envVars["BASIC_AUTH_USERNAME"] = cfg.JamfUsername
		envVars["BASIC_AUTH_PASSWORD"] = cfg.JamfPassword
	}

	// Set all environment variables
	for key, value := range envVars {
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("failed to set environment variable %s: %w", key, err)
		}
	}

	return nil
}

// isValidURL validates URL format - ADDED helper function
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// initializeToolsets initializes the available toolsets
func (s *Server) initializeToolsets() error {
	s.logger.Info("Initializing toolsets", zap.Strings("enabled_toolsets", s.config.Toolsets))

	// Create toolset factory
	factory := toolsets.NewFactory(s.jamfClient, s.logger)

	// Determine which toolsets to enable
	enabledToolsets := s.getEnabledToolsets()

	// Initialize each enabled toolset
	for _, toolsetName := range enabledToolsets {
		s.logger.Debug("Initializing toolset", zap.String("toolset", toolsetName))

		toolset, err := factory.CreateToolset(toolsetName)
		if err != nil {
			s.logger.Warn("Failed to create toolset",
				zap.String("toolset", toolsetName),
				zap.Error(err))
			continue
		}

		s.toolsets[toolsetName] = toolset

		tools := toolset.GetTools()
		for _, tool := range tools {
			s.mcpServer.RegisterToolDefinition(&tool)

			s.mcpServer.RegisterTool(tool.Name, s.createToolHandler(toolset, tool.Name))
			s.logger.Debug("Registered tool",
				zap.String("toolset", toolsetName),
				zap.String("tool", tool.Name))
		}
	}

	s.logger.Info("Toolset initialization complete",
		zap.Int("toolsets_initialized", len(s.toolsets)),
		zap.Int("total_tools_registered", len(s.mcpServer.GetRegisteredTools())))

	return nil
}

// getEnabledToolsets returns the list of toolsets that should be enabled
func (s *Server) getEnabledToolsets() []string {
	// If "all" is specified, return all available toolsets
	for _, toolset := range s.config.Toolsets {
		if toolset == "all" {
			return []string{
				"computers", "computer-inventory", "mobile-devices", "policies",
				"scripts", "users", "groups", "configuration-profiles", "buildings",
				"departments", "categories", "sites", "api-roles", "api-integrations",
				"inventory", "packages", "printers", "network-segments",
				"webhooks", "vpp", "advanced-searches", "extension-attributes",
				"ldap", "self-service", "patch-management", "mobile-applications",
				"mac-applications", "restrictions", "disk-encryption",
				"enrollment", "jamf-pro-information",
			}
		}
	}

	return s.config.Toolsets
}

// createToolHandler creates a tool handler for a specific tool
func (s *Server) createToolHandler(toolset toolsets.Toolset, toolName string) mcp.ToolHandler {
	return func(ctx context.Context, params mcp.CallToolParams) (*mcp.CallToolResult, error) {
		s.logger.Debug("Executing tool",
			zap.String("tool", toolName),
			zap.Any("arguments", params.Arguments))

		result, err := toolset.ExecuteTool(ctx, toolName, params.Arguments)
		if err != nil {
			s.logger.Error("Tool execution failed",
				zap.String("tool", toolName),
				zap.Error(err))

			return &mcp.CallToolResult{
				Content: []mcp.ToolContent{
					{
						Type: "text",
						Text: fmt.Sprintf("Error executing tool %s: %v", toolName, err),
					},
				},
				IsError: true,
			}, nil
		}

		s.logger.Debug("Tool execution successful", zap.String("tool", toolName))

		return &mcp.CallToolResult{
			Content: []mcp.ToolContent{
				{
					Type: "text",
					Text: result,
				},
			},
			IsError: false,
		}, nil
	}
}

// initializeResourceProvider initializes the resource provider
func (s *Server) initializeResourceProvider() error {
	s.logger.Info("Initializing resource provider")

	// Create a file resource provider with the current working directory as base path
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	resourceProvider := mcp.NewFileResourceProvider(wd)

	// Register common resource directories
	if err := s.registerCommonResourceDirectories(resourceProvider); err != nil {
		return fmt.Errorf("failed to register resource directories: %w", err)
	}

	// Set the resource provider on the MCP server
	s.mcpServer.SetResourceProvider(resourceProvider)

	s.logger.Info("Resource provider initialized")
	return nil
}

// registerCommonResourceDirectories registers commonly used directories as resources
func (s *Server) registerCommonResourceDirectories(provider *mcp.FileResourceProvider) error {
	// Register templates directory if it exists
	if dirExists("templates") {
		if err := provider.RegisterDirectory("file://templates", "templates"); err != nil {
			s.logger.Warn("Failed to register templates directory", zap.Error(err))
		} else {
			s.logger.Info("Registered templates directory as resource")
		}
	}

	// Register examples directory if it exists
	if dirExists("examples") {
		if err := provider.RegisterDirectory("file://examples", "examples"); err != nil {
			s.logger.Warn("Failed to register examples directory", zap.Error(err))
		} else {
			s.logger.Info("Registered examples directory as resource")
		}
	}

	// Register scripts directory if it exists
	if dirExists("scripts") {
		if err := provider.RegisterDirectory("file://scripts", "scripts"); err != nil {
			s.logger.Warn("Failed to register scripts directory", zap.Error(err))
		} else {
			s.logger.Info("Registered scripts directory as resource")
		}
	}

	// Register workflows directory if it exists
	if dirExists("workflows") {
		if err := provider.RegisterDirectory("file://workflows", "workflows"); err != nil {
			s.logger.Warn("Failed to register workflows directory", zap.Error(err))
		} else {
			s.logger.Info("Registered workflows directory as resource")
		}
	}

	// Register docs directory if it exists
	if dirExists("docs") {
		if err := provider.RegisterDirectory("file://docs", "docs"); err != nil {
			s.logger.Warn("Failed to register docs directory", zap.Error(err))
		} else {
			s.logger.Info("Registered docs directory as resource")
		}
	}

	return nil
}

// dirExists checks if a directory exists at the given path
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
