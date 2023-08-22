package osconfigpatchdeployment


type OsConfigPatchDeploymentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#fixed OsConfigPatchDeployment#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#percentage OsConfigPatchDeployment#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

