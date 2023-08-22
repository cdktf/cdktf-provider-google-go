package orgpolicypolicy


type OrgPolicyPolicySpecRulesValues struct {
	// List of values allowed at this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/org_policy_policy#allowed_values OrgPolicyPolicy#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// List of values denied at this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/org_policy_policy#denied_values OrgPolicyPolicy#denied_values}
	DeniedValues *[]*string `field:"optional" json:"deniedValues" yaml:"deniedValues"`
}

