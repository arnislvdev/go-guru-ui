package cli

import (
	"fmt"

	"github.com/arnislvdev/go-guru-ui/internal/usecase"
	"github.com/spf13/cobra"
)

var translateCmd = &cobra.Command{
	Use:   "translate [natural_language_query]",
	Short: "Turn normal words into computer commands",
	Long: `Turn normal words into computer commands you can run.
	
Examples:
  guruui translate "how do I check disk space"
  guruui translate "install nginx on ubuntu"
  guruui translate --context "I'm on macOS" "check network connections"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		// Get extra info
		context, _ := cmd.Flags().GetString("context")

		// Make the translator
		translator := usecase.NewCommandTranslator()

		// Turn words into command
		command, explanation, err := translator.Translate(query, context, mode)
		if err != nil {
			return fmt.Errorf("failed to translate query: %w", err)
		}

		// Show the result
		fmt.Printf("Command: %s\n\n", command)
		fmt.Printf("Explanation: %s\n", explanation)
		return nil
	},
}

func init() {
	translateCmd.Flags().StringP("context", "c", "", "additional context (e.g., operating system, environment)")
}
