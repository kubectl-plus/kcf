package main

import (
	"github.com/mhausenblas/kcf/cmd/fleet/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // required for GKE
)

func main() {
	cli.InitAndExecute()
}
