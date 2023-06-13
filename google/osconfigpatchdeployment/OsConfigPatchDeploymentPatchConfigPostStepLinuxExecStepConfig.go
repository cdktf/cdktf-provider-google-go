package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig struct {
	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#allowed_success_codes OsConfigPatchDeployment#allowed_success_codes}
	AllowedSuccessCodes *[]*float64 `field:"optional" json:"allowedSuccessCodes" yaml:"allowedSuccessCodes"`
	// gcs_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#gcs_object OsConfigPatchDeployment#gcs_object}
	GcsObject *OsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject `field:"optional" json:"gcsObject" yaml:"gcsObject"`
	// The script interpreter to use to run the script.
	//
	// If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#interpreter OsConfigPatchDeployment#interpreter}
	Interpreter *string `field:"optional" json:"interpreter" yaml:"interpreter"`
	// An absolute path to the executable on the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#local_path OsConfigPatchDeployment#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

