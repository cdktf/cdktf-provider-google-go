package folderorganizationpolicy


type FolderOrganizationPolicyListPolicy struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_organization_policy#allow FolderOrganizationPolicy#allow}
	Allow *FolderOrganizationPolicyListPolicyAllow `field:"optional" json:"allow" yaml:"allow"`
	// deny block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_organization_policy#deny FolderOrganizationPolicy#deny}
	Deny *FolderOrganizationPolicyListPolicyDeny `field:"optional" json:"deny" yaml:"deny"`
	// If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_organization_policy#inherit_from_parent FolderOrganizationPolicy#inherit_from_parent}
	InheritFromParent interface{} `field:"optional" json:"inheritFromParent" yaml:"inheritFromParent"`
	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_organization_policy#suggested_value FolderOrganizationPolicy#suggested_value}
	SuggestedValue *string `field:"optional" json:"suggestedValue" yaml:"suggestedValue"`
}

