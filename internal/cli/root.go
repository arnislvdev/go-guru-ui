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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "guruui",
	Short: "GuruUI - AI-powered error explanation and command translation",
	Long: `GuruUI is a developer tool that explains errors in plain English 
and translates natural language into CLI commands.

Features:
  • Professional Mode: Clear, concise, beginner-friendly explanations
  • WTF Mode: Same explanations with humor, sarcasm, and memes
  • Error Analysis: Understand compilation and runtime errors
  • Command Translation: Convert natural language to CLI commands

Examples:
  guruui explain "undefined: fmt"
  guruui translate "how do I check disk space"
  guruui --mode wtf explain "segmentation fault"`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".guruui" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".guruui")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Using config file: %s\n", viper.ConfigFileUsed())
		}
	}
}
