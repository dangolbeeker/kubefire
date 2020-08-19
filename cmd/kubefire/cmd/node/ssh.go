package node

import (
	"github.com/innobead/kubefire/internal/di"
	"github.com/innobead/kubefire/internal/validate"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh [name]",
	Short: "SSH into node",
	Args:  validate.OneArg("name"),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validate.NodeExist(args[0])
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return di.ClusterManager().GetNodeManager().LoginBySSH(
			args[0],
			di.ClusterManager().GetConfigManager(),
		)
	},
}
