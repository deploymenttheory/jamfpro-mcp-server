package cmd

import (
	"fmt"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	mcp "github.com/deploymenttheory/jamfpro-mcp-server/internal/jamfpromcp"
	"github.com/spf13/cobra"
)

var (
	instanceDomain    string
	authMethod        string
	clientID          string
	clientSecret      string
	username          string
	password          string
	logLevel          string
	hideSensitiveData bool
)

var rootCmd = &cobra.Command{
	Use:   "jamfpro-mcp-server",
	Short: "Jamf Pro MCP Server",
	Long:  `A Model Context Protocol (MCP) server for Jamf Pro, exposing read-only API endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		setEnvIfNotEmpty("INSTANCE_DOMAIN", instanceDomain)
		setEnvIfNotEmpty("AUTH_METHOD", authMethod)
		setEnvIfNotEmpty("CLIENT_ID", clientID)
		setEnvIfNotEmpty("CLIENT_SECRET", clientSecret)
		setEnvIfNotEmpty("BASIC_AUTH_USERNAME", username)
		setEnvIfNotEmpty("BASIC_AUTH_PASSWORD", password)
		setEnvIfNotEmpty("LOG_LEVEL", logLevel)
		os.Setenv("HIDE_SENSITIVE_DATA", fmt.Sprintf("%v", hideSensitiveData))

		client, err := jamfpro.BuildClientWithEnv()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to initialize Jamf Pro client: %v\n", err)
			os.Exit(1)
		}
		mcp.RunServer(client)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&instanceDomain, "instance-domain", "", "Jamf Pro instance domain (required)")
	rootCmd.PersistentFlags().StringVar(&authMethod, "auth-method", "", "Authentication method: oauth2 or basic (required)")
	rootCmd.PersistentFlags().StringVar(&clientID, "client-id", "", "OAuth2 client ID")
	rootCmd.PersistentFlags().StringVar(&clientSecret, "client-secret", "", "OAuth2 client secret")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "Basic auth username")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "Basic auth password")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "warning", "Log level (default: warning)")
	rootCmd.PersistentFlags().BoolVar(&hideSensitiveData, "hide-sensitive-data", true, "Hide sensitive data in logs (default: true)")
}

func setEnvIfNotEmpty(key, value string) {
	if value != "" {
		os.Setenv(key, value)
	}
}
