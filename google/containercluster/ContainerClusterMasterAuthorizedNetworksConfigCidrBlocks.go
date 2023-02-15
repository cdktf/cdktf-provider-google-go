package containercluster


type ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks struct {
	// External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#cidr_block ContainerCluster#cidr_block}
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Field for users to identify CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#display_name ContainerCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

