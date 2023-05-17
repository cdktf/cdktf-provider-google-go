package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigGoo struct {
	// goo update settings. Use this setting to override the default goo patch rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/os_config_patch_deployment#enabled OsConfigPatchDeployment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

