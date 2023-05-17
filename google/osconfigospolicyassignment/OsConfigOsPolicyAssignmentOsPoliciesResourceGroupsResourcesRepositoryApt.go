package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	// Required. Type of archive files in this repository. Possible values: ARCHIVE_TYPE_UNSPECIFIED, DEB, DEB_SRC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#archive_type OsConfigOsPolicyAssignment#archive_type}
	ArchiveType *string `field:"required" json:"archiveType" yaml:"archiveType"`
	// Required. List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#components OsConfigOsPolicyAssignment#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Required. Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#distribution OsConfigOsPolicyAssignment#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// Required. URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#uri OsConfigOsPolicyAssignment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// URI of the key file for this repository. The agent maintains a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#gpg_key OsConfigOsPolicyAssignment#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

