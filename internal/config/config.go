package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents the server configuration
type Config struct {
	// Server configuration
	Host string `json:"host"`
	Port int    `json:"port"`

	// Jamf Pro configuration
	JamfProURL   string `json:"jamfpro_url"`
	AuthMethod   string `json:"auth_method"` // "oauth2" or "basic"
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`

	// Feature flags
	ReadOnly        bool     `json:"read_only"`
	EnabledToolsets []string `json:"enabled_toolsets"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Host:            "localhost",
		Port:            8080,
		AuthMethod:      "oauth2",
		ReadOnly:        false,
		EnabledToolsets: []string{"all"},
	}
}

// Load loads the configuration from a file
func Load(path string) (*Config, error) {
	// Start with default configuration
	cfg := DefaultConfig()

	// If no config file is specified, check for environment variables
	if path == "" {
		return loadFromEnv(cfg)
	}

	// Open and read the config file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Decode the config file
	if err := json.NewDecoder(file).Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	// Override with environment variables if they exist
	return loadFromEnv(cfg)
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(cfg *Config) (*Config, error) {
	// Server configuration
	if host := os.Getenv("JAMFPRO_MCP_HOST"); host != "" {
		cfg.Host = host
	}

	// Jamf Pro configuration
	if url := os.Getenv("JAMFPRO_URL"); url != "" {
		cfg.JamfProURL = url
	}

	if authMethod := os.Getenv("JAMFPRO_AUTH_METHOD"); authMethod != "" {
		cfg.AuthMethod = authMethod
	}

	if clientID := os.Getenv("JAMFPRO_CLIENT_ID"); clientID != "" {
		cfg.ClientID = clientID
	}

	if clientSecret := os.Getenv("JAMFPRO_CLIENT_SECRET"); clientSecret != "" {
		cfg.ClientSecret = clientSecret
	}

	if username := os.Getenv("JAMFPRO_USERNAME"); username != "" {
		cfg.Username = username
	}

	if password := os.Getenv("JAMFPRO_PASSWORD"); password != "" {
		cfg.Password = password
	}

	// Feature flags
	if readOnly := os.Getenv("JAMFPRO_READ_ONLY"); readOnly == "true" {
		cfg.ReadOnly = true
	}

	// Validate the configuration
	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// validate validates the configuration
func validate(cfg *Config) error {
	if cfg.JamfProURL == "" {
		return fmt.Errorf("Jamf Pro URL is required")
	}

	if cfg.AuthMethod == "oauth2" {
		if cfg.ClientID == "" || cfg.ClientSecret == "" {
			return fmt.Errorf("client ID and client secret are required for OAuth2 authentication")
		}
	} else if cfg.AuthMethod == "basic" {
		if cfg.Username == "" || cfg.Password == "" {
			return fmt.Errorf("username and password are required for basic authentication")
		}
	} else {
		return fmt.Errorf("invalid authentication method: %s", cfg.AuthMethod)
	}

	return nil
}
