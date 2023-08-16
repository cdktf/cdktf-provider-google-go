package orgpolicypolicy


type OrgPolicyPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/org_policy_policy#create OrgPolicyPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/org_policy_policy#delete OrgPolicyPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/org_policy_policy#update OrgPolicyPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

