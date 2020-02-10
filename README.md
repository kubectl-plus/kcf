# Kubernetes cluster fleet viewer

Clusters are the new cattle and we should have tooling available that allows us to quickly get an idea what's going in a fleet of such clusters.

## What is this about?

Meet `fleet`, a simple CLI tool that provides you with the status and configuration of a fleet of Kubernetes clusters. For example: 

```sh
$ kubectl fleet
CLUSTER                                                      VERSION       NODES NAMESPACES PROVIDER      API
kind-kind-3                                                  v1.16.3       1/1   4          kind          https://127.0.0.1:32769
test-cluster-2                                               v1.16.2       1/1   4          minikube      https://192.168.64.4:8443
kind-test2                                                   v1.16.3       1/1   4          kind          https://127.0.0.1:32768
minikube                                                     v1.16.2       1/1   4          minikube      https://192.168.64.3:8443
gke_krew-release-bot-260708_us-central1-a_standard-cluster-1 v1.15.8-gke.3 3/3   4          GKE           https://104.197.42.183
do-sfo2-k8s-1-16-6-do-0-sfo2-1581265844177                   v1.16.6       3/3   4          Digital Ocean https://f048f314-4f77-47c2-9264-764da91d35e0.k8s.ondigitalocean.com
```

Above, you see `fleet` used as a `kubectl` plugin, available via [krew](http://krew.dev/). 
The top-level command lists all active clusters found in the [kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)
provided. Active clusters are defined as the one that you would see when you'd execute
the `kubectl config get-contexts` command. For each cluster, configuration info such as
the control plane version or API server endpoint are displayed, as well as select
stats, for example, the number of worker nodes or namespaces found in the cluster.

Note that you can also use it standalone, simply download the binary for your platform
from the [release page](https://github.com/kubectl-plus/kcf/releases).

## Getting started

To get started, visit the [usage docs](doc/USAGE.md).


