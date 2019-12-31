package fleet

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Overview creates an tabular overview of all the clusters in a fleet.
// A fleet is defined as the active clusters in the kubeconfig provided, that
// is, what you see when you execute: kubectl config get-contexts
func Overview(configFlags *genericclioptions.ConfigFlags) error {
	clientcfg := configFlags.ToRawKubeConfigLoader()
	cfg, err := clientcfg.RawConfig()
	if err != nil {
		return errors.Wrap(err, "Can't assemble raw config")
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "CLUSTER\tVERSION\tNODES\tNAMESPACES\tAPI")
	for name, context := range cfg.Contexts {
		cluster := cfg.Clusters[context.Cluster]
		clusterVersion, err := clusterVersion(cfg, name)
		if err != nil {
			clusterVersion = "?"
		}
		noinfo, err := nodesOverview(cfg, name)
		if err != nil {
			noinfo = "?"
		}
		nsinfo, err := nsOverview(cfg, name)
		if err != nil {
			nsinfo = "?"
		}
		apiServerEndpoint := "?"
		if cluster != nil {
			apiServerEndpoint = cluster.Server
		}
		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v\t%v\t%v", context.Cluster, clusterVersion, noinfo, nsinfo, apiServerEndpoint))
	}
	w.Flush()
	return nil
}

// clusterVersion returns the cluster version in given context
func clusterVersion(cfg api.Config, context string) (string, error) {
	cs, err := csForContext(cfg, context)
	if err != nil {
		return "", errors.Wrap(err, "Can't create a clientset based on config provided")
	}
	cversion, err := cs.Discovery().ServerVersion()
	if err != nil {
		return "", errors.Wrap(err, "Can't get cluster server version")
	}
	return fmt.Sprintf("%s", cversion), nil
}

// nodesOverview returns the cluster's worker nodes overview in given context
func nodesOverview(cfg api.Config, context string) (string, error) {
	cs, err := csForContext(cfg, context)
	if err != nil {
		return "", errors.Wrap(err, "Can't create a clientset based on config provided")
	}
	nodes, err := cs.CoreV1().Nodes().List(metav1.ListOptions{})
	nodeCount := len(nodes.Items)
	readyCount := 0
	for _, node := range nodes.Items {
		for _, nodeCondition := range node.Status.Conditions {
			if nodeCondition.Type == "Ready" {
				if nodeCondition.Status == "True" {
					readyCount++
				}
				break
			}
		}
	}
	if err != nil {
		return "", errors.Wrap(err, "Can't get nodes in cluster")
	}
	noverview := fmt.Sprintf("%v/%v", readyCount, nodeCount)
	return noverview, nil
}

// nsOverview returns the cluster's namespaces overview in given context
func nsOverview(cfg api.Config, context string) (string, error) {
	cs, err := csForContext(cfg, context)
	if err != nil {
		return "", errors.Wrap(err, "Can't create a clientset based on config provided")
	}
	ns, err := cs.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return "", errors.Wrap(err, "Can't get namespaces in cluster")
	}
	nsverview := fmt.Sprintf("%v", len(ns.Items))
	return nsverview, nil
}
