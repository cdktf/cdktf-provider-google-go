// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type OsConfigOsPolicyAssignmentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#fixed OsConfigOsPolicyAssignment#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#percent OsConfigOsPolicyAssignment#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

