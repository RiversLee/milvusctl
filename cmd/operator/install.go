package operator

import "github.com/spf13/cobra"

type operatorInstallArgs struct {
	//crFilename is the path to the input milvusOperator CR.
	crFilename string

	//kubeConfigPath is the path to kube config file.
	kubeConfigPath string

	//context is the cluster context in the kube config
	context string

}
func addOperatorInstallFlags(cmd *cobra.Command,args *operatorInstallArgs) {
	cmd.PersistentFlags().StringVarP(&args.crFilename, "filename", "f","",filenameFlagHelpStr)
	cmd.PersistentFlags().StringVarP(&args.kubeConfigPath,"kubeconfig","c","",KubeConfigFlagHelpStr)
	cmd.PersistentFlags().StringVar(&args.context,"context","",ContextFlagHelpStr)
}
func operatorInstallCmd(args []string) *cobra.Command {
	installCmd := &cobra.Command{
		Use: "install",
		Short: "Install the milvus operator controller in the cluster",
		Long: "The install subcommand installs the milvus operator controller in the cluster",
	}
	return installCmd
}

//operatorInstall installs the milvus operator controller into the cluster.
func operatorInstall(args *rootArgs,oiArgs *operatorInstallArgs) {
	config,err := BuildClientConfig(oiArgs.kubeConfigPath,oiArgs.context)
	if err != nil {
		return
	}
	if clientset,err := GetClientSet(config); err != nil {
		return
	}
}