package accesscontextmanageregresspolicy


type AccessContextManagerEgressPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_egress_policy#create AccessContextManagerEgressPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_egress_policy#delete AccessContextManagerEgressPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

