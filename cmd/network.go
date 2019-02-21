package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/google/uuid"

	"github.com/containerd/containerd/cio"
	gocni "github.com/containerd/go-cni"
)

var netCmd = &cobra.Command{
	Use: "net",
	Run: netRunner,
}

func netRunner(cmd *cobra.Command, args []string) {
	container, err := client.LoadContainer(rootCtx, "helloweb")
	if err != nil {
		log.Fatalf("Fail to load container helloweb: %v\n", err)
	}

	task, err := container.NewTask(rootCtx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		log.Fatalf("Fail to create task: %v", err)
	}

	id := uuid.New().String()
	netns := getNetns(task.Pid())

	cni, err := gocni.New(gocni.WithPluginConfDir("./net.d/"),
		gocni.WithPluginDir([]string{"/usr/lib/cni"}))

	// Load the cni configuration

	if err := cni.Load(gocni.WithLoNetwork, gocni.WithDefaultConf); err != nil {
		log.Fatalf("Failed to load cni configuration: %v", err)
	}

	result, err := cni.Setup(id, netns)
	if err != nil {
		log.Fatalf("failed to setup network for namespace %q: %v", id, err)
	}

	for name, config := range result.Interfaces {
		fmt.Printf("Config of interface %s: %v\n",
			name, config)
	}

	if err := task.Start(rootCtx); err != nil {
		log.Fatalf("Fail to start task: %v\n", err)
	}
}

func getNetns(pid uint32) string {
	return fmt.Sprintf("/proc/%d/ns/net", pid)
}
