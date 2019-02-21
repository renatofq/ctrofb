package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/oci"
)

var createCmd = &cobra.Command{
	Use: "create",
	Run: createRunner,
}


func createRunner(cmd *cobra.Command, args []string) {
	image, err := client.GetImage(rootCtx, "docker.io/renatofq/helloweb:latest")
	if err != nil {
		log.Fatalf("Fail to get image: %v\n", err)
	}

	// create a container
	_, err = client.NewContainer(
		rootCtx,
		"helloweb",
		containerd.WithImage(image),
		containerd.WithNewSnapshot("hello-snapshot", image),
		containerd.WithNewSpec(oci.WithImageConfig(image),
			oci.WithCapabilities([]string{"CAP_NET_RAW"})),
	)

	if err != nil {
		log.Fatalf("Fail to create container: %v\n", err)
	}
}
