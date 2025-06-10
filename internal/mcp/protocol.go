package mcp

import (
	"context"
	"encoding/json"
	"fmt"
)

// MCP Protocol types based on the Model Context Protocol specification

// Message represents a base MCP message
type Message struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

// Error represents an MCP error
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Standard MCP error codes
const (
	ParseError           = -32700
	InvalidRequest       = -32600
	MethodNotFound       = -32601
	InvalidParams        = -32602
	InternalError        = -32603
	ServerError          = -32000
	ResourceNotFound     = -32001
	InvalidResourceURI   = -32002
	ResourceAccessDenied = -32003
	ToolNotFound         = -32004
	ToolExecutionError   = -32005
)

// InitializeParams represents the initialize request parameters
type InitializeParams struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    ClientCapabilities     `json:"capabilities"`
	ClientInfo      ClientInfo             `json:"clientInfo"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
}

// ClientCapabilities represents client capabilities
type ClientCapabilities struct {
	Experimental map[string]interface{} `json:"experimental,omitempty"`
	Sampling     *SamplingCapabilities  `json:"sampling,omitempty"`
}

// SamplingCapabilities represents sampling capabilities
type SamplingCapabilities struct{}

// ClientInfo represents client information
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// InitializeResult represents the initialize response
type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo         `json:"serverInfo"`
	Instructions    string             `json:"instructions,omitempty"`
}

// ServerCapabilities represents server capabilities
type ServerCapabilities struct {
	Experimental map[string]interface{} `json:"experimental,omitempty"`
	Logging      *LoggingCapabilities   `json:"logging,omitempty"`
	Prompts      *PromptsCapabilities   `json:"prompts,omitempty"`
	Resources    *ResourcesCapabilities `json:"resources,omitempty"`
	Tools        *ToolsCapabilities     `json:"tools,omitempty"`
}

// LoggingCapabilities represents logging capabilities
type LoggingCapabilities struct{}

// PromptsCapabilities represents prompts capabilities
type PromptsCapabilities struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// ResourcesCapabilities represents resources capabilities
type ResourcesCapabilities struct {
	Subscribe   bool `json:"subscribe,omitempty"`
	ListChanged bool `json:"listChanged,omitempty"`
}

// ToolsCapabilities represents tools capabilities
type ToolsCapabilities struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// ServerInfo represents server information
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ListToolsParams represents the list tools request parameters
type ListToolsParams struct {
	Cursor string `json:"cursor,omitempty"`
}

// ListToolsResult represents the list tools response
type ListToolsResult struct {
	Tools      []Tool  `json:"tools"`
	NextCursor *string `json:"nextCursor,omitempty"`
}

// Tool represents an MCP tool
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema ToolInputSchema        `json:"inputSchema"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

// ToolInputSchema represents the tool input schema
type ToolInputSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
}

// CallToolParams represents the call tool request parameters
type CallToolParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

// CallToolResult represents the call tool response
type CallToolResult struct {
	Content []ToolContent          `json:"content"`
	IsError bool                   `json:"isError,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

// ToolContent represents tool content
type ToolContent struct {
	Type string      `json:"type"`
	Text string      `json:"text,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Name string      `json:"name,omitempty"`
}

// ListResourcesParams represents the list resources request parameters
type ListResourcesParams struct {
	Cursor string `json:"cursor,omitempty"`
}

// ListResourcesResult represents the list resources response
type ListResourcesResult struct {
	Resources  []Resource `json:"resources"`
	NextCursor *string    `json:"nextCursor,omitempty"`
}

