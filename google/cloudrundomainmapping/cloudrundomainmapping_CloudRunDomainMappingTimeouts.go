package cloudrundomainmapping


type CloudRunDomainMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_domain_mapping#create CloudRunDomainMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_domain_mapping#delete CloudRunDomainMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

