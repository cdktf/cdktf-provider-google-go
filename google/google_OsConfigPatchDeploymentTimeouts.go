// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type OsConfigPatchDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#create OsConfigPatchDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#delete OsConfigPatchDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

