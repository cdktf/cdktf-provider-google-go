package containerazurecluster


type ContainerAzureClusterControlPlane struct {
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#ssh_config ContainerAzureCluster#ssh_config}
	SshConfig *ContainerAzureClusterControlPlaneSshConfig `field:"required" json:"sshConfig" yaml:"sshConfig"`
	// The ARM ID of the subnet where the control plane VMs are deployed. Example: `/subscriptions//resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#subnet_id ContainerAzureCluster#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAzureServerConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#version ContainerAzureCluster#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// database_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#database_encryption ContainerAzureCluster#database_encryption}
	DatabaseEncryption *ContainerAzureClusterControlPlaneDatabaseEncryption `field:"optional" json:"databaseEncryption" yaml:"databaseEncryption"`
	// main_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#main_volume ContainerAzureCluster#main_volume}
	MainVolume *ContainerAzureClusterControlPlaneMainVolume `field:"optional" json:"mainVolume" yaml:"mainVolume"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#proxy_config ContainerAzureCluster#proxy_config}
	ProxyConfig *ContainerAzureClusterControlPlaneProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// replica_placements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#replica_placements ContainerAzureCluster#replica_placements}
	ReplicaPlacements interface{} `field:"optional" json:"replicaPlacements" yaml:"replicaPlacements"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#root_volume ContainerAzureCluster#root_volume}
	RootVolume *ContainerAzureClusterControlPlaneRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional. A set of tags to apply to all underlying control plane Azure resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#tags ContainerAzureCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Optional.
	//
	// The Azure VM size name. Example: `Standard_DS2_v2`. For available VM sizes, see https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions. When unspecified, it defaults to `Standard_DS2_v2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_azure_cluster#vm_size ContainerAzureCluster#vm_size}
	VmSize *string `field:"optional" json:"vmSize" yaml:"vmSize"`
}

