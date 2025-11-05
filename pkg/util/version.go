package util

import (
	"fmt"
	"os"
	"text/tabwriter"

	"kubectl-multi/pkg/cluster"

	"github.com/spf13/cobra"
)

var Version = "dev" // overridden by goreleaser during build process

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of kubectl plugin and Kubernetes servers",
	Long:  `Shows the client version of kubectl-multi and server versions from all managed clusters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get global flags from root command
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		_, _ = cmd.Flags().GetString("remote-context") // Ignored, using empty to avoid warnings

		// Display client version
		fmt.Printf("Client Version: %s\n\n", Version)

		// Discover clusters and display their server versions
		// Pass empty remoteCtx to avoid KubeStellar managed cluster discovery warnings
		clusters, err := cluster.DiscoverClusters(kubeconfig, "")
		if err != nil {
			// Silently continue if cluster discovery fails
			fmt.Println("Server Versions: (no clusters available)")
			return nil
		}

		if len(clusters) == 0 {
			fmt.Println("No clusters found")
			return nil
		}

		// Display server versions
		fmt.Println("Server Versions:")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		for _, clusterInfo := range clusters {
			if clusterInfo.DiscoveryClient != nil {
				serverVersion, err := clusterInfo.DiscoveryClient.ServerVersion()
				if err != nil {
					fmt.Fprintf(w, "%s\tError: %v\n", clusterInfo.Name, err)
					continue
				}
				fmt.Fprintf(w, "%s\t%s\n", clusterInfo.Name, serverVersion.GitVersion)
			}
		}
		w.Flush()
		return nil
	},
}
