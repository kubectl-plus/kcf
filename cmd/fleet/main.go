package main

import (
	"github.com/mhausenblas/kcf/cmd/fleet/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"  // required for GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc" // required for oidc authentication
)

func main() {
	cli.InitAndExecute()
}
