/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/milvus-io/milvusctl/internal/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewMilvusCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

