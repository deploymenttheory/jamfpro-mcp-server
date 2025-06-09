package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config holds the configuration for the Jamf Pro MCP server
type Config struct {
	// Server configuration
	LogLevel           string   `mapstructure:"log_level"`
	Toolsets           []string `mapstructure:"toolsets"`
	DynamicToolsets    bool     `mapstructure:"dynamic_toolsets"`
	ExportTranslations bool     `mapstructure:"export_translations"`

	// Jamf Pro configuration
	JamfInstanceURL  string `mapstructure:"jamf_instance_url"`
	JamfClientID     string `mapstructure:"jamf_client_id"`
	JamfClientSecret string `mapstructure:"jamf_client_secret"`
	JamfUsername     string `mapstructure:"jamf_username"`
	JamfPassword     string `mapstructure:"jamf_password"`
	AuthMethod       string `mapstructure:"auth_method"`

	// Advanced Jamf Pro client settings
	MaxRetryAttempts            int  `mapstructure:"max_retry_attempts"`
	EnableDynamicRateLimiting   bool `mapstructure:"enable_dynamic_rate_limiting"`
	MaxConcurrentRequests       int  `mapstructure:"max_concurrent_requests"`
	TokenRefreshBufferSeconds   int  `mapstructure:"token_refresh_buffer_seconds"`
	TotalRetryDurationSeconds   int  `mapstructure:"total_retry_duration_seconds"`
	CustomTimeoutSeconds        int  `mapstructure:"custom_timeout_seconds"`
	FollowRedirects             bool `mapstructure:"follow_redirects"`
	MaxRedirects                int  `mapstructure:"max_redirects"`
	EnableConcurrencyManagement bool `mapstructure:"enable_concurrency_management"`
	JamfLoadBalancerLock        bool `mapstructure:"jamf_load_balancer_lock"`
	HideSensitiveData           bool `mapstructure:"hide_sensitive_data"`

	// Tool description overrides
	ToolDescriptions map[string]string `mapstructure:"tool_descriptions"`
}

