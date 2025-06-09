package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/deploymenttheory/jamfpro-mcp-server/internal/config"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/mcp"
	"github.com/deploymenttheory/jamfpro-mcp-server/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "jamfpro-mcp-server",
		Short:   "Jamf Pro MCP (Model Context Protocol) Server",
		Long:    `A Model Context Protocol server that provides seamless integration with Jamf Pro APIs, enabling advanced automation and interaction capabilities for AI tools and applications.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
		Run:     runServer,
	}

	// Global flags
	rootCmd.PersistentFlags().String("config", "", "config file path")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error)")
	rootCmd.PersistentFlags().StringSlice("toolsets", []string{"all"}, "comma-separated list of toolsets to enable (computers,mobile-devices,policies,users,groups,configuration-profiles,scripts,buildings,all)")
	rootCmd.PersistentFlags().Bool("dynamic-toolsets", false, "enable dynamic toolset discovery")
	rootCmd.PersistentFlags().Bool("export-translations", false, "export tool descriptions to config file")

	// Environment variable support
	rootCmd.PersistentFlags().String("jamf-instance-url", "", "Jamf Pro instance URL (can also use JAMF_INSTANCE_URL)")
	rootCmd.PersistentFlags().String("jamf-client-id", "", "Jamf Pro client ID (can also use JAMF_CLIENT_ID)")
	rootCmd.PersistentFlags().String("jamf-client-secret", "", "Jamf Pro client secret (can also use JAMF_CLIENT_SECRET)")
	rootCmd.PersistentFlags().String("jamf-username", "", "Jamf Pro username for basic auth (can also use JAMF_USERNAME)")
	rootCmd.PersistentFlags().String("jamf-password", "", "Jamf Pro password for basic auth (can also use JAMF_PASSWORD)")
	rootCmd.PersistentFlags().String("auth-method", "oauth2", "authentication method: oauth2 or basic (can also use JAMF_AUTH_METHOD)")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runServer(cmd *cobra.Command, args []string) {
	// Load configuration
	cfg, err := config.Load(cmd)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger, err := initLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Handle export translations
	if cfg.ExportTranslations {
		if err := exportTranslations(logger); err != nil {
			logger.Fatal("Failed to export translations", zap.Error(err))
		}
		return
	}

	// Validate toolsets
	if err := validateToolsets(cfg.Toolsets); err != nil {
		logger.Fatal("Invalid toolsets", zap.Error(err))
	}

	logger.Info("Starting Jamf Pro MCP Server",
		zap.String("version", version),
		zap.Strings("toolsets", cfg.Toolsets),
		zap.Bool("dynamic-toolsets", cfg.DynamicToolsets),
		zap.String("auth-method", cfg.AuthMethod),
	)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Initialize and start the MCP server
	mcpServer, err := server.New(cfg, logger)
	if err != nil {
		logger.Fatal("Failed to create MCP server", zap.Error(err))
	}

	// Start server in a goroutine
	errChan := make(chan error, 1)
	go func() {
		errChan <- mcpServer.Start(ctx)
	}()

	// Wait for shutdown signal or error
	select {
	case sig := <-sigChan:
		logger.Info("Received shutdown signal", zap.String("signal", sig.String()))
		cancel()
	case err := <-errChan:
		if err != nil {
			logger.Error("Server error", zap.Error(err))
		}
	}

	logger.Info("Shutting down Jamf Pro MCP Server")
}

func initLogger(level string) (*zap.Logger, error) {
	var zapLevel zap.AtomicLevel
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		zapLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		zapLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		zapLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		zapLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	config := zap.NewProductionConfig()
	config.Level = zapLevel
	config.Development = false
	config.DisableStacktrace = true

	return config.Build()
}

func validateToolsets(toolsets []string) error {
	validToolsets := map[string]bool{
		"all":                     true,
		"computers":               true,
		"mobile-devices":          true,
		"policies":                true,
		"users":                   true,
		"groups":                  true,
		"configuration-profiles":  true,
		"scripts":                 true,
		"buildings":               true,
		"departments":             true,
		"categories":              true,
		"sites":                   true,
		"api-roles":               true,
		"api-integrations":        true,
		"inventory":               true,
		"packages":                true,
		"printers":                true,
		"network-segments":        true,
		"webhooks":                true,
		"vpp":                     true,
		"advanced-searches":       true,
		"extension-attributes":    true,
		"ldap":                    true,
		"self-service":            true,
		"patch-management":        true,
		"mobile-applications":     true,
		"mac-applications":        true,
		"restrictions":            true,
		"disk-encryption":         true,
		"enrollment":              true,
		"jamf-pro-information":    true,
		"computer-checkin":        true,
		"smtp":                    true,
		"gsx-connection":          true,
		"sso":                     true,
		"volume-purchasing":       true,
		"licensed-software":       true,
		"allowed-file-extensions": true,
		"removable-mac-addresses": true,
		"restricted-software":     true,
		"software-update-servers": true,
		"dock-items":              true,
		"directory-bindings":      true,
		"distribution-points":     true,
		"ebooks":                  true,
		"ibeacons":                true,
		"byo-profiles":            true,
	}

	for _, toolset := range toolsets {
		if !validToolsets[toolset] {
			return fmt.Errorf("invalid toolset: %s", toolset)
		}
	}
	return nil
}

func exportTranslations(logger *zap.Logger) error {
	translations := mcp.GetDefaultTranslations()

	configFile := "jamfpro-mcp-server-config.json"

	// Check if file exists and load existing translations
	if data, err := os.ReadFile(configFile); err == nil {
		var existing map[string]string
		if err := json.Unmarshal(data, &existing); err == nil {
			// Merge existing with defaults, keeping existing values
			for key, value := range existing {
				translations[key] = value
			}
		}
	}

	data, err := json.MarshalIndent(translations, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal translations: %w", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	logger.Info("Exported translations", zap.String("file", configFile))
	return nil
}
