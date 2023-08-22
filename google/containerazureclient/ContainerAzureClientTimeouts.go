package containerazureclient


type ContainerAzureClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_azure_client#create ContainerAzureClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_azure_client#delete ContainerAzureClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

