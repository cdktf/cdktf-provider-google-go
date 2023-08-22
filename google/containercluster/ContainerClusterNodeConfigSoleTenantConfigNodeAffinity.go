package containercluster


type ContainerClusterNodeConfigSoleTenantConfigNodeAffinity struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#key ContainerCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#operator ContainerCluster#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#values ContainerCluster#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

