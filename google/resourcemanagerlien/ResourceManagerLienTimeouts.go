package resourcemanagerlien


type ResourceManagerLienTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#create ResourceManagerLien#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#delete ResourceManagerLien#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

