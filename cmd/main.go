package main

import (
	"fmt"
	"os"

	"github.com/elliotxx/go-cli-prototype/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cmdShort   = i18n.T(`A cli application`)
	cmdLong    = templates.LongDesc(i18n.T(`This is a cli application with go and cobra.`))
	cmdExample = templates.Examples(i18n.T(`
		# Show version info
		go-cli-prototype -V
		# Show echo info
		go-cli-prototype -e "hello world"`))
)

type MainOptions struct {
	ShowVersion bool
	Echo        string
}

// NewEncryptOptions returns a new EncryptOptions instance
func NewMainOptions() *MainOptions {
	return &MainOptions{}
}

func configureCLI() *cobra.Command {
	o := NewMainOptions()
	rootCmd := &cobra.Command{
		Use:          "go-cli-prototype",
		Short:        cmdShort,
		Long:         cmdLong,
		Example:      cmdExample,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Show version info
			if o.ShowVersion {
				fmt.Println(version.ReleaseVersion())
				return nil
			}
			if len(o.Echo) > 0 {
				fmt.Println(o.Echo)
				return nil
			}
			return nil
		},
	}

	rootCmd.Flags().BoolVarP(&o.ShowVersion, "version", "V", false, "Show version info")
	rootCmd.Flags().StringVarP(&o.Echo, "echo", "e", "", "Show echo info")
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
