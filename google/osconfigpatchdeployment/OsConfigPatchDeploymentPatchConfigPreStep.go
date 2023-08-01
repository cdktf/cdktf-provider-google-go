package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigPreStep struct {
	// linux_exec_step_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#linux_exec_step_config OsConfigPatchDeployment#linux_exec_step_config}
	LinuxExecStepConfig *OsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig `field:"optional" json:"linuxExecStepConfig" yaml:"linuxExecStepConfig"`
	// windows_exec_step_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#windows_exec_step_config OsConfigPatchDeployment#windows_exec_step_config}
	WindowsExecStepConfig *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig `field:"optional" json:"windowsExecStepConfig" yaml:"windowsExecStepConfig"`
}

