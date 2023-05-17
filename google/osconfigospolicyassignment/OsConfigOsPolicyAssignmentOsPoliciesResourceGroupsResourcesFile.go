package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile struct {
	// Required. The absolute path of the file within the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#path OsConfigOsPolicyAssignment#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Required. Desired state of the file. Possible values: OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED, COMPLIANT, NON_COMPLIANT, UNKNOWN, NO_OS_POLICIES_APPLICABLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#state OsConfigOsPolicyAssignment#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// A a file with this content. The size of the content is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#content OsConfigOsPolicyAssignment#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_os_policy_assignment#file OsConfigOsPolicyAssignment#file}
	File *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileFile `field:"optional" json:"file" yaml:"file"`
}

