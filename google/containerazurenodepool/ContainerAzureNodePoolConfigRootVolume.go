package containerazurenodepool


type ContainerAzureNodePoolConfigRootVolume struct {
	// Optional.
	//
	// The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_node_pool#size_gib ContainerAzureNodePool#size_gib}
	SizeGib *float64 `field:"optional" json:"sizeGib" yaml:"sizeGib"`
}

