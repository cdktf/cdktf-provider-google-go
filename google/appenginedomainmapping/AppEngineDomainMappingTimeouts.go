package appenginedomainmapping


type AppEngineDomainMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_domain_mapping#create AppEngineDomainMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_domain_mapping#delete AppEngineDomainMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_domain_mapping#update AppEngineDomainMapping#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

