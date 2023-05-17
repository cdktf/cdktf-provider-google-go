package cloudrundomainmapping


type CloudRunDomainMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_domain_mapping#create CloudRunDomainMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_domain_mapping#delete CloudRunDomainMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

