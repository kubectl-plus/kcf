package cli

import (
	"github.com/kubectl-plus/kcf/pkg/fleet"
	"github.com/pkg/errors"
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
			if err := fleet.Details(KubernetesConfigFlags, args); err != nil {
				return errors.Cause(err)
			}
			return nil
		},
	}
	return cmd
}
