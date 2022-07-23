package actions

type Kubeconfig struct {
	ApiVersion string  `yaml:"apiVersion"`
	Clusters []KubeconfigClusterWithName `yaml:"clusters"`
	Contexts []KubeconfigContextWithName`yaml:"contexts"`
	Users []KubeconfigUserWithName 	   `yaml:"users"`
	CurrentContext string `yaml:"current-context"`
	Kind string `yaml:"kind"`
}

// User
type KubeconfigUserWithName struct {
	User map[string]string `yaml:"user"`
	Name string `yaml:"name"`
}

// Context
type KubeconfigContextWithName struct {
	Context map[string]string `yaml:"context"`
	Name string `yaml:"name"`
}

// Cluster
type KubeconfigClusterWithName struct {
	Cluster map[string]string `yaml:"cluster"`
	Name string `yaml:"name"`
}



