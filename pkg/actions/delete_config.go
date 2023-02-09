package actions

import "fmt"

type KubeconfigComponent interface {
	KubeconfigClusterWithName | KubeconfigContextWithName | KubeconfigUserWithName
}

func DeleteConfig(clusterName, mainKubeconfigFilePath string) error {
	err := backup(mainKubeconfigFilePath, "mainconfig-bkp")
	if err != nil {
		return err
	}
	mainKubeconfig, err := readKubeconfigFile(mainKubeconfigFilePath)
	if err != nil {
		return err
	}

	for i, cl := range mainKubeconfig.Clusters {
		if cl.Name == clusterName {
			sl, err := removeSliceElement(mainKubeconfig.Clusters, i)
			if err != nil {
				return fmt.Errorf(
					"could not remove cluster %s from %s: %w", 
					clusterName, mainKubeconfigFilePath, err,
				)
			}
			mainKubeconfig.Clusters = sl 
			fmt.Printf("Deleted cluster %s from clusters \n", clusterName)
		}
	}

	for j, ctx := range mainKubeconfig.Contexts {
		if ctx.Name == clusterName {
			sl, err := removeSliceElement(mainKubeconfig.Contexts, j)
			if err != nil {
				return fmt.Errorf(
					"could not remove context %s from %s: %w", 
					clusterName, mainKubeconfigFilePath, err,
				)
			}
			mainKubeconfig.Contexts = sl 
			fmt.Printf("Deleted context %s from contexts \n", clusterName)
		}
	}

	for k, u := range mainKubeconfig.Users {
		if u.Name == clusterName+"-admin" {
			sl, err := removeSliceElement(mainKubeconfig.Users, k)
			if err != nil {
				return fmt.Errorf(
					"could not remove user %s-admin from %s: %w", 
					clusterName, mainKubeconfigFilePath, err,
				)
			}
			mainKubeconfig.Users = sl
			fmt.Printf("Deleted user %s-admin from users \n", clusterName)
		}
	}
	return writeKubeconfigFile(mainKubeconfig, mainKubeconfigFilePath)

}

func removeSliceElement[S KubeconfigComponent](sl []S, i int) ([]S, error) {
	slLen := len(sl)
	if slLen < i+1 {
		return sl, fmt.Errorf("index is out of range")
	} 

	sl[i] = sl[slLen-1] 
	return sl[:slLen-1], nil 
}