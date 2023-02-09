package actions

import (
	"fmt"
	"testing"
)

func Test_DeleteConfig(t *testing.T) {
	copyFile("./samples/main-test-merged-config.yaml", "./samples/tmp-delete-config.yaml")
	DeleteConfig("newcl", "./samples/tmp-delete-config.yaml")

	if !areFilesEqual("./samples/tmp-delete-config.yaml", "./samples/delete-config.yaml") {
		t.Fatalf("could not delete newcl from ./samples/main-test-merged-config.yaml")
	}
}

func Test_removeSliceElement(t *testing.T){
	sl := []KubeconfigUserWithName{
		{
			User: map[string]string{"ctx": "newctx"},
			Name: "one",
		},
		{
			User: map[string]string{"ctx": "newerctx"},
			Name: "two",
		},
	}
	new, _ := removeSliceElement(sl, 0)
	fmt.Printf("%v \n", new)
}
