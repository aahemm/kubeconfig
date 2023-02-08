/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/user"

	"github.com/aahemm/kubeconfig/pkg/actions"

	"github.com/spf13/cobra"
)

var configFile string 
var clusterName string
var mainKubeconfigFilePath string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a kubeconfig to your ~/.kube/config file",
	Long: `Add a kubeconfig to your ~/.kube/config file. The new config
	file must only have one cluster, user and context`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if clusterName == "" {
			return fmt.Errorf("provide a cluster name with -c or --cluster")
		}

		if mainKubeconfigFilePath == "" {
			currentUser, _ := user.Current()
			mainKubeconfigFilePath = currentUser.HomeDir + "/.kube/config"
		}

		err := actions.AddConfig(clusterName, mainKubeconfigFilePath, configFile)
		if err != nil {
			return fmt.Errorf("there was an error while adding config file: %w", err)
		}
		return nil
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

	addCmd.Flags().StringVarP(
		&mainKubeconfigFilePath, 
		"mainconfig", 
		"m", 
		"",
		"Path to main kubeconfig file",
	)
}
