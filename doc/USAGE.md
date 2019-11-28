# Usage

The following assumes you have the plugin from this repo:

```sh
$ git clone https://github.com/mhausenblas/kcf.git && cd kcf

$ kubectl krew install --manifest=deploy/krew/fleet.yaml
```

## Cluster fleet overview

```sh
$ kubectl fleet
CLUSTER                     VERSION            NODES NAMESPACES API
mngbase.us-west-2.eksctl.io v1.14.8-eks-41be3d 2     5          https://123456789ABCDEF.gr7.us-west-2.eks.amazonaws.com
kind-mh9local               v1.16.3            1     4          https://127.0.0.1:58836
```


