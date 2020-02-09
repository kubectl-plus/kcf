package main

import (
	"github.com/kubectl-plus/kcf/cmd/fleet/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // required for GKE
)

func main() {
	cli.InitAndExecute()
}
