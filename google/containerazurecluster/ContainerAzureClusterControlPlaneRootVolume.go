package containerazurecluster


type ContainerAzureClusterControlPlaneRootVolume struct {
	// Optional.
	//
	// The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.0/docs/resources/container_azure_cluster#size_gib ContainerAzureCluster#size_gib}
	SizeGib *float64 `field:"optional" json:"sizeGib" yaml:"sizeGib"`
}

