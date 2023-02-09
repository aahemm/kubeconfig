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


var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a config from kubeconfig file",
	Long: `Delete a config from kubeconfig file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if clusterName == "" {
			return fmt.Errorf("provide a cluster name with -c or --cluster")
		}

		if mainKubeconfigFilePath == "" {
			currentUser, _ := user.Current()
			mainKubeconfigFilePath = currentUser.HomeDir + "/.kube/config"
		}

		err := actions.DeleteConfig(clusterName, mainKubeconfigFilePath)
		if err != nil {
			return fmt.Errorf("there was an error while adding config file: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	delCmd.Flags().StringVarP(
		&clusterName, 
		"cluster", 
		"c", 
		"",
		"Cluster name of kubeconfig file",
	)

	delCmd.Flags().StringVarP(
		&mainKubeconfigFilePath, 
		"mainconfig", 
		"m", 
		"",
		"Path to main kubeconfig file",
	)
}
