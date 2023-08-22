package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResources struct {
	// The id of the resource with the following restrictions: Must contain only lowercase letters, numbers, and hyphens.
	//
	// Must start with a letter.
	// Must be between 1-63 characters.
	// Must end with a number or a letter.
	// Must be unique within the OS policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_os_policy_assignment#id OsConfigOsPolicyAssignment#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_os_policy_assignment#exec OsConfigOsPolicyAssignment#exec}
	Exec *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec `field:"optional" json:"exec" yaml:"exec"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_os_policy_assignment#file OsConfigOsPolicyAssignment#file}
	File *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile `field:"optional" json:"file" yaml:"file"`
	// pkg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_os_policy_assignment#pkg OsConfigOsPolicyAssignment#pkg}
	Pkg *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg `field:"optional" json:"pkg" yaml:"pkg"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_os_policy_assignment#repository OsConfigOsPolicyAssignment#repository}
	Repository *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository `field:"optional" json:"repository" yaml:"repository"`
}

