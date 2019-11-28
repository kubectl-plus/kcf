package fleet

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Resources creates a detailed report of the resources in
// a particular cluster in a fleet
func Resources(configFlags *genericclioptions.ConfigFlags, args []string) error {
	clientcfg := configFlags.ToRawKubeConfigLoader()
	cfg, err := clientcfg.RawConfig()
	if err != nil {
		return errors.Wrap(err, "Can't assemble raw config")
	}
	if len(args) < 1 {
		return errors.New("need a cluster to operate on, please provide the cluster name")
	}
	clusterID := args[0]
	context := contextOf(cfg, clusterID)
	err = resourceDetails(cfg, context)
	if err != nil {
		return err
	}
	return nil
}

// resourceDetails prints the supported resources in the cluster
func resourceDetails(cfg api.Config, context string) error {
	cs, err := csForContext(cfg, context)
	if err != nil {
		return errors.Wrap(err, "Can't create a clientset based on config provided")
	}
	_, reslist, err := cs.Discovery().ServerGroupsAndResources()
	if err != nil {
		return errors.Wrap(err, "Can't get cluster server version")
	}
	fmt.Println("Resources supported in this cluster:")
	for _, res := range reslist {
		fmt.Println(strings.Repeat("-", 80))
		fmt.Printf("%v:\n ", res.GroupVersion)
		for _, r := range res.APIResources {
			if !strings.Contains(r.Name, "/") {
				fmt.Printf("%v (namespaced: %v) ", r.Name, r.Namespaced)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println(strings.Repeat("*", 80))
	return nil
}
