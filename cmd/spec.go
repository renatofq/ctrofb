package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	jsoniter "github.com/json-iterator/go"
)

var specCmd = &cobra.Command{
	Use: "spec",
	Run: specRunner,
}

func specRunner(cmd *cobra.Command, args []string) {
	container, err := client.LoadContainer(rootCtx, "helloweb")
	if err != nil {
		log.Fatalf("Fail to load container helloweb: %v\n", err)
	}

	spec, err := container.Spec(rootCtx)
	if err != nil {
		log.Fatalf("Fail to get container spec")
	}

	specJSON, err := jsoniter.MarshalToString(spec)
	if err != nil {
		log.Fatalf("Fail to marshal spec\n")
	}

	fmt.Printf("%v\n", specJSON)
}
