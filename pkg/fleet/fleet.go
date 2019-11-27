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
	for name, cluster := range config.Clusters {
		fmt.Printf("%v\n", name)
		fmt.Printf("Server: %v\n", cluster.Server)
		fmt.Println()
	}
	return nil
}
