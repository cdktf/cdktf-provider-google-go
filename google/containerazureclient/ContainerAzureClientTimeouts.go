package containerazureclient


type ContainerAzureClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#create ContainerAzureClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_azure_client#delete ContainerAzureClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

