package jamfpromcp

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

type jsonrpcRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type jsonrpcResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *rpcError   `json:"error,omitempty"`
}

type rpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type initializeResult struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	ServerInfo      map[string]interface{} `json:"serverInfo"`
	Capabilities    map[string]interface{} `json:"capabilities"`
}

type toolDef struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema json.RawMessage `json:"inputSchema"`
}

type listToolsResult struct {
	Tools []toolDef `json:"tools"`
}

type callToolParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

type callToolResult struct {
	Content []toolContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}

type toolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type getComputerByIDArgs struct {
	ID int `json:"id"`
}

func RunServer(client *jamfpro.Client) {
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	log := func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}

	listComputersSchema := json.RawMessage(`{"type":"object","properties":{},"additionalProperties":false}`)
	getComputerByIDSchema := json.RawMessage(`{"type":"object","properties":{"id":{"type":"integer","description":"Computer ID"}},"required":["id"],"additionalProperties":false}`)

	for {
		var req jsonrpcRequest
		if err := decoder.Decode(&req); err != nil {
			if err == io.EOF {
				break
			}
			log("Decode error: %v", err)
			break
		}

		switch req.Method {
		case "initialize":
			resp := jsonrpcResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result: initializeResult{
					ProtocolVersion: "2024-11-05",
					ServerInfo: map[string]interface{}{
						"name":    "jamfpro-mcp-server",
						"version": "0.1.0",
					},
					Capabilities: map[string]interface{}{
						"tools": map[string]interface{}{},
					},
				},
			}
			encoder.Encode(resp)

		case "tools/list":
			resp := jsonrpcResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result: listToolsResult{
					Tools: []toolDef{
						{
							Name:        "list_computers",
							Description: "List all computers in Jamf Pro",
							InputSchema: listComputersSchema,
						},
						{
							Name:        "get_computer_by_id",
							Description: "Get a computer by its ID",
							InputSchema: getComputerByIDSchema,
						},
					},
				},
			}
			encoder.Encode(resp)

		case "tools/call":
			var params callToolParams
			if err := json.Unmarshal(req.Params, &params); err != nil {
				encoder.Encode(jsonrpcResponse{
					JSONRPC: "2.0",
					ID:      req.ID,
					Error:   &rpcError{Code: -32602, Message: "Invalid params", Data: err.Error()},
				})
				continue
			}
			var result callToolResult
			switch params.Name {
			case "list_computers":
				computers, err := client.GetComputers()
				if err != nil {
					result = callToolResult{
						Content: []toolContent{{Type: "text", Text: "Error: " + err.Error()}},
						IsError: true,
					}
				} else {
					b, _ := json.MarshalIndent(computers, "", "  ")
					result = callToolResult{
						Content: []toolContent{{Type: "text", Text: string(b)}},
					}
				}
			case "get_computer_by_id":
				var args getComputerByIDArgs
				if err := json.Unmarshal(params.Arguments, &args); err != nil {
					result = callToolResult{
						Content: []toolContent{{Type: "text", Text: "Invalid arguments: " + err.Error()}},
						IsError: true,
					}
				} else {
					comp, err := client.GetComputerByID(strconv.Itoa(args.ID))
					if err != nil {
						result = callToolResult{
							Content: []toolContent{{Type: "text", Text: "Error: " + err.Error()}},
							IsError: true,
						}
					} else {
						b, _ := json.MarshalIndent(comp, "", "  ")
						result = callToolResult{
							Content: []toolContent{{Type: "text", Text: string(b)}},
						}
					}
				}
			default:
				result = callToolResult{
					Content: []toolContent{{Type: "text", Text: "Unknown tool: " + params.Name}},
					IsError: true,
				}
			}
			encoder.Encode(jsonrpcResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  result,
			})

		case "resources/list":
			encoder.Encode(jsonrpcResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  map[string]interface{}{"resources": []interface{}{}},
			})
		case "prompts/list":
			encoder.Encode(jsonrpcResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Result:  map[string]interface{}{"prompts": []interface{}{}},
			})
		case "notifications/initialized", "initialized":
			// No response needed for notifications
			continue
		default:
			encoder.Encode(jsonrpcResponse{
				JSONRPC: "2.0",
				ID:      req.ID,
				Error:   &rpcError{Code: -32601, Message: "Method not found"},
			})
		}
	}
}
