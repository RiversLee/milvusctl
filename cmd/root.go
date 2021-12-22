/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/milvus-io/milvusctl/cmd/operator"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func GetRootCmd(args []string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "milvusctl",
		Short: "milvus application control interface",
		Long: `milvus configuration command line utility for service operators to debug and diagnose their milvus application`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello cobra cli")
		},
	}
	rootCmd.AddCommand(operator.GetOperatorCmd(args))
	return rootCmd
}
