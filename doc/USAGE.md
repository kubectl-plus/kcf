# Usage

The following assumes you have the plugin from this repo:

```sh
$ git clone https://github.com/mhausenblas/kcf.git && cd kcf

$ kubectl krew install --manifest=deploy/krew/fleet.yaml
```

## Cluster fleet overview

```sh
$ kubectl fleet
NAME                                         CLUSTER                            SERVER
kind-mh9local                                kind-mh9local                      https://127.0.0.1:56208
example@something.us-west-2.eksctl.io        something.us-west-2.eksctl.io      https://123456789ABCDEF.gr7.us-west-2.eks.amazonaws.com
```


