package actions

import (
	"fmt"
	"io"
	"os"
	"os/user"
)

func backup(configPath, backupName string) error {
	currentUser, _ := user.Current()
	backupDir := currentUser.HomeDir + "/.cache/kubeconfig/"
	os.Mkdir(backupDir, 0755)
	
    mainKubeconfigFile, err := os.Open(configPath)
    if err != nil {
        return fmt.Errorf("could not open kubeconfig file %s: %w", configPath, err)
    }
    defer mainKubeconfigFile.Close()

    backupFile, err := os.Create(backupDir+backupName)
    if err != nil {
        return fmt.Errorf("could not create kubeconfig backup %s: %w", backupDir+backupName, err)
    }
    defer backupFile.Close()

    byteNum, err := io.Copy(backupFile, mainKubeconfigFile)
    if err != nil {
        return fmt.Errorf("could not write kubeconfig backup %s: %w", backupDir+backupName, err)
    }
	fmt.Printf("Backed up %d bytes to %s", byteNum, backupDir+backupName)
	return nil 
}