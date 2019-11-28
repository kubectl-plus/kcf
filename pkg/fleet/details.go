package fleet

import (
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Details creates detailed information on a particular clusters in a fleet
func Details(configFlags *genericclioptions.ConfigFlags, args []string) error {
	clientcfg := configFlags.ToRawKubeConfigLoader()
	cfg, err := clientcfg.RawConfig()
	if err != nil {
		return errors.Wrap(err, "Can't assemble raw config")
	}
	if len(args) < 1 {
		return errors.New("need a cluster to operate on, please provide the cluster name")
	}
	_ = cfg
	clusterID := args[0]
	fmt.Println(clusterID)
	return nil
}
