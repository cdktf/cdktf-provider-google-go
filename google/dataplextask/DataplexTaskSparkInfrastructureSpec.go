package dataplextask


type DataplexTaskSparkInfrastructureSpec struct {
	// batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#batch DataplexTask#batch}
	Batch *DataplexTaskSparkInfrastructureSpecBatch `field:"optional" json:"batch" yaml:"batch"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#container_image DataplexTask#container_image}
	ContainerImage *DataplexTaskSparkInfrastructureSpecContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// vpc_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#vpc_network DataplexTask#vpc_network}
	VpcNetwork *DataplexTaskSparkInfrastructureSpecVpcNetwork `field:"optional" json:"vpcNetwork" yaml:"vpcNetwork"`
}

