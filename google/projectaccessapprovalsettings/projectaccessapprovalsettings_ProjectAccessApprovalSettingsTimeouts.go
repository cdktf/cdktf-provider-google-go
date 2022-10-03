package projectaccessapprovalsettings


type ProjectAccessApprovalSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_access_approval_settings#create ProjectAccessApprovalSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_access_approval_settings#delete ProjectAccessApprovalSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_access_approval_settings#update ProjectAccessApprovalSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

