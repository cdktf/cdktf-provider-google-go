// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigPostStep struct {
	// linux_exec_step_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/os_config_patch_deployment#linux_exec_step_config OsConfigPatchDeployment#linux_exec_step_config}
	LinuxExecStepConfig *OsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig `field:"optional" json:"linuxExecStepConfig" yaml:"linuxExecStepConfig"`
	// windows_exec_step_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/os_config_patch_deployment#windows_exec_step_config OsConfigPatchDeployment#windows_exec_step_config}
	WindowsExecStepConfig *OsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig `field:"optional" json:"windowsExecStepConfig" yaml:"windowsExecStepConfig"`
}

