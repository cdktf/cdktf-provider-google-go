package containerazurecluster


type ContainerAzureClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_azure_cluster#project ContainerAzureCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

