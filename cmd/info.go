package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	jsoniter "github.com/json-iterator/go"
)

var infoCmd = &cobra.Command{
	Use: "info",
	Run: infoRunner,
}

func infoRunner(cmd *cobra.Command, args []string) {
	container, err := client.LoadContainer(rootCtx, "helloweb")
	if err != nil {
		log.Fatalf("Fail to load container helloweb: %v\n", err)
	}

	info, err := container.Info(rootCtx)
	if err != nil {
		log.Fatalf("Fail to get container info")
	}

	infoJSON, err := jsoniter.MarshalToString(info)
	if err != nil {
		log.Fatalf("Fail to marshal info\n")
	}

	fmt.Printf("%v\n", infoJSON)
}
