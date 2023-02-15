package containercluster


type ContainerClusterDefaultSnatStatus struct {
	// When disabled is set to false, default IP masquerade rules will be applied to the nodes to prevent sNAT on cluster internal traffic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#disabled ContainerCluster#disabled}
	Disabled interface{} `field:"required" json:"disabled" yaml:"disabled"`
}

