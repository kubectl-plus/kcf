package main

import (
	"github.com/kubectl-plus/kcf/cmd/fleet/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // required for GKE/OIDC/Dex style auth
)

func main() {
	cli.InitAndExecute()
}
