package fleet

import (
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// RunFleetCommand runs the top-level fleet command
func RunFleetCommand(configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read kubeconfig")
	}
	fmt.Printf("NAME\tCLUSTER\tSERVER\n")
	for name, context := range config.Contexts {
		cluster := config.Clusters[context.Cluster]
		server := "unknown"
		if cluster != nil {
			server = cluster.Server
		}
		fmt.Printf("%v\t%v\t%v\n", name, context.Cluster, server)
	}
	return nil
}
