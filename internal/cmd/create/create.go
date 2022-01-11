package create

import (
	"context"
	"fmt"
	"helm.sh/helm/v3/pkg/strvals"
	"github.com/milvus-io/milvus-operator/apis/milvus.io/v1alpha1"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
	pkgerr "github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kubectlcreate "k8s.io/kubectl/pkg/cmd/create"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/templates"
	"k8s.io/kubectl/pkg/util/i18n"
	"log"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/fatih/structs"
	"github.com/goinggo/mapstructure"
)
var (
	createLong = templates.LongDesc(i18n.T(`
		The create subcommand installs the milvus version like standalone or cluster in the cluster
    `))
)
type printFn func(format string, v ...interface{})
type MilvusCreateOptions struct {
	Mode string
	Type string
	Values []string
	CreateOptions *kubectlcreate.CreateOptions
	ResouceSetting map[string]interface{}
}
func NewMivlusCreateOptions(ioStreams genericclioptions.IOStreams) *MilvusCreateOptions {
	return &MilvusCreateOptions{
		Type: "",
		Mode: "",
		CreateOptions: kubectlcreate.NewCreateOptions(ioStreams),
	}
}
func NewMilvusCreateCmd(f cmdutil.Factory, ioStreams genericclioptions.IOStreams,client *client.Client) *cobra.Command {
	o := NewMivlusCreateOptions(ioStreams)
	createCmd := &cobra.Command{
		Use: "create {-f filename | -t type -m model}",
		Short: "create milvuse in kubernetes cluster",
		Long: createLong,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(o.CreateOptions.FilenameOptions.Filenames) > 0 && (o.Mode != "" || o.Type != ""){
				ioStreams.ErrOut.Write([]byte("Error:-f conflict with other flag, if you want to specify filename,it can't set another flag"))
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			//if cmdutil.IsFilenameSliceEmpty(o.CreateOptions.FilenameOptions.Filenames,o.CreateOptions.FilenameOptions.Kustomize) {
			//	ioStreams.ErrOut.Write([]byte("Error: must specify one of -f and -k\\n\\n"))
			//}.
			cmdutil.CheckErr(o.Complete(f, cmd))
			cmdutil.CheckErr(o.ValidateArgs(cmd, args))
			//cmdutil.CheckErr(o.Run(f, cmd,client))
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
	createCmd.Flags().StringArrayVar(&o.Values,"set",[]string{},"the resource requirement requests for milvus cluster")
	//_ = createCmd.MarkFlagRequired("mode")
	
	return createCmd
}

func (o *MilvusCreateOptions) Complete(f cmdutil.Factory, cmd *cobra.Command) error {
	var err error
	err = o.CreateOptions.Complete(f,cmd)
	if err != nil {
		return err
	}
	return err
}
func (o *MilvusCreateOptions) ValidateArgs(cmd *cobra.Command,args []string) error{
	var err error
	err = o.CreateOptions.ValidateArgs(cmd,args)
	if err != nil {
		return err
	}
	spec := new(v1alpha1.MilvusSpec)
	fmt.Println(spec)
	if len(o.Values)  > 0 {
		base := map[string]interface{}{}
		for _,value := range o.Values {
			if err := strvals.ParseInto(value, base); err != nil {
				return  pkgerr.Wrap(err, "failed parsing --set data")
			}
		}
		o.ResouceSetting = base
	}
	return err
}
func (o *MilvusCreateOptions) Run(f cmdutil.Factory, cmd *cobra.Command,client *client.Client) error {
	var err error
	if len(o.CreateOptions.FilenameOptions.Filenames) > 0 {
		err = o.CreateOptions.RunCreate(f,cmd)
		if err != nil {
			return err
		}
	}
	if o.Mode == "cluster" {
		

	}
	if o.Mode == "standalone" {
		if _,err = newMilvusStandalone(*client,context.TODO(),o.Type);err != nil {
			return err
		}
	}
	return err
}
//func runSetMilvusConfig(o *MilvusCreateOptions, milvus *v1alpha1.Milvus) error {
//	if len(o.Values) == 0 {
//		return nil
//	}
//	for key, vaylue :=  range o.ResouceSetting {
//		tags := strings.Split(key,".")
//	}
//	milvus.Spec.
//	return nil
//}
func runSetMilvusClusterConfig(setting string, milvusCluster *v1alpha1.MilvusCluster) error {
	if len(setting) == 0 {
		return nil
	}
	return nil
}
func newMilvusStandalone(client client.Client,ctx context.Context,Type string) (*v1alpha1.Milvus,error) {
	log.Printf("Creating the milvus in default namespace")
	switch Type {
	case "minimal":
	case "medium":
	case "large":
		
	}
	namespacedName := types.NamespacedName{
		Name: "milvus",
		Namespace: "default",
	}
	milvus := &v1alpha1.Milvus{}
	err := client.Get(ctx,namespacedName,milvus)
	if errors.IsNotFound(err) {
		milvus = &v1alpha1.Milvus{
			ObjectMeta:metav1.ObjectMeta{
				Name:"milvus",
				Namespace: "default",
			},
		}
		//milvus.Spec = v1alpha1.MilvusSpec{
		//	Dep: v1alpha1.MilvusDependencies{
		//
		//	},
		//	ComponentSpec:v1alpha1.ComponentSpec{
		//
		//	},
		//	ServiceType: "",
		//	Conf: v1alpha1.Values{
		//
		//	},
		//}
	}
	err = client.Create(ctx,milvus)
	if err != nil{
		return nil ,err
	}
	return milvus,nil
}
func newMilvusCluster(client client.Client,ctx context.Context,Type string) (*v1alpha1.MilvusCluster,error){
	log.Printf("Creating the milvus cluster in default namespace")
	switch Type {
	case "minimal":
	case "medium":
	case "large":
	}
	namespacedName := types.NamespacedName{
		Name: "milvus-cluster",
		Namespace: "default",
	}
	milvusCluster := &v1alpha1.MilvusCluster{}
	err := client.Get(ctx,namespacedName,milvusCluster)
	if errors.IsNotFound(err) {
		milvusCluster = &v1alpha1.MilvusCluster{
			ObjectMeta:metav1.ObjectMeta{
				Name:"milvus",
				Namespace: "default",
			},
		}
		milvusCluster.Spec = v1alpha1.MilvusClusterSpec{
			Dep: v1alpha1.MilvusClusterDependencies{

			},
			Com:v1alpha1.MilvusComponents{

			},
			Conf: v1alpha1.Values{

			},
		}
	}
	err = client.Create(ctx,milvusCluster)
	if err != nil {
		return nil,err
	}
	return milvusCluster,nil
}

// coalesceValues builds up a values map for a particular CRD.
//
// Values in v will override the src values in the crd.
func coalesceValues(printf printFn,src,values map[string]interface{}) map[string]interface{} {
	if values == nil {
		return src
	}
	for key,val := range src {
		if value,ok := values[key]; ok {
			if value == nil {
				delete(values,key)
			}else if  dest,ok := value.(map[string]interface{});ok {
				src,ok := val.(map[string]interface{})
				if !ok {
					if val != nil {
						printf("warning:skipped value for %s.%s: Not a table.")
					}
				}else {
					coalesceTablesFullKey(printf, dest, src)

				}
			}
		}else{
			values[key] = val
		}

	}
	return values
}
// coalesceTablesFullKey merges a source map into a destination map.
//
// dest is considered authoritative.
func coalesceTablesFullKey(printf printFn, dst, src map[string]interface{}) map[string]interface{} {
	// When --reuse-values is set but there are no modifications yet, return new values
	if src == nil {
		return dst
	}
	if dst == nil {
		return src
	}
	// Because dest has higher precedence than src, dest values override src
	// values.
	for key, val := range src {
		if dv, ok := dst[key]; ok && dv == nil {
			delete(dst, key)
		} else if !ok {
			dst[key] = val
		} else if istable(val) {
			if istable(dv) {
				coalesceTablesFullKey(printf, dv.(map[string]interface{}), val.(map[string]interface{}))
			} else {
				printf("warning: cannot overwrite table with non table for %s (%v)", val)
			}
		} else if istable(dv) && val != nil {
			printf("warning: destination for %s is a table. Ignoring non-table value (%v)", val)
		}
	}
	return dst
}

// istable is a special-purpose function to see if the present thing matches the definition of a YAML table.
func istable(v interface{}) bool {
	_, ok := v.(map[string]interface{})
	return ok
}
func  milvusStruct2Map(milvusSpec *v1alpha1.MilvusSpec) map[string]interface{} {
	return structs.Map(milvusSpec)
}
func milvusMap2Struect(milvusMap map[string]interface{}) *v1alpha1.MilvusSpec {
	var milvus v1alpha1.MilvusSpec
	mapstructure.Decode(milvusMap, &milvus)
	return &milvus
}
func milvusClusterStruct2Map(milvusClusterSpec *v1alpha1.MilvusClusterSpec) map[string]interface{} {
	return structs.Map(milvusClusterSpec)
}
func milvusClusterMap2Struct(milvusClusterMap map[string]interface{}) *v1alpha1.MilvusClusterSpec {
	var milvusCluster v1alpha1.MilvusClusterSpec
	mapstructure.Decode(milvusClusterMap, &milvusCluster)
	return &milvusCluster
}
