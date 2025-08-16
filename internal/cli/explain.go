package cli

import (
	"fmt"

	"github.com/arnislvdev/go-guru-ui/internal/usecase"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain [error_message]",
	Short: "Explain an error message in plain English",
	Long: `Explain a programming error in clear, understandable terms.
	
Examples:
  guruui explain "undefined: fmt"
  guruui explain "cannot use nil as type string in assignment"
  guruui explain --file main.go --line 42`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		errorMsg := args[0]

		// Get flags
		file, _ := cmd.Flags().GetString("file")
		line, _ := cmd.Flags().GetInt("line")

		// Create explainer service
		explainer := usecase.NewErrorExplainer()

		// Explain the error
		explanation, err := explainer.Explain(errorMsg, file, line, mode)
		if err != nil {
			return fmt.Errorf("failed to explain error: %w", err)
		}

		// Output the explanation
		fmt.Println(explanation)
		return nil
	},
}

func init() {
	explainCmd.Flags().StringP("file", "f", "", "source file path for context")
	explainCmd.Flags().IntP("line", "l", 0, "line number for context")
}
