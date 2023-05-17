package folderorganizationpolicy


type FolderOrganizationPolicyListPolicyDeny struct {
	// The policy allows or denies all values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/folder_organization_policy#all FolderOrganizationPolicy#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// The policy can define specific values that are allowed or denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/folder_organization_policy#values FolderOrganizationPolicy#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

