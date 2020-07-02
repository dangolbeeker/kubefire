package cluster

import (
	"github.com/innobead/kubefire/internal/di"
	"github.com/innobead/kubefire/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var downloadCmd = &cobra.Command{
	Use:   "download [name]",
	Short: "Download the kubeconfig of cluster",
	Args:  util.Validate1thArg("name"),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cluster, err := di.ClusterManager().Get(name)
		if err != nil {
			return errors.WithMessagef(err, "failed to get cluster (%s) info", name)
		}

		wd, _ := os.Getwd()
		if err := di.Bootstrapper().DownloadKubeConfig(cluster, wd); err != nil {
			return errors.WithMessagef(err, "failed to download kubeconfig of cluster (%s)", name)
		}

		return nil
	},
}
