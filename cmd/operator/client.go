package operator

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"github.com/go-logr/logr"
)

//BuildClientConfig builds a client set from a kubeconfig filepath and context
func BuildClientConfig(kubeconfig,context string) (*rest.Config,error) {
	config, err := BuildClientCmd(kubeconfig,context).ClientConfig()
	if err != nil {
		logr.Logger{}.Error(err,"can't find the kubeconfig in computer")
		return nil, err
	}
	return config,nil
}

//BuildClientCmd builds a client cmd config from a kubeconfig filepath and context
func BuildClientCmd(kubeconfig,context string ,overrides ...func(configOverrides *clientcmd.ConfigOverrides)) clientcmd.ClientConfig {
	if kubeconfig !=  "" {
		info, err := os.Stat(kubeconfig)
		if err != nil || info.Size() == 0 {
			kubeconfig = ""
		}

	}

	/**
		config loading rules:
		1. kubeconfig if it not empty string
		2. Config(s) in KUBECONFIG environment variable
		3. In cluster config if running in-cluster
		4. Use $HOME/.kube/config
	 */
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	loadingRules.ExplicitPath = kubeconfig
	configOverrides := &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmd.ClusterDefaults,
		CurrentContext: context,
	}
	for _,fn := range overrides {
		fn(configOverrides)
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules,configOverrides)
}

//GetClientSet use rest.Config get the kubernetes clientset.
func GetClientSet(config *rest.Config) (*kubernetes.Clientset,error){
	clientSet,err := kubernetes.NewForConfig(config)
	if err != nil {
		logr.Logger{}.Error(err,"can't create the kubernetes clientset")
		return nil,err
	}
	return clientSet,nil
}