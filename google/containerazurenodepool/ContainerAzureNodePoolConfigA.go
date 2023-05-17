package containerazurenodepool


type ContainerAzureNodePoolConfigA struct {
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_node_pool#ssh_config ContainerAzureNodePool#ssh_config}
	SshConfig *ContainerAzureNodePoolConfigSshConfig `field:"required" json:"sshConfig" yaml:"sshConfig"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_node_pool#proxy_config ContainerAzureNodePool#proxy_config}
	ProxyConfig *ContainerAzureNodePoolConfigProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_node_pool#root_volume ContainerAzureNodePool#root_volume}
	RootVolume *ContainerAzureNodePoolConfigRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional.
	//
	// A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_node_pool#tags ContainerAzureNodePool#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to `Standard_DS2_v2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_node_pool#vm_size ContainerAzureNodePool#vm_size}
	VmSize *string `field:"optional" json:"vmSize" yaml:"vmSize"`
}

