package organizationpolicy


type OrganizationPolicyListPolicy struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/organization_policy#allow OrganizationPolicy#allow}
	Allow *OrganizationPolicyListPolicyAllow `field:"optional" json:"allow" yaml:"allow"`
	// deny block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/organization_policy#deny OrganizationPolicy#deny}
	Deny *OrganizationPolicyListPolicyDeny `field:"optional" json:"deny" yaml:"deny"`
	// If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/organization_policy#inherit_from_parent OrganizationPolicy#inherit_from_parent}
	InheritFromParent interface{} `field:"optional" json:"inheritFromParent" yaml:"inheritFromParent"`
	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/organization_policy#suggested_value OrganizationPolicy#suggested_value}
	SuggestedValue *string `field:"optional" json:"suggestedValue" yaml:"suggestedValue"`
}

