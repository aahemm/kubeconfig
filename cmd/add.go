/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aahemm/kubeconfig/pkg/actions"
	
	"github.com/spf13/cobra"
)

var configFile string 
var clusterName string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a kubeconfig to your ~/.kube/config file",
	Long: `Add a kubeconfig to your ~/.kube/config file. The new config
	file must only have one cluster, user and context`,
	Run: func(cmd *cobra.Command, args []string) {
		if clusterName == "" {
			fmt.Println("Please provide a cluster name with -c or --cluster")
			os.Exit(1)
		}

		err := actions.AddConfig(clusterName, configFile)
		if err != nil {
			fmt.Printf("There was an error while adding config file: %v \n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(
		&configFile, 
		"file", 
		"f", 
		"./config",
		"Path to the kubeconfig file to be added",
	)

	addCmd.Flags().StringVarP(
		&clusterName, 
		"cluster", 
		"c", 
		"",
		"Cluster name of kubeconfig file",
	)
}
