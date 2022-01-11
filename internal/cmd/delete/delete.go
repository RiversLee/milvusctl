package delete

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kubectldelete "k8s.io/kubectl/pkg/cmd/delete"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)
type MilvusDeleteOptions struct {
	WithDeletions bool
	Namespace string
	Deleteflags *kubectldelete.DeleteFlags
	DeletionOptions *kubectldelete.DeleteOptions
}
func NewMivlusDeleteOptions(ioStreams genericclioptions.IOStreams) *MilvusDeleteOptions {
	deletflags := kubectldelete.NewDeleteFlags("containing the milvus  to delete.")
	o,_ := deletflags.ToOptions(nil,ioStreams)
	return &MilvusDeleteOptions{
		Deleteflags: deletflags,
		WithDeletions: false,
		Namespace: "default",
		DeletionOptions: o,
	}
}
func NewMilvusDeleteCmd(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := NewMivlusDeleteOptions(ioStreams)
	deleteCmd := &cobra.Command{
		Use: "delete",
		Short: "delete milvus in kubernetes cluster",
		Long: "The deelte subcommand uninstalls the milvus version like standalone or cluster in the cluster",
		PreRun: func(cmd *cobra.Command, args []string) {
			
		},
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}
	o.Deleteflags.AddFlags(deleteCmd)
	cmdutil.AddDryRunFlag(deleteCmd)
	deleteCmd.Flags().BoolVar(&o.WithDeletions,"withe-deletion",o.WithDeletions,"automatically add pvc deletion parameter on deletion")
	return deleteCmd
}
func (o *MilvusDeleteOptions) Run() error {
	if o.WithDeletions == false {

	}else {

	}
	return nil
}