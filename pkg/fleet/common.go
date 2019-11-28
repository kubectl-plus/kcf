package fleet

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// csForContext returns a client for a given context
func csForContext(cfg api.Config, context string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.NewNonInteractiveClientConfig(
		cfg,
		context,
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		},
		nil).ClientConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Can't switch context")
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "Can't create a client based on config and/or context provided")
	}
	return cs, nil
}

// contextOf returns the context name of a given cluster
func contextOf(cfg api.Config, clusterID string) string {
	for name, context := range cfg.Contexts {
		if clusterID == context.Cluster {
			return name
		}
	}
	return ""
}
