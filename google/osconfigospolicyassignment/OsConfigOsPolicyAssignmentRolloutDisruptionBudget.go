package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_os_policy_assignment#fixed OsConfigOsPolicyAssignment#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_os_policy_assignment#percent OsConfigOsPolicyAssignment#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

