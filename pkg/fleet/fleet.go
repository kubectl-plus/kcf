package fleet

import (
	"fmt"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"os"
	"text/tabwriter"
)

// RunFleetCommand runs the top-level fleet command
func RunFleetCommand(configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return errors.Wrap(err, "Can't read kubeconfig.")
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
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

	nodesOverview(configFlags)
	return nil
}

func nodesOverview(configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return errors.Wrap(err, "Can't read kubeconfig")
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "Can't create a clientset based on kubeconfig provided.")
	}

	nodes, err := cs.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "Can't list nodes in cluster")
	}

	for _, node := range nodes.Items {
		fmt.Println(node.Name)
	}

	return nil
}
