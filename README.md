# Intro
This is a CLI tool to add new config files to your main kubeconfig file at
`~/.kube/config`. 

# Usage 

## Add new config
```
kubeconfig add -f ./newconfig -c mycluster 
```
This command adds the kubeconfig in `./newconfig` path to 
`~/.kube/config`. Right now it supports only one cluster 
in the `newconfig`. The resulting cluster and context name 
in `~/.kube/config` will be `mycluster` and the user will be
`mycluster-admin`.

## Delete config
```
kubeconfig del -c mycluster 
```
This command deletes cluster `mycluster` from `~/.kube/config`. It also backs up `~/.kube/config` to 
`~/.cache/kubeconfig/` directory.

## Update config 
You can update existing config using a combination of `add` and `delete` commands.

# Installation
Download it from the release page of this repo.
After downloading, move the binary to a directory in your PATH.
