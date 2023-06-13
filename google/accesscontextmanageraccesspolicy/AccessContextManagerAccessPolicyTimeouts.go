package accesscontextmanageraccesspolicy


type AccessContextManagerAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_policy#create AccessContextManagerAccessPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_policy#delete AccessContextManagerAccessPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_policy#update AccessContextManagerAccessPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

