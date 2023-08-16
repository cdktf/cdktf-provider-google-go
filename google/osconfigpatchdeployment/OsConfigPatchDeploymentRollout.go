package osconfigpatchdeployment


type OsConfigPatchDeploymentRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_patch_deployment#disruption_budget OsConfigPatchDeployment#disruption_budget}
	DisruptionBudget *OsConfigPatchDeploymentRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// Mode of the patch rollout. Possible values: ["ZONE_BY_ZONE", "CONCURRENT_ZONES"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_patch_deployment#mode OsConfigPatchDeployment#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

