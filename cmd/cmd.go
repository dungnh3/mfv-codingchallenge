package cmd

import "github.com/spf13/cobra"

var usage = `mfv-coding-challenge
	- server: start server
@money-forward-vn
`

// New rootCommand
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Short:        "mfv platform",
		Long:         usage,
		SilenceUsage: true,
	}
	cmd.AddCommand(newServerCmd())
	return cmd
}
