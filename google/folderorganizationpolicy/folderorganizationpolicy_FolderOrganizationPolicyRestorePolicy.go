package folderorganizationpolicy


type FolderOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/folder_organization_policy#default FolderOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}

