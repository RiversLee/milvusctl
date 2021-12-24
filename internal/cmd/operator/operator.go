package operator

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func NewOperatorCmd(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	operatorCmd := &cobra.Command{
		Use: "operator",
		Short: "command related to The Milvus operator",
		Run: runHelp,
	}
	operatorCmd.AddCommand(NewOperatorInstallCmd(f,ioStreams))
	operatorCmd.AddCommand(NewOperatorUninstallCmd(f,ioStreams))
	return operatorCmd
}
func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}