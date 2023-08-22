package containernodepool


type ContainerNodePoolNetworkConfig struct {
	// Whether to create a new range for pod IPs in this node pool.
	//
	// Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_node_pool#create_pod_range ContainerNodePool#create_pod_range}
	CreatePodRange interface{} `field:"optional" json:"createPodRange" yaml:"createPodRange"`
	// Whether nodes have internal IP addresses only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_node_pool#enable_private_nodes ContainerNodePool#enable_private_nodes}
	EnablePrivateNodes interface{} `field:"optional" json:"enablePrivateNodes" yaml:"enablePrivateNodes"`
	// pod_cidr_overprovision_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_node_pool#pod_cidr_overprovision_config ContainerNodePool#pod_cidr_overprovision_config}
	PodCidrOverprovisionConfig *ContainerNodePoolNetworkConfigPodCidrOverprovisionConfig `field:"optional" json:"podCidrOverprovisionConfig" yaml:"podCidrOverprovisionConfig"`
	// The IP address range for pod IPs in this node pool.
	//
	// Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_node_pool#pod_ipv4_cidr_block ContainerNodePool#pod_ipv4_cidr_block}
	PodIpv4CidrBlock *string `field:"optional" json:"podIpv4CidrBlock" yaml:"podIpv4CidrBlock"`
	// The ID of the secondary range for pod IPs.
	//
	// If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_node_pool#pod_range ContainerNodePool#pod_range}
	PodRange *string `field:"optional" json:"podRange" yaml:"podRange"`
}

