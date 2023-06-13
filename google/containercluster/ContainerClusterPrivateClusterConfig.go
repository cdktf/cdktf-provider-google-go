package containercluster


type ContainerClusterPrivateClusterConfig struct {
	// When true, the cluster's private endpoint is used as the cluster endpoint and access through the public endpoint is disabled.
	//
	// When false, either endpoint can be used. This field only applies to private clusters, when enable_private_nodes is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#enable_private_endpoint ContainerCluster#enable_private_endpoint}
	EnablePrivateEndpoint interface{} `field:"optional" json:"enablePrivateEndpoint" yaml:"enablePrivateEndpoint"`
	// Enables the private cluster feature, creating a private endpoint on the cluster.
	//
	// In a private cluster, nodes only have RFC 1918 private addresses and communicate with the master's private endpoint via private networking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#enable_private_nodes ContainerCluster#enable_private_nodes}
	EnablePrivateNodes interface{} `field:"optional" json:"enablePrivateNodes" yaml:"enablePrivateNodes"`
	// master_global_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#master_global_access_config ContainerCluster#master_global_access_config}
	MasterGlobalAccessConfig *ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig `field:"optional" json:"masterGlobalAccessConfig" yaml:"masterGlobalAccessConfig"`
	// The IP range in CIDR notation to use for the hosted master network.
	//
	// This range will be used for assigning private IP addresses to the cluster master(s) and the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network, and it must be a /28 subnet. See Private Cluster Limitations for more details. This field only applies to private clusters, when enable_private_nodes is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#master_ipv4_cidr_block ContainerCluster#master_ipv4_cidr_block}
	MasterIpv4CidrBlock *string `field:"optional" json:"masterIpv4CidrBlock" yaml:"masterIpv4CidrBlock"`
	// Subnetwork in cluster's network where master's endpoint will be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#private_endpoint_subnetwork ContainerCluster#private_endpoint_subnetwork}
	PrivateEndpointSubnetwork *string `field:"optional" json:"privateEndpointSubnetwork" yaml:"privateEndpointSubnetwork"`
}

