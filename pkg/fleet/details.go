package fleet

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd/api"
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
	if cluster == nil {
		return errors.New(fmt.Sprintf("cluster %q not found in kubeconfig", clusterID))
	}

	fmt.Printf("API server endpoint: %v\n", cluster.Server)
	context := contextOf(cfg, clusterID)
	coreres, err := coreResDetails(cfg, context)
	if err != nil {
		return err
	}
	fmt.Printf("\n%v\n", coreres)
	return nil
}

// coreResDetails returns details about useful core resources in given context.
// Useful core resources include pods, services, deployments,
func coreResDetails(cfg api.Config, context string) (result string, err error) {
	cs, err := csForContext(cfg, context)
	if err != nil {
		return "", errors.Wrap(err, "Can't create a clientset based on config provided")
	}
	namespaces, err := cs.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return "", errors.Wrap(err, "Can't get namespaces in cluster")
	}
	for _, ns := range namespaces.Items {
		nsname := ns.Name
		result += fmt.Sprintf("# namespace [%v]\n", nsname)
		// pod stats in namespace:
		pods, err := cs.CoreV1().Pods(nsname).List(metav1.ListOptions{})
		if err != nil {
			return "", errors.Wrap(err, "Can't get pods")
		}
		switch len(pods.Items) {
		case 0:
			result += fmt.Sprintf("has no pods\n")
		default:
			result += fmt.Sprintf("has %v pod(s) overall:\n", len(pods.Items))
			for _, pod := range pods.Items {
				result += fmt.Sprintf("- %v", podInfo(pod))
			}
		}
		// service stats in namespace:
		svcs, err := cs.CoreV1().Services(nsname).List(metav1.ListOptions{})
		if err != nil {
			return "", errors.Wrap(err, "Can't get services")
		}
		switch len(svcs.Items) {
		case 0:
			result += fmt.Sprintf("has no services\n")
		default:
			result += fmt.Sprintf("has %v service(s) overall:\n", len(svcs.Items))
			for _, svc := range svcs.Items {
				result += fmt.Sprintf("- %v", svcInfo(svc))
			}
		}
		// mark end of namespace stats:
		result += strings.Repeat("-", 80) + "\n\n"
	}
	return result, nil
}

// podInfo renders details of the status and config of a pod given
func podInfo(pod v1.Pod) (result string) {
	podname := pod.Name
	podstatus := strings.ToLower(string(pod.Status.Phase))
	images := ""
	for _, container := range pod.Spec.Containers {
		images += fmt.Sprintf("%v ", container.Image)
	}
	result += fmt.Sprintf("pod [%v] is %v and uses image(s) %v\n", podname, podstatus, images)
	return result
}

// svcInfo renders details of the status and config of a service given
func svcInfo(svc v1.Service) (result string) {
	svcname := svc.Name
	svctype := svc.Spec.Type
	svcclusterip := svc.Spec.ClusterIP
	ports := ""
	for _, port := range svc.Spec.Ports {
		ports += fmt.Sprintf(" %v %v/%v", port.Name, port.Protocol, port.Port)
	}
	result += fmt.Sprintf("service [%v] of type %v uses IP %v and port(s)%v\n", svcname, svctype, svcclusterip, ports)
	return result
}
