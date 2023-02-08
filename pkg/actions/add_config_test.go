package actions

import (
	"testing"
)

func Test_readKubeconfigFile_normal(t *testing.T) {
	type kubeconfigTestCase struct {
		input  string
		config Kubeconfig
	}

	testTable := []kubeconfigTestCase{
		{
			input: "./samples/test-config.yaml",
			config: Kubeconfig{
				ApiVersion:     "v1",
				CurrentContext: "cl-test",
				Kind:           "Config",
				Clusters: []KubeconfigClusterWithName{
					{
						Cluster: map[string]interface{}{
							"certificate-authority-data": "test-ca-RUsrMVdOU",
							"server":                     "https://192.168.0.6:6443",
						},
						Name: "cl-test",
					},
				},
				Users: []KubeconfigUserWithName{
					{
						User: map[string]string{
							"client-certificate-data": "test-cert-JTiBDRVJU",
							"client-key-data":         "test-key-EgUFJJVkFURSBLRVkt",
						},
						Name: "cl-test-admin",
					},
				},
				Contexts: []KubeconfigContextWithName{
					{
						Context: map[string]string{
							"cluster":   "cl-test",
							"namespace": "monitor",
							"user":      "cl-test-admin",
						},
						Name: "cl-test",
					},
				},
			},
		},
	}

	for _, tcase := range testTable {
		kcon, err := readKubeconfigFile(tcase.input)
		if err != nil {
			t.Fatalf("Reading %s by readKubeconfigFile returned error that was not expected", tcase.input)
		}
		if kcon.ApiVersion != tcase.config.ApiVersion {
			t.Fatalf("readKubeconfigFile apiVersion error: %s != %s", kcon.ApiVersion, tcase.config.ApiVersion)
		}
		if kcon.CurrentContext != tcase.config.CurrentContext {
			t.Fatalf("readKubeconfigFile currentContext error: %s != %s", kcon.CurrentContext, tcase.config.CurrentContext)
		}
		if kcon.Kind != tcase.config.Kind {
			t.Fatalf("readKubeconfigFile kind error: %s != %s", kcon.Kind, tcase.config.Kind)
		}
		if kcon.Contexts[0].Name != tcase.config.Contexts[0].Name {
			t.Fatalf("readKubeconfigFile context name error: %s != %s", kcon.Contexts[0], tcase.config.Contexts[0].Name)
		}
		if kcon.Clusters[0].Cluster["certificate-authority-data"] != tcase.config.Clusters[0].Cluster["certificate-authority-data"] {
			t.Fatalf(
				"readKubeconfigFile cluster ca error: %s != %s", 
				kcon.Clusters[0].Cluster["certificate-authority-data"], 
				tcase.config.Clusters[0].Cluster["certificate-authority-data"],
			)
		}
	}
}

func Test_readKubeconfigFile_insecure(t *testing.T) {
	type kubeconfigTestCase struct {
		input  string
		config Kubeconfig
	}

	testTable := []kubeconfigTestCase{
		{
			input: "./samples/insecure-config.yaml",
			config: Kubeconfig{
				ApiVersion:     "v1",
				CurrentContext: "cl-insecure",
				Kind:           "Config",
				Clusters: []KubeconfigClusterWithName{
					{
						Cluster: map[string]interface{}{
							"insecure-skip-tls-verify": true,
							"server":                   "https://192.168.0.7:6443",
						},
						Name: "cl-insecure",
					},
				},
				Users: []KubeconfigUserWithName{
					{
						User: map[string]string{
							"client-certificate-data": "insc-cert-1CRUdJT",
							"client-key-data":         "insc-key-EgUFtLQpN",
						},
						Name: "cl-insecure-admin",
					},
				},
				Contexts: []KubeconfigContextWithName{
					{
						Context: map[string]string{
							"cluster":   "cl-insecure",
							"namespace": "platform",
							"user":      "cl-inscure-admin",
						},
						Name: "cl-insecure",
					},
				},
			},
		},
	}

	for _, tcase := range testTable {
		kcon, err := readKubeconfigFile(tcase.input)
		if err != nil {
			t.Fatalf("Reading %s by readKubeconfigFile returned error that was not expected", tcase.input)
		}
		if kcon.ApiVersion != tcase.config.ApiVersion {
			t.Fatalf("readKubeconfigFile apiVersion error: %s != %s", kcon.ApiVersion, tcase.config.ApiVersion)
		}
		if kcon.CurrentContext != tcase.config.CurrentContext {
			t.Fatalf("readKubeconfigFile currentContext error: %s != %s", kcon.CurrentContext, tcase.config.CurrentContext)
		}
		if kcon.Kind != tcase.config.Kind {
			t.Fatalf("readKubeconfigFile kind error: %s != %s", kcon.Kind, tcase.config.Kind)
		}
		if kcon.Contexts[0].Name != tcase.config.Contexts[0].Name {
			t.Fatalf("readKubeconfigFile context name error: %s != %s", kcon.Contexts[0], tcase.config.Contexts[0].Name)
		}
		if _, ok := kcon.Clusters[0].Cluster["insecure-skip-tls-verify"].(bool); !ok {
			t.Fatalf("readKubeconfigFile context insecure-skip-tls-verify is not boolean: %v", kcon.Clusters[0].Cluster["insecure-skip-tls-verify"])
		}
	}
}
