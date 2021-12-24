package operator

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kubectldelete "k8s.io/kubectl/pkg/cmd/delete"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"time"
)

func NewOperatorUninstallCmd(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	deletflags := kubectldelete.NewDeleteFlags("containing the operator to delete.")
	deleteCmd := &cobra.Command{
		Use: "uninstall",
		Short: "Uninstall the milvus operator controller in the cluster",
		Long: "The uninstall subcommand uninstalls the milvus operator controller in the cluster",
		Run: func(cmd *cobra.Command, args []string) {
			o,err := deletflags.ToOptions(nil,ioStreams)
			cmdutil.CheckErr(err)
			cmdutil.CheckErr(o.Complete(f,args,cmd))
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.RunDelete(f))
		},
	}
	deletflags.AddFlags(deleteCmd)
	cmdutil.AddDryRunFlag(deleteCmd)
	return deleteCmd
}

// NewDeleteCommandFlags provides default flags and values for use with the "delete" command
func NewDeleteCommandFlags(usage string) *kubectldelete.DeleteFlags {
	cascadingStrategy := "background"
	gracePeriod := -1

	// setup command defaults
	all := false
	allNamespaces := false
	force := false
	ignoreNotFound := false
	now := false
	output := ""
	labelSelector := ""
	fieldSelector := ""
	timeout := time.Duration(0)
	wait := true
	raw := ""

	filenames := []string{}
	recursive := false
	kustomize := ""

	return &kubectldelete.DeleteFlags{
		// Not using helpers.go since it provides function to add '-k' for FileNameOptions, but not FileNameFlags
		FileNameFlags: &genericclioptions.FileNameFlags{Usage: usage, Filenames: &filenames, Kustomize: &kustomize, Recursive: &recursive},
		LabelSelector: &labelSelector,
		FieldSelector: &fieldSelector,

		CascadingStrategy: &cascadingStrategy,
		GracePeriod:       &gracePeriod,

		All:            &all,
		AllNamespaces:  &allNamespaces,
		Force:          &force,
		IgnoreNotFound: &ignoreNotFound,
		Now:            &now,
		Timeout:        &timeout,
		Wait:           &wait,
		Output:         &output,
		Raw:            &raw,
	}
}