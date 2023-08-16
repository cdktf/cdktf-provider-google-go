package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#disruption_budget OsConfigOsPolicyAssignment#disruption_budget}
	DisruptionBudget *OsConfigOsPolicyAssignmentRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// This determines the minimum duration of time to wait after the configuration changes are applied through the current rollout.
	//
	// A VM continues to count towards the 'disruption_budget' at least until this duration of time has passed after configuration changes are applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#min_wait_duration OsConfigOsPolicyAssignment#min_wait_duration}
	MinWaitDuration *string `field:"required" json:"minWaitDuration" yaml:"minWaitDuration"`
}

