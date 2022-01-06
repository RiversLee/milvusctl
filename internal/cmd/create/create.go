package create

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	kubectlcreate "k8s.io/kubectl/pkg/cmd/create"
	"os"
)
type MilvusCreateOptions struct {
	Mode string
	Type string
	Sets string
	CreateOptions *kubectlcreate.CreateOptions
}
func NewMivlusCreateOptions(ioStreams genericclioptions.IOStreams) *MilvusCreateOptions {
	return &MilvusCreateOptions{
		Type: "",
		Mode: "",
		Sets: "",
		CreateOptions: kubectlcreate.NewCreateOptions(ioStreams),
	}
}
func NewMilvusCreateCmd(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	o := NewMivlusCreateOptions(ioStreams)
	createCmd := &cobra.Command{
		Use: "create {-f filename | -t type -m model}",
		Short: "create milvuse in kubernetes cluster",
		Long: "The create subcommand installs the milvus version like standalone or cluster in the cluster",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(o.CreateOptions.FilenameOptions.Filenames) > 0 && (o.Mode != "" || o.Type != ""){
				ioStreams.ErrOut.Write([]byte("Error:-f conflict with other flag, if you want to specify filename,it can't set another flag"))
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if cmdutil.IsFilenameSliceEmpty(o.CreateOptions.FilenameOptions.Filenames,o.CreateOptions.FilenameOptions.Kustomize) {
				ioStreams.ErrOut.Write([]byte("Error: must specify one of -f and -k\\n\\n"))
			}
			cmdutil.CheckErr(o.CreateOptions.Complete(f, cmd))
			cmdutil.CheckErr(o.CreateOptions.ValidateArgs(cmd, args))
			cmdutil.CheckErr(o.CreateOptions.RunCreate(f, cmd))
		},
	}
	o.CreateOptions.RecordFlags.AddFlags(createCmd)

	usage := "to use to create the resouce"
	cmdutil.AddFilenameOptionFlags(createCmd,&o.CreateOptions.FilenameOptions,usage)
	cmdutil.AddValidateFlags(createCmd)
	o.CreateOptions.PrintFlags.AddFlags(createCmd)
	cmdutil.AddApplyAnnotationFlags(createCmd)
	cmdutil.AddDryRunFlag(createCmd)

	createCmd.Flags().StringVarP(&o.Mode,"mode","m",o.Mode,"use mode parameter to choose milvus standalone or cluster")
	createCmd.Flags().StringVarP(&o.Type,"type","t",o.Type,"use type parameter to choose milvus cluster minimal,medium or large")
	createCmd.Flags().StringVar(&o.Sets,"set",o.Sets,"the resource requirement requests for milvus cluster")
	_ = createCmd.MarkFlagRequired("mode")
	
	return createCmd
}