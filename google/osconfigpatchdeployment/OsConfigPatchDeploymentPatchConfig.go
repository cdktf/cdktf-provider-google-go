package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfig struct {
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#apt OsConfigPatchDeployment#apt}
	Apt *OsConfigPatchDeploymentPatchConfigApt `field:"optional" json:"apt" yaml:"apt"`
	// goo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#goo OsConfigPatchDeployment#goo}
	Goo *OsConfigPatchDeploymentPatchConfigGoo `field:"optional" json:"goo" yaml:"goo"`
	// Allows the patch job to run on Managed instance groups (MIGs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#mig_instances_allowed OsConfigPatchDeployment#mig_instances_allowed}
	MigInstancesAllowed interface{} `field:"optional" json:"migInstancesAllowed" yaml:"migInstancesAllowed"`
	// post_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#post_step OsConfigPatchDeployment#post_step}
	PostStep *OsConfigPatchDeploymentPatchConfigPostStep `field:"optional" json:"postStep" yaml:"postStep"`
	// pre_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#pre_step OsConfigPatchDeployment#pre_step}
	PreStep *OsConfigPatchDeploymentPatchConfigPreStep `field:"optional" json:"preStep" yaml:"preStep"`
	// Post-patch reboot settings. Possible values: ["DEFAULT", "ALWAYS", "NEVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#reboot_config OsConfigPatchDeployment#reboot_config}
	RebootConfig *string `field:"optional" json:"rebootConfig" yaml:"rebootConfig"`
	// windows_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#windows_update OsConfigPatchDeployment#windows_update}
	WindowsUpdate *OsConfigPatchDeploymentPatchConfigWindowsUpdate `field:"optional" json:"windowsUpdate" yaml:"windowsUpdate"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#yum OsConfigPatchDeployment#yum}
	Yum *OsConfigPatchDeploymentPatchConfigYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#zypper OsConfigPatchDeployment#zypper}
	Zypper *OsConfigPatchDeploymentPatchConfigZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

