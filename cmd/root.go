package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

var (
	client *containerd.Client
	rootCtx context.Context
)

var rootCmd = &cobra.Command{ Use: "catraia" }

func init() {
	c, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatalf("fail to connect to containerd: %v", err)
	}

	client = c

	rootCtx = namespaces.WithNamespace(context.Background(), "default")

	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(netCmd)
}

// Execute is the entry point of the application
func Execute() {
	if err:= rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
