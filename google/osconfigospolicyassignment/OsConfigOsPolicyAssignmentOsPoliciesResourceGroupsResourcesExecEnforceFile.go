package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFile struct {
	// Defaults to false.
	//
	// When false, files are subject to validations based on the file type: Remote: A checksum must be specified. Cloud Storage: An object generation number must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#allow_insecure OsConfigOsPolicyAssignment#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#gcs OsConfigOsPolicyAssignment#gcs}
	Gcs *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFileGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// A local path within the VM to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#local_path OsConfigOsPolicyAssignment#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#remote OsConfigOsPolicyAssignment#remote}
	Remote *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFileRemote `field:"optional" json:"remote" yaml:"remote"`
}

