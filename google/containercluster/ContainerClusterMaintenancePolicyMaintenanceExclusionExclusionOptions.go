package containercluster


type ContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions struct {
	// The scope of automatic upgrades to restrict in the exclusion window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#scope ContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

