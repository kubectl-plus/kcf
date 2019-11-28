package fleet

import (
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Details creates a detailed report on a particular clusters in a fleet
func Details(configFlags *genericclioptions.ConfigFlags, args []string) error {
	clientcfg := configFlags.ToRawKubeConfigLoader()
	cfg, err := clientcfg.RawConfig()
	if err != nil {
		return errors.Wrap(err, "Can't assemble raw config")
	}
	if len(args) < 1 {
		return errors.New("need a cluster to operate on, please provide the cluster name")
	}
	clusterID := args[0]
	cluster := cfg.Clusters[clusterID]
	fmt.Printf("API server endpoint: %v\n", cluster.Server)
	// context := contextOf(cfg, clusterID)
	return nil
}
