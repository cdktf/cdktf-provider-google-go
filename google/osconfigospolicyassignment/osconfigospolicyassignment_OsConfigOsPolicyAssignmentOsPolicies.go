package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPolicies struct {
	// Required.
	//
	// The id of the OS policy with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the assignment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#id OsConfigOsPolicyAssignment#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Required. Policy mode Possible values: MODE_UNSPECIFIED, VALIDATION, ENFORCEMENT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#mode OsConfigOsPolicyAssignment#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// resource_groups block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#resource_groups OsConfigOsPolicyAssignment#resource_groups}
	ResourceGroups interface{} `field:"required" json:"resourceGroups" yaml:"resourceGroups"`
	// This flag determines the OS policy compliance status when none of the resource groups within the policy are applicable for a VM.
	//
	// Set this value to `true` if the policy needs to be reported as compliant even if the policy has nothing to validate or enforce.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#allow_no_resource_group_match OsConfigOsPolicyAssignment#allow_no_resource_group_match}
	AllowNoResourceGroupMatch interface{} `field:"optional" json:"allowNoResourceGroupMatch" yaml:"allowNoResourceGroupMatch"`
	// Policy description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#description OsConfigOsPolicyAssignment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

