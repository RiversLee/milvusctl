package operator

import (
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kubectlcreate "k8s.io/kubectl/pkg/cmd/create"
)


type operatorInstallArgs struct {
	//crFilename is the path to the input milvusOperator CR.
	crFilename string

	//kubeConfigPath is the path to kube config file.
	kubeConfigPath string

	//kubectl create options
	createOptions kubectlcreate.CreateOptions

}
func NewOperatorInstallCmd(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := kubectlcreate.NewCreateOptions(ioStreams)
	installCmd := &cobra.Command{
		Use: "install",
		Short: "Install the milvus operator controller in the cluster",
		Long: "The install subcommand installs the milvus operator controller in the cluster",
		Run: func(cmd *cobra.Command, args []string) {
			if cmdutil.IsFilenameSliceEmpty(o.FilenameOptions.Filenames,o.FilenameOptions.Kustomize) {
				ioStreams.ErrOut.Write([]byte("Error: must specify one of -f and -k\\n\\n"))
			}
			cmdutil.CheckErr(o.Complete(f,cmd))
			cmdutil.CheckErr(o.ValidateArgs(cmd,args))
			cmdutil.CheckErr(o.RunCreate(f,cmd))
		},
	}

	o.RecordFlags.AddFlags(installCmd)
	usage := "to use to create the resouce"
	cmdutil.AddFilenameOptionFlags(installCmd,&o.FilenameOptions,usage)
	cmdutil.AddValidateFlags(installCmd)
	o.PrintFlags.AddFlags(installCmd)
	cmdutil.AddApplyAnnotationFlags(installCmd)
	cmdutil.AddDryRunFlag(installCmd)
	//cmdutil.AddFieldManagerFlagVar(installCmd,)
	return installCmd
}
