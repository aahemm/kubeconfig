package actions

import (
	"os"
	"os/user"
	"testing"
)

func Test_backup(t *testing.T){
	backup("./samples/wrong-config.yaml", "wrong-config-bkp")

	currentUser, _ := user.Current()
	backupDir := currentUser.HomeDir + "/.cache/kubeconfig/"
	if !areFilesEqual(backupDir+"wrong-config-bkp", "./samples/wrong-config.yaml") {
		t.Fatal("backup function did not work correct.")
	}

	os.Remove(backupDir+"wrong-config-bkp")
}