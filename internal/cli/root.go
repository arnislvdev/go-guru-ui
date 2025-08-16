package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	mode    string
	verbose bool
)

// This is the main command - what runs when you just type 'guruui'
var rootCmd = &cobra.Command{
	Use:   "guruui",
	Short: "GuruUI - Makes errors easy to understand and turns words into commands",
	Long: `GuruUI is a tool that explains errors in simple English 
and turns normal words into computer commands.

What it does:
  • Professional Mode: Clear and simple explanations
  • WTF Mode: Same explanations but funny
  • Error Help: Understand what went wrong
  • Command Help: Turn words into commands

Examples:
  guruui explain "undefined: fmt"
  guruui translate "how do I check disk space"
  guruui --mode wtf explain "segmentation fault"`,
	Version: "0.1.0",
}

// Execute runs the main command and adds all the smaller commands
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.guruui.yaml)")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "professional", "output mode: professional or wtf")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Add subcommands
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(translateCmd)
	rootCmd.AddCommand(configCmd)
}

// initConfig reads the settings file and environment variables
func initConfig() {
	if cfgFile != "" {
		// Use the settings file they specified
		viper.SetConfigFile(cfgFile)
	} else {
		// Find their home folder
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Look for settings file called ".guruui" in their home folder
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".guruui")
	}

	viper.AutomaticEnv() // read environment variables

	// If we found a settings file, read it
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Using settings file: %s\n", viper.ConfigFileUsed())
		}
	}
}
