package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigGoo struct {
	// goo update settings. Use this setting to override the default goo patch rules.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#enabled OsConfigPatchDeployment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

