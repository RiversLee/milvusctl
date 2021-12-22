package operator

import (
	"fmt"
	"github.com/spf13/cobra"
)
const (
	KubeConfigFlagHelpStr = "Path to kube config."
	ContextFlagHelpStr = "The name of the kubeconfig context to use"
	filenameFlagHelpStr = `Path to file containing MilvusOperator custom resource
		This flag can be specified multipe times to overlay multiple files. Multiple files are overlaid in left to right order`
)
type rootArgs struct {
	//Dry run performs all steps except actually applying the manifests or creating output dirs/files
	dryRun bool
}

func GetOperatorCmd(args []string) *cobra.Command {
	operatorCmd := &cobra.Command{
		Use: "operator",
		Short: "command related to The Milvus operator",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("it's milvus operator")
		},
	}
	return operatorCmd
}
