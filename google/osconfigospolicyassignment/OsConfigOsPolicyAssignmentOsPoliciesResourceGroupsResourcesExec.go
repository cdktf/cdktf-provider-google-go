package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec struct {
	// validate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#validate OsConfigOsPolicyAssignment#validate}
	Validate *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate `field:"required" json:"validate" yaml:"validate"`
	// enforce block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#enforce OsConfigOsPolicyAssignment#enforce}
	Enforce *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce `field:"optional" json:"enforce" yaml:"enforce"`
}

