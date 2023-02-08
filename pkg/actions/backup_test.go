package actions

import (
	"os"
	"os/user"
	"testing"
)

func Test_backup(t *testing.T){
	backup("./samples/wrong-config.yaml", "wrong-config-bkp")
	mainData, _ := os.ReadFile("./samples/wrong-config.yaml")

	currentUser, _ := user.Current()
	backupDir := currentUser.HomeDir + "/.cache/kubeconfig/"
	bkpData, _ := os.ReadFile(backupDir+"wrong-config-bkp")

	if string(bkpData) != string(mainData) {
		t.Fatalf("backup function did not work correct. %s != %s", string(bkpData), string(mainData))
	}

	os.Remove(backupDir+"wrong-config-bkp")
}