// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerClusterNodePoolNodeConfigGcfsConfig struct {
	// Whether or not GCFS is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

