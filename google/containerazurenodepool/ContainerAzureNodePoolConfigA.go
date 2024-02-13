// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerazurenodepool


type ContainerAzureNodePoolConfigA struct {
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/container_azure_node_pool#ssh_config ContainerAzureNodePool#ssh_config}
	SshConfig *ContainerAzureNodePoolConfigSshConfig `field:"required" json:"sshConfig" yaml:"sshConfig"`
	// Optional.
	//
	// The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/container_azure_node_pool#labels ContainerAzureNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/container_azure_node_pool#proxy_config ContainerAzureNodePool#proxy_config}
	ProxyConfig *ContainerAzureNodePoolConfigProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/container_azure_node_pool#root_volume ContainerAzureNodePool#root_volume}
	RootVolume *ContainerAzureNodePoolConfigRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional.
	//
	// A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/container_azure_node_pool#tags ContainerAzureNodePool#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to `Standard_DS2_v2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/container_azure_node_pool#vm_size ContainerAzureNodePool#vm_size}
	VmSize *string `field:"optional" json:"vmSize" yaml:"vmSize"`
}

