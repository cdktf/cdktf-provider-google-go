package projectorganizationpolicy


type ProjectOrganizationPolicyListPolicy struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_organization_policy#allow ProjectOrganizationPolicy#allow}
	Allow *ProjectOrganizationPolicyListPolicyAllow `field:"optional" json:"allow" yaml:"allow"`
	// deny block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_organization_policy#deny ProjectOrganizationPolicy#deny}
	Deny *ProjectOrganizationPolicyListPolicyDeny `field:"optional" json:"deny" yaml:"deny"`
	// If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_organization_policy#inherit_from_parent ProjectOrganizationPolicy#inherit_from_parent}
	InheritFromParent interface{} `field:"optional" json:"inheritFromParent" yaml:"inheritFromParent"`
	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_organization_policy#suggested_value ProjectOrganizationPolicy#suggested_value}
	SuggestedValue *string `field:"optional" json:"suggestedValue" yaml:"suggestedValue"`
}

