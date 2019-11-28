package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DetailsCmd runs the fleet details command
func DetailsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "details",
		Short:         "Details about a cluster in the fleet",
		Long:          `.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("DETAILS")
			return nil
		},
	}
	return cmd
}
