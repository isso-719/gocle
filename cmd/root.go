package cmd

import (
	"gocle/src"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocle [package name]",
	Short: "Delete Packages from GOBIN.",
	Long: `Gocle is a CLI tool to remove the Go Packages installed on GOBIN.
For example,
 $ gocle air  // remove air package from GOBIN
 $ gocle air -y // remove air package from GOBIN, and don't show confirmation message
 $ gocle air golangci-lint // you can remove multiple packages at once
`,
	Run: func(cmd *cobra.Command, args []string) {
		src.Remover(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("yes", "y", false, "don't show confirmation message")
}
