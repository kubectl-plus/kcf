# The `fleet` user manual

## Installation

Most typically, you'd want to install `fleet` as a `kubectl` plugin.
In order to be able to do that, make sure you have [krew installed](https://github.com/kubernetes-sigs/krew/#installation) and then you can install `fleet` as follows:

```sh
$ kubectl krew install fleet
```

Alternatively, for a version that is not yet updated and hence available in the [krew index](https://index.krew.dev/), you can install `fleet` from this repo as follows:

```sh
$ git clone https://github.com/mhausenblas/kcf.git && cd kcf

$ kubectl krew install --manifest=deploy/krew/fleet.yaml
```

If the installation fails, check if `krew` is available on your local system
 and also, make sure you're using the most recent index (run `kubectl krew update`) to ensure this.

## Usage

There are two levels `fleet` operates on: the top-level command, without any 
further sub-commands, operates on all the clusters in your fleet 
(all clusters in your contexts, that is, what you see when you execute 
`kubectl config get-contexts`). If a sub-command such as `details` or `resources` 
is provided, it operates on a particular cluster and hence the name of the 
cluster to use needs to be specified.

### Getting an overview of the fleet

To get an overview of your entire fleet, do:

```sh
$ kubectl fleet
CLUSTER                     VERSION            NODES NAMESPACES API
mngbase.us-west-2.eksctl.io v1.14.8-eks-41be3d 2     5          https://123456789ABCDEF.gr7.us-west-2.eks.amazonaws.com
kind-mh9local               v1.16.3            1     4          https://127.0.0.1:58836
```

### Exploring clusters in the fleet

If you're interested in what is running in a specific cluster in your fleet, 
use the `details` sub-command like so:

```sh
$ kubectl fleet details kind-mh9local
API server endpoint: https://127.0.0.1:51366

# namespace [default]
has no pods
has 1 service(s) overall:
- service [kubernetes] of type ClusterIP uses IP 10.96.0.1 and port(s) https TCP/443
--------------------------------------------------------------------------------

# namespace [kube-node-lease]
has no pods
has no services
--------------------------------------------------------------------------------

# namespace [kube-public]
has no pods
has no services
--------------------------------------------------------------------------------

# namespace [kube-system]
has 8 pod(s) overall:
- pod [coredns-5644d7b6d9-kwzw7] is running and uses image(s) k8s.gcr.io/coredns:1.6.2
- pod [coredns-5644d7b6d9-x5nn5] is running and uses image(s) k8s.gcr.io/coredns:1.6.2
- pod [etcd-mh9local-control-plane] is running and uses image(s) k8s.gcr.io/etcd:3.3.15-0
- pod [kindnet-4df5j] is running and uses image(s) kindest/kindnetd:0.5.3@sha256:bc1833b3da442bb639008dd5a62861a0419d3f64b58fce6fb38b749105232555
- pod [kube-apiserver-mh9local-control-plane] is running and uses image(s) k8s.gcr.io/kube-apiserver:v1.16.3
- pod [kube-controller-manager-mh9local-control-plane] is running and uses image(s) k8s.gcr.io/kube-controller-manager:v1.16.3
- pod [kube-proxy-j5dtw] is running and uses image(s) k8s.gcr.io/kube-proxy:v1.16.3
- pod [kube-scheduler-mh9local-control-plane] is running and uses image(s) k8s.gcr.io/kube-scheduler:v1.16.3
has 1 service(s) overall:
- service [kube-dns] of type ClusterIP uses IP 10.96.0.10 and port(s) dns UDP/53 dns-tcp TCP/53 metrics TCP/9153
--------------------------------------------------------------------------------
```

If you want to learn what kinds of resources are supported in a specific cluster
then you'd use the `resources` sub-command as shown below:

```sh
$ kubectl fleet resources kind-mh9local
Resources supported in this cluster:
--------------------------------------------------------------------------------
v1:
 bindings (namespaced: true) componentstatuses (namespaced: false) configmaps (namespaced: true) endpoints (namespaced: true) events (namespaced: true) limitranges (namespaced: true) namespaces (namespaced: false) nodes (namespaced: false
) persistentvolumeclaims (namespaced: true) persistentvolumes (namespaced: false) pods (namespaced: true) podtemplates (namespaced: true) replicationcontrollers (namespaced: true) resourcequotas (namespaced: true) secrets (namespaced: tru
e) serviceaccounts (namespaced: true) services (namespaced: true)
--------------------------------------------------------------------------------
apiregistration.k8s.io/v1:
 apiservices (namespaced: false)
--------------------------------------------------------------------------------

...

coordination.k8s.io/v1:
 leases (namespaced: true)
--------------------------------------------------------------------------------
coordination.k8s.io/v1beta1:
 leases (namespaced: true)
--------------------------------------------------------------------------------
node.k8s.io/v1beta1:
 runtimeclasses (namespaced: false)
********************************************************************************
```

It's in general a good idea, for above `resources` command, to pipe it through 
`less` … because less is more ;).


