package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage GuruUI configuration",
	Long: `Configure GuruUI settings including AI provider API keys and preferences.
	
Examples:
  guruui config set --ai-provider openai --api-key <your-key>
  guruui config show
  guruui config reset`,
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration values",
	Long: `Set configuration values for GuruUI.
	
Examples:
  guruui config set --ai-provider openai --api-key <your-key>
  guruui config set --default-mode wtf`,
	RunE: func(cmd *cobra.Command, args []string) error {
		aiProvider, _ := cmd.Flags().GetString("ai-provider")
		apiKey, _ := cmd.Flags().GetString("api-key")
		defaultMode, _ := cmd.Flags().GetString("default-mode")

		if aiProvider != "" {
			viper.Set("ai.provider", aiProvider)
		}
		if apiKey != "" {
			viper.Set("ai.api_key", apiKey)
		}
		if defaultMode != "" {
			viper.Set("default_mode", defaultMode)
		}

		// Save config
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("Configuration saved successfully!")
		return nil
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Current Configuration:")
		fmt.Printf("AI Provider: %s\n", viper.GetString("ai.provider"))
		fmt.Printf("Default Mode: %s\n", viper.GetString("default_mode"))
		fmt.Printf("Config File: %s\n", viper.ConfigFileUsed())
		return nil
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to defaults",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Reset to defaults
		viper.Set("ai.provider", "openai")
		viper.Set("default_mode", "professional")

		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("Configuration reset to defaults!")
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configResetCmd)

	configSetCmd.Flags().String("ai-provider", "", "AI provider (openai, anthropic)")
	configSetCmd.Flags().String("api-key", "", "API key for AI provider")
	configSetCmd.Flags().String("default-mode", "", "default output mode (professional, wtf)")
}