// Load loads configuration from various sources
func Load(cmd *cobra.Command) (*Config, error) {
	v := viper.New()

	// Set configuration file path
	configFile, _ := cmd.Flags().GetString("config")
	if configFile != "" {
		v.SetConfigFile(configFile)
	} else {
		// Look for config file in current directory
		v.SetConfigName("jamfpro-mcp-server-config")
		v.SetConfigType("json")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME/.jamfpro-mcp-server")
		v.AddConfigPath("/etc/jamfpro-mcp-server")
	}

	// Read config file if it exists
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Set up environment variable mapping
	v.SetEnvPrefix("JAMF")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// Map environment variables to config keys
	envMappings := map[string]string{
		"JAMF_INSTANCE_URL":                  "jamf_instance_url",
		"JAMF_CLIENT_ID":                     "jamf_client_id",
		"JAMF_CLIENT_SECRET":                 "jamf_client_secret",
		"JAMF_USERNAME":                      "jamf_username",
		"JAMF_PASSWORD":                      "jamf_password",
		"JAMF_AUTH_METHOD":                   "auth_method",
		"JAMF_TOOLSETS":                      "toolsets",
		"JAMF_DYNAMIC_TOOLSETS":              "dynamic_toolsets",
		"JAMF_LOG_LEVEL":                     "log_level",
		"JAMF_MAX_RETRY_ATTEMPTS":            "max_retry_attempts",
		"JAMF_ENABLE_DYNAMIC_RATE_LIMITING":  "enable_dynamic_rate_limiting",
		"JAMF_MAX_CONCURRENT_REQUESTS":       "max_concurrent_requests",
		"JAMF_TOKEN_REFRESH_BUFFER_SECONDS":  "token_refresh_buffer_seconds",
		"JAMF_TOTAL_RETRY_DURATION_SECONDS":  "total_retry_duration_seconds",
		"JAMF_CUSTOM_TIMEOUT_SECONDS":        "custom_timeout_seconds",
		"JAMF_FOLLOW_REDIRECTS":              "follow_redirects",
		"JAMF_MAX_REDIRECTS":                 "max_redirects",
		"JAMF_ENABLE_CONCURRENCY_MANAGEMENT": "enable_concurrency_management",
		"JAMF_LOAD_BALANCER_LOCK":            "jamf_load_balancer_lock",
		"JAMF_HIDE_SENSITIVE_DATA":           "hide_sensitive_data",
	}

	for envVar, configKey := range envMappings {
		if value := os.Getenv(envVar); value != "" {
			v.Set(configKey, value)
		}
	}

	// Bind command line flags
	if err := bindFlags(cmd, v); err != nil {
		return nil, fmt.Errorf("failed to bind flags: %w", err)
	}

	// Set defaults
	setDefaults(v)

	// Unmarshal into config struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Handle special cases for toolsets from environment variable
	if toolsetsEnv := os.Getenv("JAMF_TOOLSETS"); toolsetsEnv != "" {
		cfg.Toolsets = strings.Split(toolsetsEnv, ",")
		for i, toolset := range cfg.Toolsets {
			cfg.Toolsets[i] = strings.TrimSpace(toolset)
		}
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) error {
	flagMappings := map[string]string{
		"log-level":           "log_level",
		"toolsets":            "toolsets",
		"dynamic-toolsets":    "dynamic_toolsets",
		"export-translations": "export_translations",
		"jamf-instance-url":   "jamf_instance_url",
		"jamf-client-id":      "jamf_client_id",
		"jamf-client-secret":  "jamf_client_secret",
		"jamf-username":       "jamf_username",
		"jamf-password":       "jamf_password",
		"auth-method":         "auth_method",
	}

	for flag, configKey := range flagMappings {
		if err := v.BindPFlag(configKey, cmd.Flags().Lookup(flag)); err != nil {
			return err
		}
	}

	return nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("log_level", "info")
	v.SetDefault("toolsets", []string{"all"})
	v.SetDefault("dynamic_toolsets", false)
	v.SetDefault("export_translations", false)
	v.SetDefault("auth_method", "oauth2")
	v.SetDefault("max_retry_attempts", 3)
	v.SetDefault("enable_dynamic_rate_limiting", true)
	v.SetDefault("max_concurrent_requests", 5)
	v.SetDefault("token_refresh_buffer_seconds", 300)
	v.SetDefault("total_retry_duration_seconds", 300)
	v.SetDefault("custom_timeout_seconds", 60)
	v.SetDefault("follow_redirects", true)
	v.SetDefault("max_redirects", 5)
	v.SetDefault("enable_concurrency_management", true)
	v.SetDefault("jamf_load_balancer_lock", false)
	v.SetDefault("hide_sensitive_data", true)
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.JamfInstanceURL == "" {
		return fmt.Errorf("jamf_instance_url is required")
	}

	if !strings.HasPrefix(c.JamfInstanceURL, "http://") && !strings.HasPrefix(c.JamfInstanceURL, "https://") {
		return fmt.Errorf("jamf_instance_url must start with http:// or https://")
	}

	switch c.AuthMethod {
	case "oauth2":
		if c.JamfClientID == "" {
			return fmt.Errorf("jamf_client_id is required for oauth2 authentication")
		}
		if c.JamfClientSecret == "" {
			return fmt.Errorf("jamf_client_secret is required for oauth2 authentication")
		}
	case "basic":
		if c.JamfUsername == "" {
			return fmt.Errorf("jamf_username is required for basic authentication")
		}
		if c.JamfPassword == "" {
			return fmt.Errorf("jamf_password is required for basic authentication")
		}
	default:
		return fmt.Errorf("auth_method must be either 'oauth2' or 'basic'")
	}

	if len(c.Toolsets) == 0 {
		return fmt.Errorf("at least one toolset must be specified")
	}

	return nil
}

// GetJamfProClientConfig returns a config suitable for the Jamf Pro SDK client
func (c *Config) GetJamfProClientConfig() map[string]interface{} {
	config := map[string]interface{}{
		"instance_domain":                     c.JamfInstanceURL,
		"auth_method":                         c.AuthMethod,
		"log_level":                           c.LogLevel,
		"log_output_format":                   "json",
		"hide_sensitive_data":                 c.HideSensitiveData,
		"max_retry_attempts":                  c.MaxRetryAttempts,
		"enable_dynamic_rate_limiting":        c.EnableDynamicRateLimiting,
		"max_concurrent_requests":             c.MaxConcurrentRequests,
		"token_refresh_buffer_period_seconds": c.TokenRefreshBufferSeconds,
		"total_retry_duration_seconds":        c.TotalRetryDurationSeconds,
		"custom_timeout_seconds":              c.CustomTimeoutSeconds,
		"follow_redirects":                    c.FollowRedirects,
		"max_redirects":                       c.MaxRedirects,
		"enable_concurrency_management":       c.EnableConcurrencyManagement,
		"jamf_load_balancer_lock":             c.JamfLoadBalancerLock,
	}

	if c.AuthMethod == "oauth2" {
		config["client_id"] = c.JamfClientID
		config["client_secret"] = c.JamfClientSecret
	} else {
		config["basic_auth_username"] = c.JamfUsername
		config["basic_auth_password"] = c.JamfPassword
	}

	return config
}
