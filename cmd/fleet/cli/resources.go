package cli

import (
	"github.com/mhausenblas/kcf/pkg/fleet"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ResourcesCmd runs the fleet resources command
func ResourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "resources",
		Short:         "Details about the resources of a cluster in the fleet",
		Long:          `.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fleet.Resources(KubernetesConfigFlags, args); err != nil {
				return errors.Cause(err)
			}
			return nil
		},
	}
	return cmd
}