// Resource represents an MCP resource
type Resource struct {
	URI         string                 `json:"uri"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	MimeType    string                 `json:"mimeType,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

// ReadResourceParams represents the read resource request parameters
type ReadResourceParams struct {
	URI string `json:"uri"`
}

// ReadResourceResult represents the read resource response
type ReadResourceResult struct {
	Contents []ResourceContent      `json:"contents"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

// ResourceContent represents resource content
type ResourceContent struct {
	URI      string      `json:"uri"`
	MimeType string      `json:"mimeType,omitempty"`
	Text     string      `json:"text,omitempty"`
	Blob     string      `json:"blob,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
}

// LoggingLevel represents logging levels
type LoggingLevel string

const (
	LoggingLevelDebug     LoggingLevel = "debug"
	LoggingLevelInfo      LoggingLevel = "info"
	LoggingLevelNotice    LoggingLevel = "notice"
	LoggingLevelWarning   LoggingLevel = "warning"
	LoggingLevelError     LoggingLevel = "error"
	LoggingLevelCritical  LoggingLevel = "critical"
	LoggingLevelAlert     LoggingLevel = "alert"
	LoggingLevelEmergency LoggingLevel = "emergency"
)

// LoggingParams represents the logging message parameters
type LoggingParams struct {
	Level  LoggingLevel           `json:"level"`
	Data   interface{}            `json:"data,omitempty"`
	Logger string                 `json:"logger,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}

// Handler represents an MCP handler interface
type Handler interface {
	HandleMessage(ctx context.Context, msg *Message) (*Message, error)
}

// Server represents the MCP server - FIXED with missing fields and methods
type Server struct {
	capabilities     ServerCapabilities
	serverInfo       ServerInfo
	toolHandlers     map[string]ToolHandler
	toolRegistry     map[string]*Tool // ADDED: Store actual tool definitions
	resourceProvider ResourceProvider
	initialized      bool
}

// ToolHandler represents a tool handler function
type ToolHandler func(ctx context.Context, params CallToolParams) (*CallToolResult, error)

// NewServer creates a new MCP server - FIXED to initialize all fields
func NewServer(name, version string) *Server {
	return &Server{
		capabilities: ServerCapabilities{
			Tools: &ToolsCapabilities{
				ListChanged: true,
			},
			Resources: &ResourcesCapabilities{
				Subscribe:   false,
				ListChanged: true,
			},
			Logging: &LoggingCapabilities{},
		},
		serverInfo: ServerInfo{
			Name:    name,
			Version: version,
		},
		toolHandlers: make(map[string]ToolHandler),
		toolRegistry: make(map[string]*Tool), // ADDED: Initialize tool registry
		initialized:  false,
	}
}

// SetResourceProvider sets the resource provider for the server - ADDED missing method
func (s *Server) SetResourceProvider(provider ResourceProvider) {
	s.resourceProvider = provider
}

// GetRegisteredTools returns the list of registered tool names - ADDED missing method
func (s *Server) GetRegisteredTools() []string {
	tools := make([]string, 0, len(s.toolHandlers))
	for name := range s.toolHandlers {
		tools = append(tools, name)
	}
	return tools
}

// RegisterTool registers a tool handler
func (s *Server) RegisterTool(name string, handler ToolHandler) {
	s.toolHandlers[name] = handler
}

// RegisterToolDefinition registers a tool definition - ADDED to store schemas
func (s *Server) RegisterToolDefinition(tool *Tool) {
	if s.toolRegistry == nil {
		s.toolRegistry = make(map[string]*Tool)
	}
	toolCopy := *tool // Make a copy to avoid pointer issues
	s.toolRegistry[tool.Name] = &toolCopy
}

// HandleMessage handles an incoming MCP message
func (s *Server) HandleMessage(ctx context.Context, msg *Message) (*Message, error) {
	response := &Message{
		JSONRPC: "2.0",
		ID:      msg.ID,
	}

	switch msg.Method {
	case "initialize":
		result, err := s.handleInitialize(ctx, msg.Params)
		if err != nil {
			response.Error = &Error{
				Code:    InternalError,
				Message: err.Error(),
			}
		} else {
			response.Result = result
		}

	case "tools/list":
		result, err := s.handleListTools(ctx, msg.Params)
		if err != nil {
			response.Error = &Error{
				Code:    InternalError,
				Message: err.Error(),
			}
		} else {
			response.Result = result
		}

	case "tools/call":
		result, err := s.handleCallTool(ctx, msg.Params)
		if err != nil {
			response.Error = &Error{
				Code:    ToolExecutionError,
				Message: err.Error(),
			}
		} else {
			response.Result = result
		}

	case "resources/list":
		result, err := s.handleListResources(ctx, msg.Params)
		if err != nil {
			response.Error = &Error{
				Code:    InternalError,
				Message: err.Error(),
			}
		} else {
			response.Result = result
		}

	case "resources/read":
		result, err := s.handleReadResource(ctx, msg.Params)
		if err != nil {
			response.Error = &Error{
				Code:    ResourceNotFound,
				Message: err.Error(),
			}
		} else {
			response.Result = result
		}

	default:
		response.Error = &Error{
			Code:    MethodNotFound,
			Message: fmt.Sprintf("method not found: %s", msg.Method),
		}
	}

	return response, nil
}

func (s *Server) handleInitialize(ctx context.Context, params interface{}) (*InitializeResult, error) {
	s.initialized = true

	return &InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities:    s.capabilities,
		ServerInfo:      s.serverInfo,
		Instructions: "This server provides access to Jamf Pro APIs for managing Apple devices, mobile devices, policies, scripts, configuration profiles, and more. " +
			"Use the available tools to interact with your Jamf Pro environment. Authentication is handled automatically based on the server configuration.",
	}, nil
}

func (s *Server) handleListTools(ctx context.Context, params interface{}) (*ListToolsResult, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	tools := make([]Tool, 0, len(s.toolHandlers))
	for name := range s.toolHandlers {
		tool := s.getToolDefinition(name)
		if tool != nil {
			tools = append(tools, *tool)
		}
	}

	return &ListToolsResult{
		Tools: tools,
	}, nil
}

func (s *Server) handleCallTool(ctx context.Context, params interface{}) (*CallToolResult, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	var callParams CallToolParams
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}

	if err := json.Unmarshal(data, &callParams); err != nil {
		return nil, fmt.Errorf("failed to unmarshal call tool params: %w", err)
	}

	handler, exists := s.toolHandlers[callParams.Name]
	if !exists {
		return nil, fmt.Errorf("tool not found: %s", callParams.Name)
	}

	return handler(ctx, callParams)
}

// getToolDefinition returns the tool definition for a given tool name - FIXED to use registry
func (s *Server) getToolDefinition(name string) *Tool {
	// Try to get from registry first
	if tool, exists := s.toolRegistry[name]; exists {
		return tool
	}

	// Fallback to basic definition
	return &Tool{
		Name:        name,
		Description: fmt.Sprintf("Tool: %s", name),
		InputSchema: ToolInputSchema{
			Type:       "object",
			Properties: make(map[string]interface{}),
			Required:   []string{},
		},
	}
}

// handleListResources handles the resources/list method
func (s *Server) handleListResources(ctx context.Context, params interface{}) (*ListResourcesResult, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	if s.resourceProvider == nil {
		return &ListResourcesResult{
			Resources: []Resource{},
		}, nil
	}

	resources, err := s.resourceProvider.ListResources()
	if err != nil {
		return nil, fmt.Errorf("failed to list resources: %w", err)
	}

	return &ListResourcesResult{
		Resources: resources,
	}, nil
}

// handleReadResource handles the resources/read method
func (s *Server) handleReadResource(ctx context.Context, params interface{}) (*ReadResourceResult, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	if s.resourceProvider == nil {
		return nil, fmt.Errorf("resource provider not configured")
	}

	var readParams ReadResourceParams
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}

	if err := json.Unmarshal(data, &readParams); err != nil {
		return nil, fmt.Errorf("failed to unmarshal read resource params: %w", err)
	}

	return s.resourceProvider.ReadResource(readParams.URI)
}
