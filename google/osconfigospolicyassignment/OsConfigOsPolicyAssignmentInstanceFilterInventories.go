package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentInstanceFilterInventories struct {
	// The OS short name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#os_short_name OsConfigOsPolicyAssignment#os_short_name}
	OsShortName *string `field:"required" json:"osShortName" yaml:"osShortName"`
	// The OS version Prefix matches are supported if asterisk(*) is provided as the last character.
	//
	// For example, to match all versions with a major version of '7', specify the following value for this field '7.*' An empty string matches all OS versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#os_version OsConfigOsPolicyAssignment#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
}

