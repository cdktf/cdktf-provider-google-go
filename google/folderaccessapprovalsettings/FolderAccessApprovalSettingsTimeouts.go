package folderaccessapprovalsettings


type FolderAccessApprovalSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_access_approval_settings#create FolderAccessApprovalSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_access_approval_settings#delete FolderAccessApprovalSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/folder_access_approval_settings#update FolderAccessApprovalSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

