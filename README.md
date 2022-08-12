# Intro
This is a CLI tool to add new config files to your main kubeconfig file at
`~/.kube/config`. 

# Usage 
```
kubeconfig add -f ./newconfig -c mycluster 
```
This command adds the kubeconfig in `./newconfig` path to 
`~/.kube/config`. Right now it supports only one cluster 
in the `newconfig`.
