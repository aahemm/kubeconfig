package actions

import (
	"fmt"
	"io/ioutil"
	"os/user"

	yaml "gopkg.in/yaml.v2"
)

func readKubeconfigFile (configFilePath string) (Kubeconfig, error) {
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		newErr := fmt.Errorf("error reading config file %s: %w", configFilePath, err)
		return Kubeconfig{}, newErr
    }

	var kubeconfig Kubeconfig
	err = yaml.Unmarshal(configFile, &kubeconfig)
	if err != nil {
		newErr := fmt.Errorf("error parsing config file %s: %w", configFilePath, err)
		return Kubeconfig{}, newErr
	}

	return kubeconfig, nil
}

func writeKubeconfigFile (kubeconfig Kubeconfig, configFilePath string) error {
	configFile, err := yaml.Marshal(kubeconfig)
	if err != nil {
		newErr := fmt.Errorf("error serializing config file: %w", err)
		return newErr
	}

	err = ioutil.WriteFile(configFilePath, configFile, 0)
	if err != nil {
		newErr := fmt.Errorf("error writing to config file %s: %w", configFilePath, err)
		return newErr
	}

	return nil 
} 

func AddConfig(clusterName, configFilePath string) error {
	newKubeconfig, err := readKubeconfigFile(configFilePath)
	if err != nil {
		return err 
	}

	currentUser, _ := user.Current()
	mainKubeconfigFilePath := currentUser.HomeDir + "/.kube/config"
	mainKubeconfig, err := readKubeconfigFile(mainKubeconfigFilePath)
	if err != nil {
		return err
	}

	newKubeconfig.Clusters[0].Name = clusterName
	newKubeconfig.Users[0].Name = clusterName + "-admin"
	newKubeconfig.Contexts[0] = KubeconfigContextWithName{
		Context: map[string]string {
			"cluster": clusterName,
			"user": clusterName + "-admin",
		},
		Name: clusterName,
	}  

	mainKubeconfig.Clusters = append(mainKubeconfig.Clusters, newKubeconfig.Clusters...)
	mainKubeconfig.Contexts = append(mainKubeconfig.Contexts, newKubeconfig.Contexts...)
	mainKubeconfig.Users = append(mainKubeconfig.Users, newKubeconfig.Users...)

	// fmt.Printf("%v \n", mainKubeconfig)
	return writeKubeconfigFile(mainKubeconfig, mainKubeconfigFilePath)

}
