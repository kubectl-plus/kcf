# Using `fleet`

## Installation

Most typically, you'd want to install `fleet` as a `kubectl` plugin.
In order to be able to do that, make sure you have [krew installed](https://github.com/kubernetes-sigs/krew/#installation) and then you can install `fleet` from this repo as follows:

```sh
$ git clone https://github.com/mhausenblas/kcf.git && cd kcf

$ kubectl krew install --manifest=deploy/krew/fleet.yaml
```

## Usage

There are two levels `fleet` operates on: the top-level command, without any 
further sub-commands, operates on all the clusters in your fleet 
(all clusters in your contexts, that is, what you see when you execute 
`kubectl config get-contexts`). If a sub-command such as `details` or `resources` 
is provided, it operates on a particular cluster and hence the name of the 
cluster to use needs to be specified.

### Getting an overview of the fleet

Use:

```sh
$ kubectl fleet
CLUSTER                     VERSION            NODES NAMESPACES API
mngbase.us-west-2.eksctl.io v1.14.8-eks-41be3d 2     5          https://123456789ABCDEF.gr7.us-west-2.eks.amazonaws.com
kind-mh9local               v1.16.3            1     4          https://127.0.0.1:58836
```

### Exploring clusters in the fleet

TBD.


