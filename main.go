package main

import cmd "deploymenttheory/jamfpro-mcp-server/cmd/jamfpro-mcp-server"

// Entrypoint for jamfpro-mcp-server. All logic is in cmd/jamfpro-mcp-server.

func main() {
	cmd.Execute()
}
