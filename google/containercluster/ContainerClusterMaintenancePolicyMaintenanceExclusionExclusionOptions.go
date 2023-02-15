package containercluster


type ContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions struct {
	// The scope of automatic upgrades to restrict in the exclusion window.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#scope ContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

