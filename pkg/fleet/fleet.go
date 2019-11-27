package fleet

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
	"text/tabwriter"
)

// RunFleetCommand runs the top-level fleet command
func RunFleetCommand(configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read kubeconfig")
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "NAME\tCLUSTER\tSERVER")
	for name, context := range config.Contexts {
		cluster := config.Clusters[context.Cluster]
		server := "unknown"
		if cluster != nil {
			server = cluster.Server
		}
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v", name, context.Cluster, server))
	}
	w.Flush()
	return nil
}
