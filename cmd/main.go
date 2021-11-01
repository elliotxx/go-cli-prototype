package main

import (
	"fmt"
	"os"

	"github.com/elliotxx/go-cli-prototype/pkg/version"
	"github.com/spf13/cobra"
)

var showVersion bool

func configureCLI() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "go-cli-prototype",
		Long:         "This is a cli application with go and cobra.",
		SilenceUsage: true,
		Example:      "go-cli-prototype -h",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Show version info
			if showVersion {
				fmt.Println(version.ReleaseVersion())
				return nil
			}
			return nil
		},
	}

	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show version info")
	return rootCmd
}

func main() {
	// init rootCmd
	rootCmd := configureCLI()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
