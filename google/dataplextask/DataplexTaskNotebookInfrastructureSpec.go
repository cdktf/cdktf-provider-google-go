package dataplextask


type DataplexTaskNotebookInfrastructureSpec struct {
	// batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#batch DataplexTask#batch}
	Batch *DataplexTaskNotebookInfrastructureSpecBatch `field:"optional" json:"batch" yaml:"batch"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#container_image DataplexTask#container_image}
	ContainerImage *DataplexTaskNotebookInfrastructureSpecContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// vpc_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task#vpc_network DataplexTask#vpc_network}
	VpcNetwork *DataplexTaskNotebookInfrastructureSpecVpcNetwork `field:"optional" json:"vpcNetwork" yaml:"vpcNetwork"`
}

