# Kubernetes cluster fleet viewer

Clusters are the new cattle and we should have tooling available that allows us to quickly get an idea what's going in a fleet of such clusters.

## What is this about?

Meet `fleet`, a simple CLI tool that provides you with the status and configuration of a fleet of Kubernetes clusters. For example: 

```sh
$ kubectl fleet
CLUSTER                     VERSION            NODES NAMESPACES API
mngbase.us-west-2.eksctl.io v1.14.8-eks-41be3d 2     5          https://123456789ABCDEF.gr7.us-west-2.eks.amazonaws.com
kind-mh9local               v1.16.3            1     4          https://127.0.0.1:58836
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


