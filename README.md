# kubectl-multi

A comprehensive kubectl plugin for multi-cluster operations with KubeStellar. This plugin extends kubectl to work seamlessly across all KubeStellar managed clusters, providing unified views and operations while filtering out workflow staging clusters (WDS).

## Overview

kubectl-multi is a kubectl plugin written in Go that automatically discovers KubeStellar managed clusters and executes kubectl commands across all of them simultaneously. It provides a unified tabular output with cluster context information, making it easy to monitor and manage resources across multiple clusters.

### Key Features

- **Multi-cluster resource viewing**: Get resources from all managed clusters with unified output
- **Cluster context identification**: Each resource shows which cluster it belongs to
- **All kubectl commands**: Supports all major kubectl commands across clusters
- **KubeStellar integration**: Automatically discovers managed clusters via KubeStellar APIs
- **WDS filtering**: Automatically excludes Workload Description Space clusters
- **Familiar syntax**: Uses the same command structure as kubectl


## how to install 

### Downloading step for Linux
```bash
# Step 1: Download kubectl-multi binary for Linux
TAG="v0.0.2"

# Fix: Use ${TAG#v} to remove just 'v', not 'v_'
curl -L -o "kubectl-multi_${TAG#v}_linux_amd64.tar.gz" \
  "https://github.com/kubestellar/kubectl-plugin/releases/download/${TAG}/kubectl-multi_${TAG#v}_linux_amd64.tar.gz"

# Step 2: Extract and install
tar -xzf "kubectl-multi_${TAG#v}_linux_amd64.tar.gz"
sudo mv kubectl-multi /usr/local/bin/kubectl-multi



#to test 
kubectl-multi help

```
### Downloading step for windows (commands for powershell)
```bash
# Step 1: Download kubectl-multi binary for windows

# TAG="v0.0.2"

curl.exe -LO "https://github.com/kubestellar/kubectl-plugin/releases/download/v0.0.2/kubectl-multi_0.0.2_windows_amd64.zip"

# Step 2: Extract and install
Expand-Archive .\kubectl-multi_0.0.2_windows_amd64.zip

# Step 3: making a new directory for plugin 
New-Item -ItemType Directory -Force -Path C:\kubectl-plugins

# Step 4 : navigate to the directory where kubectl-plugin installed - Downloads
Move-Item .\kubectl-multi.exe C:\kubectl-plugins\kubectl-multi.exe

# Step 5:  Add the Folder to Your System PATH
Go to Control Panel → System and Security → System → Advanced system settings → Environment Variables.  

In “System variables”, select Path, click “Edit”, then “New” and enter C:\kubectl-plugins

#to test (if this command don't work at first then try restarting the powershell terminal )
kubectl-multi.exe --help
or 
kubectl plugin list

```
### Downloading steps for windows (git bash)
```bash
# Step 1: Download kubectl-multi binary for Windows
# TAG="v0.0.2"
curl -LO "https://github.com/kubestellar/kubectl-plugin/releases/download/v0.0.2/kubectl-multi_0.0.2_windows_amd64.zip"

# Step 2: Extract the ZIP file
unzip kubectl-multi_0.0.2_windows_amd64.zip

# Step 3: Create a new directory for plugins
mkdir -p /c/kubectl-plugins

# Step 4: Move the extracted binary to the plugins directory
mv kubectl-multi.exe /c/kubectl-plugins/

# Step 5: Add the plugins directory to your PATH (session only)
export PATH=$PATH:/c/kubectl-plugins

# To make PATH permanent, add the above export line to your ~/.bashrc and restart Git Bash

# Step 6: Test the plugin
kubectl-multi.exe --help
# or
kubectl plugin list
```

## Quick Start for developer

```bash
# Install the plugin
make install

# Get nodes from all managed clusters
kubectl multi get nodes

# Get pods from all clusters in all namespaces
kubectl multi get pods -A
```


## Documentation

- **[Installation Guide](docs/installation_guide.md)** - How to install and set up kubectl-multi
- **[Usage Guide](docs/usage_guide.md)** - Detailed usage examples and commands
- **[Architecture Guide](docs/architecture_guide.md)** - Technical architecture and how it works
- **[Development Guide](docs/development_guide.md)** - Contributing and development workflow
- **[API Reference](docs/api_reference.md)** - Code organization and technical implementation

## Tech Stack

- **Go 1.21+**: Primary language for the plugin
- **Cobra**: CLI framework for command structure and parsing
- **Kubernetes client-go**: Official Kubernetes Go client library
- **KubeStellar APIs**: For managed cluster discovery

## Example Output

```
CONTEXT  CLUSTER       NAME                    STATUS  ROLES          AGE    VERSION
its1     cluster1      cluster1-control-plane  Ready   control-plane  6d23h  v1.33.1
its1     cluster2      cluster2-control-plane  Ready   control-plane  6d23h  v1.33.1
its1     its1-cluster  kubeflex-control-plane  Ready   <none>         6d23h  v1.27.2+k3s1
```

## Related Projects

- [KubeStellar](https://github.com/kubestellar/kubestellar) - Multi-cluster configuration management
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) - Kubernetes command-line tool

## Support

For issues and questions:
- File an issue in this repository  
- Check the KubeStellar documentation
- Join the KubeStellar community discussions

## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.